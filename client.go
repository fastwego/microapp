// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package microapp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var (
	ServerUrl              = "https://developer.toutiao.com" //  api 服务器地址
	UserAgent              = "fastwego/microapp"
	ErrorAccessTokenExpire = errors.New("access token expire")
	ErrorSystemBusy        = errors.New("system busy")
)

/*
HttpClient 用于向接口发送请求
*/
type Client struct {
	Ctx *MicroApp
}

// HTTPGet GET 请求
func (client *Client) HTTPGet(uri string) (resp []byte, err error) {

	req, err := http.NewRequest(http.MethodGet, ServerUrl+uri, nil)
	if err != nil {
		return
	}

	return client.HTTPDo(req)
}

//HTTPPost POST 请求
func (client *Client) HTTPPost(uri string, payload io.Reader, contentType string) (resp []byte, err error) {

	req, err := http.NewRequest(http.MethodPost, ServerUrl+uri, payload)
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", contentType)

	return client.HTTPDo(req)
}

//HTTPDo 执行 请求
func (client *Client) HTTPDo(req *http.Request) (resp []byte, err error) {

	var body, body2 []byte
	if req.Body != nil {

		body, err = ioutil.ReadAll(req.Body)
		if err != nil {
			return
		}

		body2 = make([]byte, len(body))
		copy(body2, body)
		req.Body = ioutil.NopCloser(bytes.NewReader(body))
	}

	req.Header.Add("User-Agent", UserAgent)

	if client.Ctx.Logger != nil {
		client.Ctx.Logger.Printf("%s %s Headers %v", req.Method, req.URL.String(), req.Header)
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()

	resp, err = responseFilter(response)

	// 发现 access_token 过期
	if err == ErrorAccessTokenExpire {

		// 主动 通知 access_token 过期
		err = client.Ctx.NoticeAccessTokenExpireHandler(client.Ctx)
		if err != nil {
			return
		}

		// 通知到位后 access_token 会被刷新，那么可以 retry 了
		var accessToken string
		accessToken, err = client.Ctx.GetAccessTokenHandler(client.Ctx)
		if err != nil {
			return
		}

		// 换新
		q := req.URL.Query()
		if q.Get("access_token") != "" {
			q.Set("access_token", accessToken)
			req.URL.RawQuery = q.Encode()
		} else if req.Header.Get("X-Token") != "" {
			req.Header.Set("X-Token", accessToken)
		} else {
			if len(body2) > 0 {
				jsonData := map[string]interface{}{}
				err = json.Unmarshal(body2, &jsonData)
				if err != nil {
					return
				}
				if _, ok := jsonData["access_token"]; ok {
					jsonData["access_token"] = accessToken
					body2, err = json.Marshal(jsonData)
					if err != nil {
						return
					}
				}
			}
		}

		if client.Ctx.Logger != nil {
			client.Ctx.Logger.Printf("%v retry %s %s Headers %v", ErrorAccessTokenExpire, req.Method, req.URL.String(), req.Header)
		}

		req.Body = ioutil.NopCloser(bytes.NewReader(body2))
		req.ContentLength = int64(len(body2))
		response, err = http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		resp, err = responseFilter(response)
	} else if err == ErrorSystemBusy {

		if client.Ctx.Logger != nil {
			client.Ctx.Logger.Printf("%v : retry %s %s Headers %v", ErrorSystemBusy, req.Method, req.URL.String(), req.Header)
		}

		req.Body = ioutil.NopCloser(bytes.NewReader(body2))
		req.ContentLength = int64(len(body2))
		response, err = http.DefaultClient.Do(req)
		if err != nil {
			return
		}

		resp, err = responseFilter(response)
	}

	return
}

/*
筛查 api 服务器响应，判断以下错误：

- http 状态码 不为 200

- 接口响应错误码 errcode 不为 0
*/
func responseFilter(response *http.Response) (resp []byte, err error) {
	if response.StatusCode != http.StatusOK {

		if response.StatusCode == 401 { // 401 Unauthorized
			err = ErrorAccessTokenExpire
			return
		}

		err = fmt.Errorf("Status %s", response.Status)
		return
	}

	resp, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	errorResponse := struct {
		Errcode int64  `json:"errcode"`
		Errmsg  string `json:"errmsg"`
	}{}
	err = json.Unmarshal(resp, &errorResponse)
	if err != nil {
		return
	}

	if errorResponse.Errcode == 40002 { // bad access_token
		err = ErrorAccessTokenExpire
		return
	}
	//  -1	系统繁忙，此时请开发者稍候再试
	if errorResponse.Errcode == -1 {
		err = ErrorSystemBusy
		return
	}

	if errorResponse.Errcode != 0 {
		err = errors.New(string(resp))
		return
	}
	return
}

// 防止多个 goroutine 并发刷新冲突
var refreshAccessTokenLock sync.Mutex

/*
从 公众号实例 的 AccessToken 管理器 获取 access_token

如果没有 access_token 或者 已过期，那么刷新

获得新的 access_token 后 过期时间设置为 0.9 * expiresIn 提供一定冗余
*/
func GetAccessToken(ctx *MicroApp) (accessToken string, err error) {
	accessToken, err = ctx.Cache.Fetch(ctx.Config.AppId)
	if accessToken != "" {
		return
	}

	refreshAccessTokenLock.Lock()
	defer refreshAccessTokenLock.Unlock()

	accessToken, err = ctx.Cache.Fetch(ctx.Config.AppId)
	if accessToken != "" {
		return
	}

	accessToken, expiresIn, err := refreshAccessToken(ctx.Config.AppId, ctx.Config.AppSecret)
	if err != nil {
		return
	}

	// 本地缓存 access_token
	d := time.Duration(expiresIn) * time.Second
	_ = ctx.Cache.Save(ctx.Config.AppId, accessToken, d)

	if ctx.Logger != nil {
		ctx.Logger.Printf("%s %s %d\n", "refreshAccessToken", accessToken, expiresIn)
	}

	return
}

/*
NoticeAccessTokenExpire 只需将本地存储的 access_token 删除，即完成了 access_token 已过期的 主动通知

retry 请求的时候，会发现本地没有 access_token ，从而触发refresh
*/
func NoticeAccessTokenExpire(ctx *MicroApp) (err error) {
	if ctx.Logger != nil {
		ctx.Logger.Println("NoticeAccessTokenExpire")
	}

	err = ctx.Cache.Delete(ctx.Config.AppId)
	return
}

/*
从服务器获取新的 AccessToken

See: https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_access_token.html
*/
func refreshAccessToken(appid string, secret string) (accessToken string, expiresIn int, err error) {
	params := url.Values{}
	params.Add("appid", appid)
	params.Add("secret", secret)
	params.Add("grant_type", "client_credential")
	url := ServerUrl + "/api/apps/token?" + params.Encode()

	response, err := http.Get(url)
	if err != nil {
		return
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("GET %s RETURN %s", url, response.Status)
		return
	}

	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	var result = struct {
		AccessToken string  `json:"access_token"`
		ExpiresIn   int     `json:"expires_in"`
		Errcode     float64 `json:"errcode"`
		Errmsg      string  `json:"errmsg"`
	}{}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		err = fmt.Errorf("Unmarshal error %s", string(resp))
		return
	}

	if result.AccessToken == "" {
		err = fmt.Errorf("%s", string(resp))
		return
	}

	return result.AccessToken, result.ExpiresIn, nil
}
