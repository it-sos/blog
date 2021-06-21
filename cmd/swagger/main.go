/*
   Copyright (c) [2021] IT.SOS
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package main

import (
	"fmt"
	"gitee.com/itsos/golibs/config"
	C "gitee.com/itsos/study-notes/config"
	_ "gitee.com/itsos/study-notes/docs"
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
)

var c = C.C

func main() {
	config.C.Init()
	app := iris.New()
	url := swagger.URL(c.GetSwaggerUrl() + "/swagger/doc.json")
	app.Get("/swagger/{any:path}", swagger.WrapHandler(swaggerFiles.Handler, url))

	fmt.Printf("Swagger on: %s/swagger/index.html\n", c.GetSwaggerUrl())
	app.Listen(":" + c.GetSwaggerPort())
}
