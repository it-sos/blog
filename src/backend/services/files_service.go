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
	"github.com/it-sos/blog/cerrors"
	"github.com/it-sos/blog/datamodels"
	"github.com/it-sos/blog/models/vo"
	"github.com/it-sos/blog/repositories"
	minio2 "github.com/it-sos/golibs/v2/db/minio"
	"github.com/it-sos/golibs/v2/utils"
	"github.com/minio/minio-go/v7"
	"github.com/nfnt/resize"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"strings"
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
	// ResizeImg 设置图片大小
	ResizeImg(img []byte, contentType, size string) (newImg []byte, err error)
	// RemoveFile 移除文件
	RemoveFile(fileName string) error
	// UploadFile 上传文件
	UploadFile(file *multipart.FileHeader) (fileVO vo.FileVO, err error)
	// RelFileAndArticle 关联文章与文件记录
	RelFileAndArticle(aid uint, name, media string) (err error)
	// GetStorageName 获取存储地址
	GetStorageName(ext string) string
}

type filesService struct {
	file repositories.FilesRepository
}

func (f filesService) ResizeImg(img []byte, contentType, size string) (newImg []byte, err error) {
	image, _, err := image.Decode(bytes.NewBuffer(img))
	if err != nil {
		err = cerrors.Error("read_file_err")
		return
	}
	sizes := strings.Split(size, "x")
	w, _ := strconv.Atoi(sizes[0])
	h, _ := strconv.Atoi(sizes[1])
	m := resize.Resize(uint(w), uint(h), image, resize.Lanczos3)
	buf := bytes.NewBuffer(nil)
	if contentType == "image/jpeg" {
		err = jpeg.Encode(buf, m, nil)
	} else if contentType == "image/gif" {
		err = gif.Encode(buf, m, nil)
	} else {
		err = png.Encode(buf, m)
	}
	return buf.Bytes(), nil
}

func (f filesService) RelFileAndArticle(aid uint, name, media string) (err error) {
	if f.file.Insert(&datamodels.Files{
		Aid:  aid,
		Name: name,
		File: media,
	}) < 1 {
		err = f.RemoveFile(media)
	}
	return
}

func (f filesService) GetFile(fileName string) ([]byte, string, error) {
	ctx := context.Background()
	object, err := minio2.NewMinio().GetObject(ctx, bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		log.Print(err)
		return nil, "", cerrors.Error("read_file_err")
	}
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, object)
	return buf.Bytes(), mime.TypeByExtension(filepath.Ext(fileName)), nil
}

func (f filesService) GetStorageName(ext string) string {
	year, month, day := time.Now().Date()
	return fmt.Sprintf("%d%d%d/%s%s", year, month, day, utils.Rand(16, utils.RandMix), ext)
}

func (f filesService) GetFileListByAid(aid uint) (files []datamodels.Files, err error) {
	return f.file.SelectManyByAid(aid), nil
}

func (f filesService) RemoveFile(fileName string) error {
	err := minio2.NewMinio().RemoveObject(context.Background(), bucketName, fileName, minio.RemoveObjectOptions{})
	if err != nil {
		log.Print(err)
		err = cerrors.Error("remove_file_err")
	} else {
		f.file.Delete(fileName)
	}
	return err
}

func (f filesService) UploadFile(file *multipart.FileHeader) (fileVO vo.FileVO, err error) {
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
			log.Print(err)
			return
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
	fo, _ := file.Open()
	_, err = minio2.NewMinio().PutObject(ctx, bucketName, fileName, fo, file.Size, minio.PutObjectOptions{})
	if err != nil {
		log.Print(err)
		return
	} else {
		fileVO = vo.FileVO{
			FileMedia: fileName,
			FileName:  file.Filename,
		}
	}
	return
}

var SFiles FilesService = &filesService{
	file: repositories.RFiles,
}
