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
	"gitee.com/itsos/studynotes/models/vo"
	"gitee.com/itsos/studynotes/services"
	"github.com/kataras/iris/v12"
	"strconv"
	"time"
)

type FilesController struct {
	Ctx       iris.Context
	StartTime time.Time
}

// DeleteFiles
// @Tags 博客后台接口
// @Summary 删除资源文件
// @Description 删除资源文件
// @Accept json
// @Produce json
// @Param media query string true "文章资源"
// @Success 200 {string} string ""
// @Failure 400 {object} cerrors.Errors "error"
// @Router /admin/files [delete]
func (c *FilesController) DeleteFiles() error {
	media := c.Ctx.URLParam("media")
	return services.SFiles.RemoveFile(media)
}

// PostFiles
// @Tags 博客后台接口
// @Summary 保存文件
// @Description 保存文件
// @Accept  multipart/form-data
// @Produce json
// @Param file formData file true "request file data"
// @Param aid formData integer true "文章id"
// @Success 200 {object} vo.FileVO "文件vo"
// @Failure 400 {object} cerrors.Errors "error"
// @Router /admin/files [post]
func (c *FilesController) PostFiles() (file vo.FileVO, err error) {
	_, f, errs := c.Ctx.FormFile("file")
	if errs != nil {
		err = errs
		return
	}
	return services.SFiles.UploadFile(f)
}

// PostFilesArticle
// @Tags 博客后台接口
// @Summary 关联文件与文章
// @Description 关联文件与文章
// @Accept  multipart/form-data
// @Produce json
// @Param aid formData integer true "文章id"
// @Param name formData string true "文件名称"
// @Param media formData string true "文件标志"
// @Success 200 {string} string ""
// @Failure 400 {object} cerrors.Errors "error"
// @Router /admin/files/article [post]
func (c *FilesController) PostFilesArticle() error {
	name := c.Ctx.FormValue("name")
	media := c.Ctx.FormValue("media")
	aid, _ := strconv.Atoi(c.Ctx.FormValue("aid"))
	return services.SFiles.RelFileAndArticle(uint(aid), name, media)
}

// GetFiles
// @Tags 博客后台接口
// @Summary 查询文件列表
// @Description 查询文件列表
// @Accept json
// @Produce json
// @Param aid query integer true "文章id"
// @Success 200 {object} vo.FilesEditVO "文章详情VO"
// @Failure 400 {object} cerrors.Errors "error"
// @Router /admin/article [get]
func (c *FilesController) GetFiles() ([]datamodels.Files, error) {
	aid, _ := c.Ctx.URLParamInt("aid")
	return services.SFiles.GetFileListByAid(uint(aid))
}
