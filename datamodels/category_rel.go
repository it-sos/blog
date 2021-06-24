/*
   Copyright (c) [2021] IT.SOS
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package datamodels

import "time"

// CategoryRel 分类与文章关联表
type CategoryRel struct {
	Id    uint      `json:"id" readonly:"true" example:"1" xorm:"notnull autoincr pk id"` // 关联表自增ID
	Type  uint8     `json:"type" xorm:"tinyint notnull comment('1标签；2专题')"`               // 类型1标签；2专题
	Aid   uint      `json:"aid" xorm:"notnull unique('aid_cid') comment('文章id')"`         // 文章id
	Cid   uint      `json:"cid" xorm:"notnull unique('aid_cid') comment('分类id')"`         // 分类id
	Ctime time.Time `json:"ctime" readonly:"true" xorm:"notnull created"`
}
