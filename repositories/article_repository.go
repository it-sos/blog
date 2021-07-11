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
	"strings"
	"time"
)

const (
	IsStatePrivate = 1
	IsStatePublic  = 2

	// NotDeleted 删除状态
	NotDeleted = 0
	IsDeleted  = 1
)

type ArticleRepository interface {
	// Insert 新增
	Insert(p *datamodels.Article) (id uint)
	// Update 更新
	Update(p *datamodels.Article) (id uint)
	// Select 查询文章详细
	Select(state []uint8, p *datamodels.Article) (datamodels.Article, bool)
	// Navigation 按时间续查询上一页下一页导航
	Navigation(state []uint8, utime time.Time) (prevTitle, nextTitle string)
	// SelectMany 查询文章列表
	SelectMany(state []uint8, offset int, limit int) (results []datamodels.Article)
	SelectManyByIds(state []uint8, ids []string) []datamodels.Article
}

type articleRepository struct {
}

func (ur *articleRepository) Navigation(state []uint8, utime time.Time) (prevTitle, nextTitle string) {
	uptime := utime.Format("2006-01-02 15:04:05")
	prev := &datamodels.Article{}
	var (
		has bool
		err error
	)
	has, err = db.Conn.In("is_state", state).Where("utime>? and is_del=?", uptime, NotDeleted).Asc("utime").Get(prev)
	if err != nil {
		panic(err)
	}
	if has {
		prevTitle = prev.Title
	}

	next := &datamodels.Article{}
	has, err = db.Conn.In("is_state", state).Where("utime<? and is_del=?", uptime, NotDeleted).Desc("utime").Get(next)
	if err != nil {
		panic(err)
	}
	if has {
		nextTitle = next.Title
	}
	return
}

// Select 查询文章信息
func (ur *articleRepository) Select(state []uint8, p *datamodels.Article) (datamodels.Article, bool) {
	has, err := db.Conn.In("is_state", state).Where("is_del=?", NotDeleted).Get(p)
	if err != nil {
		panic(err)
	}
	return *p, has
}

// SelectMany 查询文章列表
func (ur *articleRepository) SelectMany(state []uint8, offset int, limit int) (results []datamodels.Article) {
	article := make([]datamodels.Article, 0)
	err := db.Conn.In("is_state", state).Where("is_del=?", NotDeleted).Desc("utime").Limit(limit, offset).Find(&article)
	if err != nil {
		panic(err)
	}
	return article
}

// SelectManyByIds 通过文章id查询列表
func (ur *articleRepository) SelectManyByIds(state []uint8, ids []string) []datamodels.Article {
	article := make([]datamodels.Article, 0)
	orderBy := strings.Join(ids, ",")
	err := db.Conn.In("id", ids).In("is_state", state).Where("is_del=?", NotDeleted).OrderBy("field(id," + orderBy + ")").Find(&article)
	if err != nil {
		panic(err)
	}
	return article
}

func (ur *articleRepository) Insert(p *datamodels.Article) (id uint) {
	_, err := db.Conn.Insert(p)
	if err != nil {
		panic(err)
	}
	return p.Id
}

func (ur *articleRepository) Update(p *datamodels.Article) (id uint) {
	_, err := db.Conn.ID(p.Id).Update(p)
	if err != nil {
		panic(err)
	}
	return p.Id
}

var RArticle ArticleRepository = &articleRepository{}
