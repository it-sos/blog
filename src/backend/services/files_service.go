/*
   Copyright (c) [2021] IT.SOS
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package services

import (
	"gitee.com/itsos/golibs/v2/global/variable"
	"gitee.com/itsos/studynotes/datamodels"
	"gitee.com/itsos/studynotes/repositories"
)

// FilesService 标记为「后台」的方法按登录状态下的逻辑处理
type FilesService interface {
	// GetFileListByAid 获取文章id下文件列表
	GetFileListByAid(aid uint) (files []datamodels.Files, err error)
	// RemoveFile 移除文件
	RemoveFile(id uint) error
	// UploadFile 上传文件
	UploadFile(aid uint, name, addr string) (file string, err error)
	// GetStorage 获取存储地址
	GetStorage(aid uint) string
}

type filesService struct {
	file repositories.FilesRepository
}

func (f filesService) GetStorage(aid uint) string {
	return variable.BasePath + "/web/public/photos"
}

func (f filesService) GetFileListByAid(aid uint) (files []datamodels.Files, err error) {
	panic("implement me")
}

func (f filesService) RemoveFile(id uint) error {
	panic("implement me")
}

func (f filesService) UploadFile(aid uint, name, addr string) (file string, err error) {
	panic("implement me")
}

var SFiles FilesService = &filesService{
	file: repositories.RFiles,
}
