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
	"bytes"
	"context"
	"fmt"
	minio2 "gitee.com/itsos/golibs/v2/db/minio"
	"gitee.com/itsos/golibs/v2/utils"
	"gitee.com/itsos/studynotes/datamodels"
	"gitee.com/itsos/studynotes/errors"
	"gitee.com/itsos/studynotes/models/vo"
	"gitee.com/itsos/studynotes/repositories"
	"github.com/minio/minio-go/v7"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"path/filepath"
	"time"
)

const (
	bucketName = "media"
	location   = "bj-1"
)

// FilesService 标记为「后台」的方法按登录状态下的逻辑处理
type FilesService interface {
	// GetFileListByAid 获取文章id下文件列表
	GetFileListByAid(aid uint) (files []datamodels.Files, err error)
	// GetFile 获取文件数据
	GetFile(fileName string) ([]byte, string, error)
	// RemoveFile 移除文件
	RemoveFile(fileName string) error
	// UploadFile 上传文件
	UploadFile(aid uint, file *multipart.FileHeader) (fileVO vo.FileVO, err error)
	// GetStorageName 获取存储地址
	GetStorageName(ext string) string
}

type filesService struct {
	file repositories.FilesRepository
}

func (f filesService) GetFile(fileName string) ([]byte, string, error) {
	ctx := context.Background()
	object, err := minio2.NewMinio().GetObject(ctx, bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		log.Fatalln(err)
		return nil, "", errors.Error("read_file_err")
	}
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, object)
	return buf.Bytes(), mime.TypeByExtension(filepath.Ext(fileName)), nil
}

func (f filesService) GetStorageName(ext string) string {
	year, month, day := time.Now().Date()
	return fmt.Sprintf("%d%d%d/%s%s", year, month, day, utils.Rand(16), ext)
}

func (f filesService) GetFileListByAid(aid uint) (files []datamodels.Files, err error) {
	return f.file.SelectManyByAid(aid), nil
}

func (f filesService) RemoveFile(fileName string) error {
	err := minio2.NewMinio().RemoveObject(context.Background(), bucketName, fileName, minio.RemoveObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	} else {
		f.file.Delete(fileName)
	}
	return err
}

func (f filesService) UploadFile(aid uint, file *multipart.FileHeader) (fileVO vo.FileVO, err error) {
	// 存储地址
	fileName := f.GetStorageName(filepath.Ext(file.Filename))
	ctx := context.Background()
	err = minio2.NewMinio().MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minio2.NewMinio().BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
	fo, _ := file.Open()
	_, err = minio2.NewMinio().PutObject(ctx, bucketName, fileName, fo, file.Size, minio.PutObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	} else {
		if f.file.Insert(&datamodels.Files{
			Aid:  aid,
			Name: file.Filename,
			File: fileName,
		}) < 1 {
			f.RemoveFile(fileName)
		} else {
			fileVO = vo.FileVO{
				FileMedia: fileName,
				FileName:  file.Filename,
			}
		}
	}
	return
}

var SFiles FilesService = &filesService{
	file: repositories.RFiles,
}
