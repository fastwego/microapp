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

package content_security_test

import (
	"fmt"

	"github.com/fastwego/microapp"
	"github.com/fastwego/microapp/apis/content_security"
)

func ExampleTextAntiDirty() {
	var ctx *microapp.MicroApp

	payload := []byte("{}")
	resp, err := content_security.TextAntiDirty(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleImage() {
	var ctx *microapp.MicroApp

	payload := []byte("{}")
	resp, err := content_security.Image(ctx, payload)

	fmt.Println(resp, err)
}
