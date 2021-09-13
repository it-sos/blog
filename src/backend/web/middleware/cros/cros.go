package cros

import (
	"gitee.com/itsos/studynotes/web/bootstrap"
	"github.com/kataras/iris/v12"
)

// Configure creates a new identity middleware and registers that to the app.
// 跨域处理
func Configure(b *bootstrap.Bootstrapper) {
	b.UseGlobal(
		func(ctx iris.Context) {
			// 设置允许跨域访问
			ctx.Header("Access-Control-Allow-Origin", "*")
			ctx.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE")

			// 预检查 options 直接放行
			if ctx.Method() == "OPTIONS" {
				ctx.StatusCode(iris.StatusOK)
				return
			}
			ctx.Next()
		})
}
