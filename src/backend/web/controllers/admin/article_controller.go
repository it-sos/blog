/*
   Copyright (c) [2021] IT.SOS
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"time"
)

type ArticleController struct {
	Ctx       iris.Context
	StartTime time.Time
	Sess      *sessions.Session
}

// GetArticleContent
// @Tags 博客后台接口
// @Summary 文章内容
// @Description 通过文章id获取文章详情
// @Accept json
// @Produce json
// @Param id query integer true "文章id"
// @Success 200 {object} vo.ArticleVO "文章详情VO"
// @Failure 400 {object} errors.Errors "error"
// @Router /admin/article/content [get]
func (c *ArticleController) GetArticleContent() interface{} {
	//isLogin, _ := c.Sess.GetBoolean("authenticated")
	//id, _ := strconv.Atoi(c.Ctx.FormValue("id"))
	return "{1}"
}
