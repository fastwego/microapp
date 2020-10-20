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

// Package subscribe_notification 订阅消息
package subscribe_notification

import (
	"bytes"

	"github.com/fastwego/microapp"
)

const (
	apiNotify = "/api/apps/subscribe_notification/developer/v1/notify"
)

/*
订阅消息推送

用户产生了订阅模板消息的行为后，可以通过这个接口发送模板消息给用户，功能参考订阅消息能力。

See: https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/subscribe-notification/notify

POST https://developer.toutiao.com/api/apps/subscribe_notification/developer/v1/notify
*/
func Notify(ctx *microapp.MicroApp, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiNotify, bytes.NewReader(payload), "application/json;charset=utf-8")
}
