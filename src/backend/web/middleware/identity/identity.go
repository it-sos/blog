package identity

import (
	"gitee.com/itsos/studynotes/web/bootstrap"
	"github.com/kataras/iris/v12"
	"time"
)

// Configure creates a new identity middleware and registers that to the app.
func Configure(b *bootstrap.Bootstrapper) {
	b.UseGlobal(func(ctx iris.Context) {
		if ctx.GetStatusCode() < iris.StatusInternalServerError {
			// 重置 session 未操作计时器，长时不操作 session 自动销毁
			//b.Sessions.ShiftExpiration(ctx)

			// response headers
			ctx.Header("App-Name", b.AppName)
			ctx.Header("App-Owner", b.AppOwner)
			ctx.Header("App-Since", time.Since(b.AppSpawnDate).String())

			// view data if ctx.View or c.Tmpl = "$page.html" will be called next.
			ctx.ViewData("AppName", b.AppName)
			ctx.ViewData("AppOwner", b.AppOwner)
		}

		ctx.Next()
	})
}
