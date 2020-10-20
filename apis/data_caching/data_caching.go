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

// Package data_caching 数据缓存
package data_caching

import (
	"bytes"
	"net/url"

	"github.com/fastwego/microapp"
)

const (
	apiSetUserStorage    = "/api/apps/set_user_storage"
	apiRemoveUserStorage = "/api/apps/remove_user_storage"
)

/*
setUserStorage

以 key-value 形式存储用户数据到小程序平台的云存储服务。若开发者无内部存储服务则可接入，免费且无需申请。一般情况下只存储用户的基本信息，禁止写入大量不相干信息。

See: https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/data-caching/set-user-storage

POST https://developer.toutiao.com/api/apps/set_user_storage
*/
func SetUserStorage(ctx *microapp.MicroApp, payload []byte, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetUserStorage+"?"+params.Encode(), bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
removeUserStorage

删除存储到字节跳动的云存储服务的 key-value 数据。当开发者不需要该用户信息时，需要删除，以免占用过大的存储空间。

See: https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/data-caching/remove-user-storage

POST https://developer.toutiao.com/api/apps/remove_user_storage
*/
func RemoveUserStorage(ctx *microapp.MicroApp, payload []byte, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiRemoveUserStorage+"?"+params.Encode(), bytes.NewReader(payload), "application/json;charset=utf-8")
}
