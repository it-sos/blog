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
	"gitee.com/itsos/golibs/v2/db/mysql"
	"gitee.com/itsos/studynotes/datamodels"
)

type CategoryRelRepository interface {
	// Insert 新增
	Insert(p *datamodels.CategoryRel) bool
	// Update 更新
	Update(p *datamodels.CategoryRel) bool
	// Select 查询单条
	Select(p *datamodels.CategoryRel) (datamodels.CategoryRel, bool)
	// SelectManyByAid 通过文章id查询多条
	SelectManyByAid(aid uint) []datamodels.CategoryRel
	// SelectManyByCid 通过分类id查询多条
	SelectManyByCid(cid uint) []datamodels.CategoryRel
}

type categoryRelRepository struct {
	db mysql.GoLibMysql
}

func (c categoryRelRepository) Insert(p *datamodels.CategoryRel) bool {
	affected, err := c.db.Insert(p)
	if err != nil {
		panic(err)
	}
	return affected > 0
}

func (c categoryRelRepository) Update(p *datamodels.CategoryRel) bool {
	affected, err := c.db.Update(p)
	if err != nil {
		panic(err)
	}
	return affected > 0
}

func (c categoryRelRepository) Select(p *datamodels.CategoryRel) (datamodels.CategoryRel, bool) {
	has, err := c.db.Get(p)
	if err != nil {
		panic(err)
	}
	return *p, has
}

func (c categoryRelRepository) SelectManyByAid(aid uint) (results []datamodels.CategoryRel) {
	categoryRel := make([]datamodels.CategoryRel, 0)
	err := c.db.Where("aid=?", aid).Desc("id").Find(&categoryRel)
	if err != nil {
		panic(err)
	}
	return categoryRel
}

func (c categoryRelRepository) SelectManyByCid(cid uint) (results []datamodels.CategoryRel) {
	categoryRel := make([]datamodels.CategoryRel, 0)
	err := c.db.Where("cid=?", cid).Desc("id").Find(&categoryRel)
	if err != nil {
		panic(err)
	}
	return categoryRel
}

var RCategoryRel CategoryRelRepository = &categoryRelRepository{mysql.NewMysql()}
