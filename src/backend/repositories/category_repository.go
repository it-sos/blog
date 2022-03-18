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
	"gitee.com/itsos/blog/datamodels"
)

const (
	// CategoryTypeTag 标签
	CategoryTypeTag uint8 = 1
	// CategoryTypeTopic 专题
	CategoryTypeTopic uint8 = 2
)

type CategoryRepository interface {
	// Insert 新增
	Insert(p *datamodels.Category) (id uint)
	// Delete 删除
	Delete(id uint) bool
	// Update 更新
	Update(id uint, p *datamodels.Category) bool
	// ExistName 名称存在性校验
	ExistName(name string, cyType uint8) bool
	// Select 查询单条
	Select(p *datamodels.Category) (datamodels.Category, bool)
	// SelectMany 查询多条
	SelectMany() []datamodels.Category
	// SelectManyByType 通过类型查询多条
	SelectManyByType(cyType uint8) []datamodels.Category
}

type categoryRepository struct {
	db mysql.GoLibMysql
}

func (c categoryRepository) ExistName(name string, cyType uint8) bool {
	article := datamodels.Category{Name: name, Type: cyType}
	isExits, err := c.db.Exist(&article)
	if err != nil {
		panic(err)
	}
	return isExits
}

func (c categoryRepository) Delete(id uint) bool {
	affected, err := c.db.ID(id).Delete(new(datamodels.Category))
	if err != nil {
		panic(err)
	}
	return affected > 0
}

func (c categoryRepository) Insert(p *datamodels.Category) uint {
	_, err := c.db.Insert(p)
	if err != nil {
		panic(err)
	}
	return p.Id
}

func (c categoryRepository) Update(id uint, p *datamodels.Category) bool {
	affected, err := c.db.ID(id).Update(p)
	if err != nil {
		panic(err)
	}
	return affected > 0
}

func (c categoryRepository) Select(p *datamodels.Category) (datamodels.Category, bool) {
	has, err := c.db.Get(p)
	if err != nil {
		panic(err)
	}
	return *p, has
}

func (c categoryRepository) SelectMany() (results []datamodels.Category) {
	results = make([]datamodels.Category, 0)
	err := c.db.Desc("id").Find(&results)
	if err != nil {
		panic(err)
	}
	return
}

func (c categoryRepository) SelectManyByType(cyType uint8) (results []datamodels.Category) {
	results = make([]datamodels.Category, 0)
	err := c.db.Where("type=?", cyType).Desc("id").Find(&results)
	if err != nil {
		panic(err)
	}
	return
}

var RCategory CategoryRepository = &categoryRepository{mysql.NewMysql()}
