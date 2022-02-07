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
	"gitee.com/itsos/studynotes/datamodels"
	"gitee.com/itsos/studynotes/services"
	"github.com/kataras/iris/v12"
	"strconv"
	"time"
)

type CategoryController struct {
	Ctx       iris.Context
	StartTime time.Time
}

// DeleteCategoryRelations
// @Tags 博客后台接口
// @Summary 解除绑定关系
// @Description 通过文章id和分类id解除绑定关系
// @Accept json
// @Produce json
// @Param id query integer true "类别id"
// @Param aid query integer true "文章id"
// @Success 200 {string} string "success"
// @Failure 400 {object} cerrors.Errors "error"
// @Router /admin/category/relations [delete]
func (c *CategoryController) DeleteCategoryRelations() error {
	id, _ := strconv.Atoi(c.Ctx.FormValue("id"))
	aid, _ := strconv.Atoi(c.Ctx.FormValue("aid"))
	return services.SCategory.Unbind(uint(id), uint(aid))
}

// PostCategoryTopic
// @Tags 博客后台接口
// @Summary 新增专题
// @Description 新增专题
// @Accept json
// @Produce json
// @Param name query string true "专题名称"
// @Success 200 {integer} integer "专题id"
// @Failure 400 {object} cerrors.Errors "error"
// @Router /admin/category/topic [post]
func (c *CategoryController) PostCategoryTopic() (uint, error) {
	name := c.Ctx.FormValue("name")
	return services.SCategory.NewTopic(name)
}

// PostCategoryBindtopic
// @Tags 博客后台接口
// @Summary 绑定专题与文章
// @Description 绑定专题与文章
// @Accept json
// @Produce json
// @Param id query integer true "专题id"
// @Param aid query integer true "文章id"
// @Success 200 {integer} string ""
// @Failure 400 {object} cerrors.Errors "error"
// @Router /admin/category/bindtopic [post]
func (c *CategoryController) PostCategoryBindtopic() {
	id, _ := strconv.Atoi(c.Ctx.FormValue("id"))
	aid, _ := strconv.Atoi(c.Ctx.FormValue("aid"))
	services.SCategory.BindTopic(uint(id), uint(aid))
}

// DeleteCategoryTopic
// @Tags 博客后台接口
// @Summary 删除专题
// @Description 通过专题id删除专题
// @Accept json
// @Produce json
// @Param id query integer true "专题id"
// @Success 200 {string} string "success"
// @Failure 400 {object} cerrors.Errors "error"
// @Router /admin/category/topic [delete]
func (c *CategoryController) DeleteCategoryTopic() error {
	id, _ := strconv.Atoi(c.Ctx.FormValue("id"))
	return services.SCategory.DeleteTopic(uint(id))
}

// PutCategoryTopic
// @Tags 博客后台接口
// @Summary 更新专题
// @Description 更新专题
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id formData integer true "专题id"
// @Param name formData string true "专题名称"
// @Success 200 {string} string "success"
// @Failure 400 {object} cerrors.Errors "error"
// @Router /admin/category/topic [put]
func (c *CategoryController) PutCategoryTopic() error {
	id, _ := strconv.Atoi(c.Ctx.FormValue("id"))
	name := c.Ctx.FormValue("name")
	return services.SCategory.UpdateTopic(uint(id), name)
}

// GetCategoryTopics
// @Tags 博客后台接口
// @Summary 查询专题列表
// @Description 返回专题列表
// @Accept json
// @Produce json
// @Success 200 {object} []datamodels.Category "专题列表"
// @Failure 400 {object} cerrors.Errors "error"
// @Router /admin/category/topics [get]
func (c *CategoryController) GetCategoryTopics() []datamodels.Category {
	return services.SCategory.GetTopicList()
}

// PostCategoryTag
// @Tags 博客后台接口
// @Summary 新增标签
// @Description 新增标签
// @Accept json
// @Produce plain
// @Param name query integer true "标签名称"
// @Success 200 {integer} integer "专题id"
// @Failure 400 {object} cerrors.Errors "error"
// @Router /admin/category/tag [post]
func (c *CategoryController) PostCategoryTag() (uint, error) {
	name := c.Ctx.FormValue("name")
	return services.SCategory.NewTag(name)
}

// PostCategoryBindtag
// @Tags 博客后台接口
// @Summary 绑定标签与文章
// @Description 绑定标签与文章
// @Accept json
// @Produce json
// @Param id query integer true "专题id"
// @Param aid query integer true "文章id"
// @Success 200 {integer} string ""
// @Failure 400 {object} cerrors.Errors "error"
// @Router /admin/category/bindtag [post]
func (c *CategoryController) PostCategoryBindtag() {
	id, _ := strconv.Atoi(c.Ctx.FormValue("id"))
	aid, _ := strconv.Atoi(c.Ctx.FormValue("aid"))
	services.SCategory.BindTag(uint(id), uint(aid))
}

// DeleteCategoryTag
// @Tags 博客后台接口
// @Summary 删除标签
// @Description 删除标签
// @Accept json
// @Produce json
// @Param id query integer true "标签id"
// @Success 200 {string} string "success"
// @Failure 400 {object} cerrors.Errors "error"
// @Router /admin/category/tag [delete]
func (c *CategoryController) DeleteCategoryTag() error {
	id, _ := strconv.Atoi(c.Ctx.FormValue("id"))
	return services.SCategory.DeleteTag(uint(id))
}

// PutCategoryTag
// @Tags 博客后台接口
// @Summary 更新标签
// @Description 更新标签
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id formData integer true "标签id"
// @Param name formData string true "标签名称"
// @Success 200 {string} string "success"
// @Failure 400 {object} cerrors.Errors "error"
// @Router /admin/category/tag [put]
func (c *CategoryController) PutCategoryTag() error {
	id, _ := strconv.Atoi(c.Ctx.FormValue("id"))
	name := c.Ctx.FormValue("name")
	return services.SCategory.UpdateTag(uint(id), name)
}

// GetCategoryTags
// @Tags 博客后台接口
// @Summary 查询标签列表
// @Description 查询标签列表
// @Accept json
// @Produce json
// @Success 200 {object} []datamodels.Category "标签列表"
// @Failure 400 {object} cerrors.Errors "error"
// @Router /admin/category/tags [get]
func (c *CategoryController) GetCategoryTags() []datamodels.Category {
	return services.SCategory.GetTagList()
}

// GetCategoryBindartcount
// @Tags 博客后台接口
// @Summary 查询绑定的文章条数
// @Description 查询绑定的文章条数
// @Accept json
// @Produce plain
// @Param id query integer true "分类id"
// @Success 200 {integer} integer "绑定文章条数"
// @Failure 400 {object} cerrors.Errors "error"
// @Router /admin/category/bindartcount [get]
func (c *CategoryController) GetCategoryBindartcount() uint {
	id, _ := strconv.Atoi(c.Ctx.FormValue("id"))
	return services.SCategory.GetBindArtCount(uint(id))
}
