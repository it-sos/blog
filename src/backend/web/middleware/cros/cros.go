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
			ctx.Header("Access-Control-Allow-Origin", "https://itsos.ltd")
			ctx.Header("Access-Control-Allow-Credentials", "true")
			ctx.Header("Access-Control-Allow-Methods", "*")
			ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Type, Depth,User-Agent, X-File-Size, X-Requested-With, X-Requested-By, If-Modified-Since, X-File-Name, X-File-Type, Cache-Control, Origin, token, x-requested-with, TE")
			ctx.Header("Access-Control-Expose-Headers", "*")

			// 预检查 options 直接放行
			if ctx.Method() == "OPTIONS" {
				ctx.StatusCode(iris.StatusOK)
				return
			}
			ctx.Next()
		})
}
