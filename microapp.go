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

/*
字节小程序开发 SDK

See: https://microapp.bytedance.com/
*/
package microapp

import (
	"log"
	"os"

	"github.com/faabiosr/cachego"
	"github.com/faabiosr/cachego/file"
)

// GetAccessTokenFunc 获取 access_token 方法接口
type GetAccessTokenFunc func(ctx *MicroApp) (accessToken string, err error)

// NoticeAccessTokenExpireFunc 通知中控 刷新 access_token
type NoticeAccessTokenExpireFunc func(ctx *MicroApp) (err error)

/*
MicroApp 实例
*/
type MicroApp struct {
	Config                         Config
	Client                         Client
	Logger                         *log.Logger
	Cache                          cachego.Cache
	GetAccessTokenHandler          GetAccessTokenFunc
	NoticeAccessTokenExpireHandler NoticeAccessTokenExpireFunc
}

/*
小程序配置
*/
type Config struct {
	AppId     string
	AppSecret string
}

/*
创建小程序实例
*/
func New(config Config) (microapp *MicroApp) {
	instance := MicroApp{
		Config:                         config,
		Cache:                          file.New(os.TempDir()),
		GetAccessTokenHandler:          GetAccessToken,
		NoticeAccessTokenExpireHandler: NoticeAccessTokenExpire,
	}

	instance.Client = Client{Ctx: &instance}
	instance.Logger = log.New(os.Stdout, "[fastwego/microapp] ", log.LstdFlags|log.Llongfile)

	return &instance
}
