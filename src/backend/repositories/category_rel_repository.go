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
	"github.com/it-sos/blog/datamodels"
	"github.com/it-sos/golibs/v2/db/mysql"
)

type CategoryRelRepository interface {
	// Insert 新增
	Insert(p *datamodels.CategoryRel) uint
	// Delete 删除单条记录专题或标签与文章的绑定记录
	Delete(id uint) bool
	// DeleteByAid 删除所有与文章有关的专题与标签
	DeleteByAid(aid uint) bool
	// DeleteByAidAndCid 删除所有与文章有关的专题或标签
	DeleteByAidAndCid(aid uint, cid uint) bool
	// DeleteByCid 删除所有与专题或标签有关的文章
	DeleteByCid(cid uint) bool
	// Update 更新
	Update(id uint, p *datamodels.CategoryRel) bool
	// Select 查询单条
	Select(p *datamodels.CategoryRel) (datamodels.CategoryRel, bool)
	// SelectManyByAid 通过文章id查询多条
	SelectManyByAid(aid uint) []datamodels.CategoryRel
	// SelectManyByCid 通过分类id查询多条
	SelectManyByCid(cid uint) []datamodels.CategoryRel
	// GetCountByCid 通过分类id查询绑定的条数
	GetCountByCid(cid uint) uint
}

type categoryRelRepository struct {
	db mysql.GoLibMysql
}

func (c categoryRelRepository) DeleteByAidAndCid(aid uint, cid uint) bool {
	affected, err := c.db.Where("aid=? and cid=?", aid, cid).Delete(new(datamodels.CategoryRel))
	if err != nil {
		panic(err)
	}
	return affected > 0
}

func (c categoryRelRepository) DeleteByAid(aid uint) bool {
	affected, err := c.db.Where("aid=?", aid).Delete(new(datamodels.CategoryRel))
	if err != nil {
		panic(err)
	}
	return affected > 0
}

func (c categoryRelRepository) DeleteByCid(cid uint) bool {
	affected, err := c.db.Where("cid=?", cid).Delete(new(datamodels.CategoryRel))
	if err != nil {
		panic(err)
	}
	return affected > 0
}

func (c categoryRelRepository) Delete(id uint) bool {
	affected, err := c.db.ID(id).Delete(new(datamodels.CategoryRel))
	if err != nil {
		panic(err)
	}
	return affected > 0
}

func (c categoryRelRepository) Insert(p *datamodels.CategoryRel) uint {
	_, err := c.db.Insert(p)
	if err != nil {
		panic(err)
	}
	return p.Id
}

func (c categoryRelRepository) Update(id uint, p *datamodels.CategoryRel) bool {
	affected, err := c.db.ID(id).Update(p)
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

func (c categoryRelRepository) GetCountByCid(cid uint) uint {
	total, err := c.db.Where("cid=?", cid).Count(new(datamodels.CategoryRel))
	if err != nil {
		panic(err)
	}
	return uint(total)
}

var RCategoryRel CategoryRelRepository = &categoryRelRepository{mysql.NewMysql()}
