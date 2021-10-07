/*
   Copyright (c) [2021] IT.SOS
   study-notes is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package services

import (
	"encoding/base64"
	"mime/multipart"
	"os"
	"testing"
)

func Test_filesService_GetFileListByAid(t *testing.T) {
	t.Run("upload file", func(t *testing.T) {
		png := "/tmp/john.png"
		pngData := "iVBORw0KGgoAAAANSUhEUgAAAAIAAAACCAYAAABytg0kAAAAFElEQVQImWMQFRX9Lyoq+p8BxgAAJ04E+dJh/q8AAAAASUVORK5CYII="
		f, _ := os.Create(png)
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {

			}
		}(f)
		decoded, _ := base64.StdEncoding.DecodeString(pngData)
		_, err := f.Write(decoded)
		if err != nil {
			return
		}
		file := &multipart.FileHeader{
			Filename: "a.txt",
			Size:     int64(len(decoded)),
		}

		t.Log(SFiles.UploadFile(file))
	})

	t.Run("read file", func(t *testing.T) {
		t.Log(SFiles.GetFile("2021928/RR4lNC5CK4zXm52p.png"))
		//t.Log(router.TypeByFilename("helloworld.txt"))
		//SFiles.GetFileListByAid(1)
	})

	t.Run("file list", func(t *testing.T) {
		t.Log(SFiles.GetFileListByAid(1))
	})
}
