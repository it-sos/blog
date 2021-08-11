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

const (
	// CategoryRelTypeTag 标签
	CategoryRelTypeTag uint8 = 1
	// CategoryRelTypeTopic 专题
	CategoryRelTypeTopic uint8 = 2
)

type CategoryRepository interface {
	// Insert 新增
	Insert(p *datamodels.Category) (id uint)
	// Update 更新
	Update(p *datamodels.Category) (id uint)
	// Select 查询单条
	Select(p *datamodels.Category) (datamodels.Category, bool)
	// SelectMany 查询多条
	SelectMany() []datamodels.Category
}

type categoryRepository struct {
	db mysql.GoLibMysql
}

func (c categoryRepository) Insert(p *datamodels.Category) uint {
	_, err := c.db.Insert(p)
	if err != nil {
		panic(err)
	}
	return p.Id
}

func (c categoryRepository) Update(p *datamodels.Category) uint {
	_, err := c.db.Update(p)
	if err != nil {
		panic(err)
	}
	return p.Id
}

func (c categoryRepository) Select(p *datamodels.Category) (datamodels.Category, bool) {
	has, err := c.db.Get(p)
	if err != nil {
		panic(err)
	}
	return *p, has
}

func (c categoryRepository) SelectMany() (results []datamodels.Category) {
	category := make([]datamodels.Category, 0)
	err := c.db.Desc("id").Find(&category)
	if err != nil {
		panic(err)
	}
	return category
}

var RCategory CategoryRepository = &categoryRepository{mysql.NewMysql()}
