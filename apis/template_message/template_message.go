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

// Package template_message 模板消息
package template_message

import (
	"bytes"

	"github.com/fastwego/microapp"
)

const (
	apiSend = "/api/apps/game/template/send"
)

/*
发送模版消息

提示 本接口在服务器端调用 目前只有今日头条支持，抖音和 lite 接入中

See: https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/model-news/send

POST https://developer.toutiao.com/api/apps/game/template/send
*/
func Send(ctx *microapp.MicroApp, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSend, bytes.NewReader(payload), "application/json;charset=utf-8")
}
