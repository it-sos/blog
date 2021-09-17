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

// Category 分类信息表
type Category struct {
	Id    uint      `json:"id" readonly:"true" example:"1" xorm:"notnull autoincr pk id"` // 分类表ID
	Name  string    `json:"name" xorm:"varchar(128) index notnull comment('专题/标签名称')"`    // 专题/标签名称
	Type  uint8     `json:"type" xorm:"tinyint notnull comment('1标签；2专题')"`               // 类型1标签；2专题
	Info  string    `json:"info" xorm:"varchar(255) notnull default '' comment('简介')"`    // 简介
	Pid   uint      `json:"pid" xorm:"notnull index default '0' comment('父编号，预留字段')"`     // 父编号
	Utime time.Time `json:"utime" readonly:"true" xorm:"notnull updated"`
	Ctime time.Time `json:"ctime" readonly:"true" xorm:"notnull created"`
}
