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
	"github.com/kataras/iris/v12/sessions"
	"time"
)

type FilesController struct {
	Ctx       iris.Context
	StartTime time.Time
	Sess      *sessions.Session
}

// DeleteFiles
// @Tags 博客后台接口
// @Summary 删除文章
// @Description 删除文章
// @Accept json
// @Produce json
// @Param id query integer true "文章id"
// @Success 200 {string} string ""
// @Failure 400 {object} errors.Errors "error"
// @Router /admin/article [delete]
func (c *FilesController) DeleteFiles() {
	//id, _ := strconv.Atoi(c.Ctx.FormValue("id"))
	//services.SFiles.DeleteFiles(uint(id))
}

// PostFiles
// @Tags 博客后台接口
// @Summary 保存文件
// @Description 保存文章
// @Accept  multipart/form-data
// @Produce plain
// @Param file formData file true "request file data"
// @Param aid query integer true "文章id"
// @Success 200 {integer} integer "文章id"
// @Failure 400 {object} errors.Errors "error"
// @Router /admin/article [post]
func (c *FilesController) PostFiles() (file string, err error) {
	return "", nil
	//body, _ := c.Ctx.GetBody()
	//article := new(vo.FilesParamsVO)
	//if err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&article); err != nil {
	//	return
	//}
	//return services.SFiles.SaveFiles(*article)
}

// GetFiles
// @Tags 博客后台接口
// @Summary 查询文件列表
// @Description 查询文件列表
// @Accept json
// @Produce json
// @Param aid query integer true "文章id"
// @Success 200 {object} vo.FilesEditVO "文章详情VO"
// @Failure 400 {object} errors.Errors "error"
// @Router /admin/article [get]
func (c *FilesController) GetFiles() ([]datamodels.Files, error) {
	aid, _ := c.Ctx.URLParamInt("aid")
	return services.SFiles.GetFileListByAid(uint(aid))
}
