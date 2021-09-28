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
	"testing"
)

func Test_filesService_GetFileListByAid(t *testing.T) {
	t.Run("upload file", func(t *testing.T) {
		SFiles.UploadFile(1, "a.txt", "/tmp/a.txt")
	})

	t.Run("read file", func(t *testing.T) {
		t.Log(SFiles.GetFile("2021928/RR4lNC5CK4zXm52p.txt"))
		//t.Log(router.TypeByFilename("helloworld.txt"))
		//SFiles.GetFileListByAid(1)
	})

	t.Run("file list", func(t *testing.T) {
		t.Log(SFiles.GetFileListByAid(1))
	})
}
