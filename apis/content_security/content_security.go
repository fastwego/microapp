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

// Package content_security 内容安全
package content_security

import (
	"bytes"
	"net/http"

	"github.com/fastwego/microapp"
)

const (
	apiTextAntiDirty = "/api/v2/tags/text/antidirt"
	apiImage         = "/api/v2/tags/image/"
)

/*
内容安全检测

检测一段文本是否包含违法违规内容。

See: https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/content-security/content-security-detect

POST https://developer.toutiao.com/api/v2/tags/text/antidirt
*/
func TextAntiDirty(ctx *microapp.MicroApp, payload []byte) (resp []byte, err error) {
	req, err := http.NewRequest(http.MethodPost, microapp.ServerUrl+apiTextAntiDirty, bytes.NewReader(payload))
	if err != nil {
		return
	}

	var accessToken string
	accessToken, err = ctx.GetAccessTokenHandler(ctx)
	if err != nil {
		return
	}
	req.Header.Add("X-Token", accessToken)
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	return ctx.Client.HTTPDo(req)
}

/*
图片检测

检测图片是否包含违法违规内容。

See: https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/content-security/picture-detect

POST https://developer.toutiao.com/api/v2/tags/image/
*/
func Image(ctx *microapp.MicroApp, payload []byte) (resp []byte, err error) {

	req, err := http.NewRequest(http.MethodPost, microapp.ServerUrl+apiImage, bytes.NewReader(payload))
	if err != nil {
		return
	}

	var accessToken string
	accessToken, err = ctx.GetAccessTokenHandler(ctx)
	if err != nil {
		return
	}
	req.Header.Add("X-Token", accessToken)
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	return ctx.Client.HTTPDo(req)
}
