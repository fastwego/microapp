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

// Package qrcode 二维码
package qrcode

import (
	"bytes"

	"github.com/fastwego/microapp"
)

const (
	apiCreateQRCode = "/api/apps/qrcode"
)

/*
createQRCode

获取小程序/小游戏的二维码。该二维码可通过任意 app 扫码打开，能跳转到开发者指定的对应字节系 app 内拉起小程序/小游戏，并传入开发者指定的参数。通过该接口生成的二维码，永久有效，暂无数量限制。

See: https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/qr-code/create-qr-code

POST https://developer.toutiao.com/api/apps/qrcode
*/
func CreateQRCode(ctx *microapp.MicroApp, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreateQRCode, bytes.NewReader(payload), "application/json;charset=utf-8")
}
