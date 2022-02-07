/*
   Copyright (c) [2021] IT.SOS
   study-notes is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package auth

import (
	"gitee.com/itsos/studynotes/services"
	"github.com/kataras/iris/v12"
)

func Secret(ctx iris.Context) {
	token := ctx.GetHeader("token")
	userId, err := services.SAuthService.GetUserId(token)
	if err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.WriteString(err.Error())
		return
	}
	ctx.Values().Set("userId", userId)
	ctx.Next()
}
