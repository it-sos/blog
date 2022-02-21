/*
   Copyright (c) [2021] IT.SOS
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package routes

import (
	bootstrap2 "gitee.com/itsos/golibs/v2/framework/iris/bootstrap"
	"gitee.com/itsos/golibs/v2/framework/iris/middleware/auth"
	"gitee.com/itsos/studynotes/web/controllers"
	"gitee.com/itsos/studynotes/web/controllers/admin"
	"github.com/kataras/iris/v12/mvc"
	"time"
)

func Routes(b *bootstrap2.Bootstrapper) {
	mvc.Configure(
		b.Party("/"),
		func(app *mvc.Application) {
			app.Handle(new(controllers.IndexController)).
				Handle(new(controllers.FilesController)).
				Handle(new(controllers.AuthController))
		},
	).Register(time.Now())

	mvc.Configure(
		b.Party("/admin", auth.Secret),
		func(app *mvc.Application) {
			app.Handle(new(admin.ArticleController)).
				Handle(new(admin.FilesController)).
				Handle(new(admin.CategoryController))
		},
	).Register(time.Now())
}
