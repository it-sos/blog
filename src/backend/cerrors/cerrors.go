/*
   Copyright (c) [2021] IT.SOS
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package cerrors

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Errors struct {
	Code    int    `json:"code" example:"4000000"`
	Message string `json:"message"`
}

func (n *Errors) SetMessage(message string) {
	n.Message = message
}

func (n *Errors) GetMessage() string {
	return n.Message
}

var errCodeList = map[string]Errors{
	"parameter_error":     {4000000, "参数错误[ %s ]"},
	"unauthorized_access": {4000003, "未授权访问"},

	"article_exists_err":   {4001001, "文章标题已存在"},
	"article_notfound_err": {4001002, "该文章不存在"},
	"article_remove_err":   {4001003, "文章移除失败"},
	"search_keyword_err":   {4001004, "搜索关键子不存在"},
	"topic_exists_err":     {4002001, "专题已存在"},
	"topic_remove_err":     {4002002, "专题移除失败"},
	"tag_exists_err":       {4003001, "标签已存在"},
	"tag_remove_err":       {4003002, "标签移除失败"},
	"unbind_err":           {4004001, "解绑操作失败"},
	"read_file_err":        {4004002, "读取文件内容失败"},
	"remove_file_err":      {4004003, "删除文件操作失败"},
	"resize_img_err":       {4004004, "重置图片大小失败"},

	"user_account_exits": {4005001, "用户账号已存在"},
	"user_auth_fail":     {4005002, "账号或密码不正确"},
}

func Error(key string) error {
	ret, _ := json.Marshal(errCodeList[key])
	return errors.New(string(ret))
}

func Errorf(key string, msg interface{}) error {
	e := errCodeList[key]
	e.SetMessage(fmt.Sprintf(e.GetMessage(), msg))
	ret, _ := json.Marshal(e)
	return errors.New(string(ret))
}
