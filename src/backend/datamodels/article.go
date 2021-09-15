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

// Article 文章信息表
type Article struct {
	Id         uint      `json:"id" readonly:"true" example:"1" xorm:"notnull autoincr pk id"`       // 文章表ID
	Title      string    `json:"title" xorm:"varchar(128) index notnull comment('标题')"`              // 标题
	TitleMatch string    `json:"title_match" xorm:"-"`                                               // 匹配标题
	Intro      string    `json:"intro" xorm:"varchar(255) notnull default '' comment('简介')"`         // 简介
	IntroMatch string    `json:"intro_match" xorm:"-"`                                               // 匹配简介
	Uid        uint      `json:"uid" xorm:"notnull comment('编辑者id')"`                                // 编辑者id
	IsState    uint8     `json:"is_state" xorm:"tinyint notnull default '1' comment('状态1私密；2公开')"`   // 状态1私密；2公开
	IsEncrypt  uint8     `json:"is_encrypt" xorm:"tinyint notnull default '1' comment('状态1明文；2密文')"` // 状态1明文；2密文
	IsDel      uint8     `json:"is_del" xorm:"tinyint notnull default '1' comment('状态1未删除；2已删除')"`   // 状态1未删除；2已删除
	Utime      time.Time `json:"utime" readonly:"true" xorm:"index notnull updated"`
	Ctime      time.Time `json:"ctime" readonly:"true" xorm:"notnull created"`
}
