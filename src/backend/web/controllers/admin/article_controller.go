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
	"bytes"
	"encoding/json"
	"gitee.com/itsos/studynotes/models/vo"
	"gitee.com/itsos/studynotes/services"
	"github.com/kataras/iris/v12"
	"strconv"
	"time"
)

type ArticleController struct {
	Ctx       iris.Context
	StartTime time.Time
}

// DeleteArticle
// @Tags 博客后台接口
// @Summary 删除文章
// @Description 删除文章
// @Accept json
// @Produce json
// @Param id query integer true "文章id"
// @Success 200 {string} string ""
// @Failure 400 {object} cerrors.Errors "error"
// @Failure 401 {string} string "未授权登录"
// @Router /admin/article [delete]
func (c *ArticleController) DeleteArticle() {
	id, _ := strconv.Atoi(c.Ctx.FormValue("id"))
	services.SArticle.DeleteArticle(uint(id))
}

// PostArticle
// @Tags 博客后台接口
// @Summary 保存文章
// @Description 保存文章
// @Accept application/x-www-form-urlencoded
// @Produce plain
// @Param body body vo.ArticleParamsVO true "文章相关内容"
// @Success 200 {integer} integer "文章id"
// @Failure 400 {object} cerrors.Errors "error"
// @Failure 401 {string} string "未授权登录"
// @Router /admin/article [post]
func (c *ArticleController) PostArticle() (id uint, err error) {
	body, _ := c.Ctx.GetBody()
	article := new(vo.ArticleParamsVO)
	if err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&article); err != nil {
		return
	}
	return services.SArticle.SaveArticle(*article)
}

// GetArticle
// @Tags 博客后台接口
// @Summary 查询文章及相关信息
// @Description 查询文章及相关信息
// @Accept json
// @Produce json
// @Param id query integer true "文章id"
// @Success 200 {object} vo.ArticleEditVO "文章详情VO"
// @Failure 400 {object} cerrors.Errors "error"
// @Failure 401 {string} string "未授权登录"
// @Router /admin/article [get]
func (c *ArticleController) GetArticle() (vo.ArticleEditVO, error) {
	id, _ := c.Ctx.URLParamInt("id")
	return services.SArticle.GetArticleAndContent(uint(id))
}

// GetArticles
// @Tags 博客后台接口
// @Summary 查询文章列表
// @Description 查询文章列表
// @Accept json
// @Produce json
// @Param keyword query string false "关键词"
// @Param page query integer true "页码"
// @Param size query integer true "每页条数"
// @Success 200 {object} vo.ArticleListVO "文章列表VO"
// @Failure 400 {object} cerrors.Errors "error"
// @Failure 401 {string} string "未授权登录"
// @Router /admin/articles [get]
func (c *ArticleController) GetArticles() []vo.ArticleListVO {
	page, _ := strconv.Atoi(c.Ctx.FormValue("page"))
	size, _ := strconv.Atoi(c.Ctx.FormValue("size"))
	keyword := c.Ctx.FormValue("keyword")
	return services.SArticle.GetArticleList(page, size, keyword)
}
