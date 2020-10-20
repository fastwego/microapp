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

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const ServerUrl = `https://microapp.bytedance.com`

var apiUniqMap = map[string]bool{}

func run() {

	file, err := ioutil.ReadFile("./cmd/data/doc_links.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	pattern := `href="(/docs/zh-CN/mini-app/develop/server/.+)"`
	reg := regexp.MustCompile(pattern)
	matched := reg.FindAllStringSubmatch(string(file), -1)

	for _, match := range matched {

		link := ServerUrl + match[1]
		resp, err := http.Get(link)
		if err != nil {
			continue
		}

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		_NAME_ := doc.Find(".markdown-render-content h1").First().Text()

		_REQUEST_ := strings.TrimSpace(doc.Find(".markdown-render-content h2#请求地址").Next().Text())
		//fmt.Println(_REQUEST_)

		if _REQUEST_ == "" || _NAME_ == "" {
			fmt.Println("Error")
			continue
		}

		_DESCRIPTION_ := doc.Find(".markdown-render-content h1").Next().Text()

		_SEE_ := link
		_FUNCNAME_ := _NAME_

		//GetParams: []Param{
		//	{Name: `appid`, Type: `string`},
		//	{Name: `redirect_uri`, Type: `string`},
		//},

		_GET_PARAMS_ := ""
		fields := strings.Fields(_REQUEST_)
		parse, _ := url.Parse(fields[1])
		for param_name, _ := range parse.Query() {
			if param_name == "access_token" {
				continue
			}
			_GET_PARAMS_ += "\t\t{Name: `" + param_name + "`, Type: `string`},\n"
		}
		if _GET_PARAMS_ != "" {
			_GET_PARAMS_ = `GetParams: []Param{
` + _GET_PARAMS_ + "\t},"
		}

		if _, ok := apiUniqMap[_REQUEST_]; !ok {
			apiUniqMap[_REQUEST_] = true
		} else {
			continue
		}

		tpl := strings.ReplaceAll(itemTpl, "_NAME_", _NAME_)
		tpl = strings.ReplaceAll(tpl, "_DESCRIPTION_", strings.TrimSpace(_DESCRIPTION_))
		tpl = strings.ReplaceAll(tpl, "_REQUEST_", strings.TrimSpace(_REQUEST_))
		tpl = strings.ReplaceAll(tpl, "_SEE_", _SEE_)
		tpl = strings.ReplaceAll(tpl, "_FUNCNAME_", _FUNCNAME_)
		tpl = strings.ReplaceAll(tpl, "_GET_PARAMS_", _GET_PARAMS_)

		fmt.Println(tpl)

	}

}

var itemTpl = `{
	Name:        "_NAME_",
	Description: "_DESCRIPTION_",
	Request:     "_REQUEST_",
	See:         "_SEE_",
	FuncName:    "_FUNCNAME_",
	_GET_PARAMS_
},`
