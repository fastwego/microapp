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

package main

type Param struct {
	Name string
	Type string
}

type Api struct {
	Name        string
	Description string
	Request     string
	See         string
	FuncName    string
	GetParams   []Param
}

type ApiGroup struct {
	Name    string
	Apis    []Api
	Package string
}

var apiConfig = []ApiGroup{
	{
		Name:    `登录`,
		Package: `auth`,
		Apis: []Api{
			{
				Name:        "code2Session",
				Description: "通过login接口获取到登录凭证后，开发者可以通过服务器发送请求的方式获取 session_key 和 openId。",
				Request:     "GET https://developer.toutiao.com/api/apps/jscode2session",
				See:         "https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/log-in/code-2-session",
				FuncName:    "Code2Session",
				GetParams: []Param{
					{Name: "appid", Type: "string"},
				},
			},
		},
	},
	{
		Name:    `数据缓存`,
		Package: `data_caching`,
		Apis: []Api{
			{
				Name:        "setUserStorage",
				Description: "以 key-value 形式存储用户数据到小程序平台的云存储服务。若开发者无内部存储服务则可接入，免费且无需申请。一般情况下只存储用户的基本信息，禁止写入大量不相干信息。",
				Request:     "POST https://developer.toutiao.com/api/apps/set_user_storage",
				See:         "https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/data-caching/set-user-storage",
				FuncName:    "SetUserStorage",
				GetParams: []Param{
					{Name: "openid", Type: "string"},
					{Name: "signature", Type: "string"},
					{Name: "sig_method", Type: "string"},
				},
			},
			{
				Name:        "removeUserStorage",
				Description: "删除存储到字节跳动的云存储服务的 key-value 数据。当开发者不需要该用户信息时，需要删除，以免占用过大的存储空间。",
				Request:     "POST https://developer.toutiao.com/api/apps/remove_user_storage",
				See:         "https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/data-caching/remove-user-storage",
				FuncName:    "RemoveUserStorage",
				GetParams: []Param{
					{Name: "openid", Type: "string"},
					{Name: "signature", Type: "string"},
					{Name: "sig_method", Type: "string"},
				},
			}},
	},
	{
		Name:    `二维码`,
		Package: `qrcode`,
		Apis: []Api{{
			Name:        "createQRCode",
			Description: "获取小程序/小游戏的二维码。该二维码可通过任意 app 扫码打开，能跳转到开发者指定的对应字节系 app 内拉起小程序/小游戏，并传入开发者指定的参数。通过该接口生成的二维码，永久有效，暂无数量限制。",
			Request:     "POST https://developer.toutiao.com/api/apps/qrcode",
			See:         "https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/qr-code/create-qr-code",
			FuncName:    "CreateQRCode",
		}},
	},
	{
		Name:    `模板消息`,
		Package: `template_message`,
		Apis: []Api{{
			Name:        "发送模版消息",
			Description: "提示 本接口在服务器端调用 目前只有今日头条支持，抖音和 lite 接入中",
			Request:     "POST https://developer.toutiao.com/api/apps/game/template/send",
			See:         "https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/model-news/send",
			FuncName:    "Send",
		}},
	},
	{
		Name:    `内容安全`,
		Package: `content_security`,
		Apis: []Api{{
			Name:        "内容安全检测",
			Description: "检测一段文本是否包含违法违规内容。",
			Request:     "POST https://developer.toutiao.com/api/v2/tags/text/antidirt",
			See:         "https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/content-security/content-security-detect",
			FuncName:    "TextAntiDirty",
		},
			{
				Name:        "图片检测",
				Description: "检测图片是否包含违法违规内容。",
				Request:     "POST https://developer.toutiao.com/api/v2/tags/image/",
				See:         "https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/content-security/picture-detect",
				FuncName:    "Image",
			}},
	},
	{
		Name:    `订阅消息`,
		Package: `subscribe_notification`,
		Apis: []Api{
			{
				Name:        "订阅消息推送",
				Description: "用户产生了订阅模板消息的行为后，可以通过这个接口发送模板消息给用户，功能参考订阅消息能力。",
				Request:     "POST https://developer.toutiao.com/api/apps/subscribe_notification/developer/v1/notify",
				See:         "https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/subscribe-notification/notify",
				FuncName:    "Notify",
			}},
	},
}
