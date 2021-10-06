/*
   Copyright (c) [2021] IT.SOS
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package errors

import (
	"encoding/json"
	"errors"
)

type Errors struct {
	Code    int    `json:"code" example:"4000000"`
	Message string `json:"message"`
}

var errCodeList = map[string]Errors{
	"param_err":            {4020001, "参数异常"},
	"article_exists_err":   {4001001, "文章标题已存在"},
	"article_notfound_err": {4001002, "该文章不存在"},
	"article_remove_err":   {4001003, "文章移除失败"},
	"topic_exists_err":     {4002001, "专题已存在"},
	"topic_remove_err":     {4002002, "专题移除失败"},
	"tag_exists_err":       {4003001, "标签已存在"},
	"tag_remove_err":       {4003002, "标签移除失败"},
	"unbind_err":           {4004001, "解绑操作失败"},
	"read_file_err":        {4004002, "读取文件内容失败"},
	"remove_file_err":      {4004003, "删除文件操作失败"},
	"resize_img_err":       {4004004, "重置图片大小失败"},
}

func Error(key string) error {
	ret, _ := json.Marshal(errCodeList[key])
	return errors.New(string(ret))
}
