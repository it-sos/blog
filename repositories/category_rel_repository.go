/*
   Copyright (c) [2021] IT.SOS
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package repositories

import (
	"gitee.com/itsos/golibs/db"
	"gitee.com/itsos/studynotes/datamodels"
)

const (
	// CategoryRelTypeTag 标签
	CategoryRelTypeTag uint8 = 1
	// CategoryRelTypeTopic 专题
	CategoryRelTypeTopic uint8 = 2
)

type CategoryRelRepository interface {
	// Insert 新增
	Insert(p *datamodels.CategoryRel) bool
	// Update 更新
	Update(p *datamodels.CategoryRel) bool
	// Select 查询单条
	Select(p *datamodels.CategoryRel) (datamodels.CategoryRel, bool)
	// SelectMany 查询多条
	SelectMany(p *datamodels.CategoryRel) []datamodels.CategoryRel
}

type categoryRelRepository struct {
}

func (c categoryRelRepository) Insert(p *datamodels.CategoryRel) bool {
	affected, err := db.Conn.Insert(p)
	if err != nil {
		panic(err)
	}
	return affected > 0
}

func (c categoryRelRepository) Update(p *datamodels.CategoryRel) bool {
	affected, err := db.Conn.Update(p)
	if err != nil {
		panic(err)
	}
	return affected > 0
}

func (c categoryRelRepository) Select(p *datamodels.CategoryRel) (datamodels.CategoryRel, bool) {
	has, err := db.Conn.Get(p)
	if err != nil {
		panic(err)
	}
	return *p, has
}

func (c categoryRelRepository) SelectMany(p *datamodels.CategoryRel) (results []datamodels.CategoryRel) {
	categoryRel := make([]datamodels.CategoryRel, 0)
	err := db.Conn.Where("aid=?", p.Aid).Desc("id").Find(&categoryRel)
	if err != nil {
		panic(err)
	}
	return categoryRel
}

var RCategoryRel CategoryRelRepository = &categoryRelRepository{}
