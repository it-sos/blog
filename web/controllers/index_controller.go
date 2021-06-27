/*
   Copyright (c) [2021] IT.SOS
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package controllers

import (
	"gitee.com/itsos/studynotes/services"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"time"
)

type IndexController struct {
	Ctx       iris.Context
	StartTime time.Time
	Sess      *sessions.Session
}

// GetCookie
// @Tags 测试方法
// @Summary 临时用
// @Description 临时用
// @Accept json
// @Produce json
// @Success 200 {string} string "测试结果"
// @Failure 400 {object} errors.Errors "error"
// @Router /cookie [get]
func (c *IndexController) GetCookie() interface{} {
	return services.SArticle.GetRank()
	////id := c.Ctx.FormValue("id")
	//visits := c.Sess.Increment("visits", 1)
	//// 写下当前的，更新的访问.
	//since := time.Now().Sub(c.StartTime).Seconds()
	//return fmt.Sprintf("%d visit(s) from my current session in %0.1f seconds of server's up-time",
	//	visits, since)
}
