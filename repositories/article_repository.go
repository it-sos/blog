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
	IS_STATE_PRIVATE = 1
	IS_STATE_PUBLIC  = 2
)

type UserRepository interface {
	// Insert 新增
	Insert(p *datamodels.Article) (id uint)
	// Update 更新
	Update(p *datamodels.Article) (id uint)
	// Select 查询用户详细
	Select(p *datamodels.Article) (datamodels.Article, bool)
	// SelectMany 查询用户列表
	SelectMany(p *datamodels.Article, offset int, limit int) (results []datamodels.Article)
}

type userRepository struct {
}

var err error

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (ur *userRepository) Select(p *datamodels.Article) (datamodels.Article, bool) {
	has, err := db.Conn.Get(p)
	if err != nil {
		panic(err)
	}
	return *p, has
}

func (ur *userRepository) SelectMany(p *datamodels.Article, offset int, limit int) (results []datamodels.Article) {
	article := make([]datamodels.Article, 0)
	state := []int{IS_STATE_PUBLIC}
	err := db.Conn.In("is_state", state).Desc("update_time").Limit(limit, offset).Find(&article)
	if err != nil {
		panic(err)
	}
	return article
}

func (ur *userRepository) Insert(p *datamodels.Article) (id uint) {
	_, err = db.Conn.Insert(p)
	if err != nil {
		panic(err)
	}
	return p.Id
}

func (ur *userRepository) Update(p *datamodels.Article) (id uint) {
	_, err = db.Conn.ID(p.Id).Update(p)
	if err != nil {
		panic(err)
	}
	return p.Id
}
