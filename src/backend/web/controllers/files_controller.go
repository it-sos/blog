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
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"time"
)

type FilesController struct {
	Ctx       iris.Context
	StartTime time.Time
	Sess      *sessions.Session
}

// GetFile
// @Tags 博客前台接口
// @Summary 输出文件流
// @Description 输出文件流图片or视频or其他
// @Accept json
// @Produce json
// @Param media query string true "资源标志"
// @Success 200 {string} string "文件流"
// @Failure 400 {object} errors.Errors "error"
// @Router /file [get]
func (c *FilesController) GetFile() (mvc.Result, error) {
	media := c.Ctx.URLParam("media")
	data, contentType, err := services.SFiles.GetFile(media)
	return mvc.Response{
		ContentType: contentType,
		Content:     data,
	}, err
}
