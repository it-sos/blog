/*
   Copyright (c) [2021] itsos
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
               http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package studynotes

import (
	"gitee.com/itsos/golibs/global/variable"
	"gitee.com/itsos/studynotes/config"
	"gitee.com/itsos/studynotes/web/bootstrap"
	"gitee.com/itsos/studynotes/web/middleware/identity"
	"gitee.com/itsos/studynotes/web/routes"
	"github.com/kataras/iris/v12"
)

func newApp() *bootstrap.Bootstrapper {
	println(variable.BasePath)
	app := bootstrap.New("studynotes", "peng.yu@qjfu.cn")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func Listen() {
	newApp().Listen(":"+config.C.GetPort(), iris.WithOptimizations)
}
