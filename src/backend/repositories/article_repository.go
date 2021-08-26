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
	"strings"
	"time"
)

const (
	// IsStatePrivate 私有
	IsStatePrivate = 1
	// IsStatePublic 公开
	IsStatePublic = 2

	// IsEncrypt 密文
	IsEncrypt = 1
	// IsPlaintext 明文
	IsPlaintext = 2

	// NotDeleted 删除状态
	NotDeleted = 0
	IsDeleted  = 1
)

type ArticleRepository interface {
	// InsertTrans 新增文章与内容
	InsertTrans(p *datamodels.Article, content string) (id uint)
	// UpdateTrans 更新文章与内容
	UpdateTrans(id uint, p *datamodels.Article, content string)
	// DeleteTrans 删除文章与内容
	DeleteTrans(id uint)
	// SoftDelete 软删除
	SoftDelete(id uint)
	// TitleExists title存在性校验
	TitleExists(title string) bool
	// GetInfoById 通过id查询文章信息
	GetInfoById(id uint) (datamodels.Article, bool)
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
	db mysql.GoLibMysql
}

func (ur *articleRepository) SoftDelete(id uint) {
	_, err := ur.db.ID(id).Update(datamodels.Article{IsDel: IsDeleted})
	if err != nil {
		panic(err)
	}
}

func (ur *articleRepository) GetInfoById(id uint) (datamodels.Article, bool) {
	article := new(datamodels.Article)
	has, err := ur.db.ID(id).Get(article)
	if err != nil {
		panic(err)
	}
	return *article, has
}

func (ur *articleRepository) TitleExists(title string) bool {
	article := datamodels.Article{Title: title}
	isExits, err := ur.db.Exist(&article)
	if err != nil {
		panic(err)
	}
	return isExits
}

func (ur *articleRepository) DeleteTrans(id uint) {
	session := ur.db.NewSession()
	defer session.Close()
	err := session.Begin()
	_, err = session.ID(id).Delete(new(datamodels.ArticleContent))
	if err == nil {
		_, err = session.ID(id).Delete(new(datamodels.Article))
	}
	if err != nil {
		session.Rollback()
		panic(err)
	}
	err = session.Commit()
	if err != nil {
		panic(err)
	}
}

func (ur *articleRepository) InsertTrans(p *datamodels.Article, content string) (id uint) {
	session := ur.db.NewSession()
	defer session.Close()
	err := session.Begin()
	_, err = session.Insert(p)
	if err != nil {
		session.Rollback()
		panic(err)
	}

	_, err = session.Insert(&datamodels.ArticleContent{
		Aid:  p.Id,
		Data: content,
	})
	if err != nil {
		session.Rollback()
		panic(err)
	}
	err = session.Commit()
	if err != nil {
		panic(err)
	}
	return p.Id
}

func (ur *articleRepository) UpdateTrans(id uint, p *datamodels.Article, content string) {
	session := ur.db.NewSession()
	defer session.Close()
	err := session.Begin()
	_, err = session.ID(id).Update(p)
	if err != nil {
		session.Rollback()
		panic(err)
	}

	_, err = session.ID(id).Update(&datamodels.ArticleContent{Data: content})
	if err != nil {
		session.Rollback()
		panic(err)
	}
	err = session.Commit()
	if err != nil {
		panic(err)
	}
}

func (ur *articleRepository) Navigation(state []uint8, utime time.Time) (prevTitle, nextTitle string) {
	uptime := utime.Format("2006-01-02 15:04:05")
	prev := &datamodels.Article{}
	var (
		has bool
		err error
	)
	has, err = ur.db.In("is_state", state).Where("utime>? and is_del=?", uptime, NotDeleted).Asc("utime").Get(prev)
	if err != nil {
		panic(err)
	}
	if has {
		prevTitle = prev.Title
	}

	next := &datamodels.Article{}
	has, err = ur.db.In("is_state", state).Where("utime<? and is_del=?", uptime, NotDeleted).Desc("utime").Get(next)
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
	has, err := ur.db.In("is_state", state).Where("is_del=?", NotDeleted).Get(p)
	if err != nil {
		panic(err)
	}
	return *p, has
}

// SelectMany 查询文章列表
func (ur *articleRepository) SelectMany(state []uint8, offset int, limit int) (results []datamodels.Article) {
	article := make([]datamodels.Article, 0)
	err := ur.db.In("is_state", state).Where("is_del=?", NotDeleted).Desc("utime").Limit(limit, offset).Find(&article)
	if err != nil {
		panic(err)
	}
	return article
}

// SelectManyByIds 通过文章id查询列表
func (ur *articleRepository) SelectManyByIds(state []uint8, ids []string) []datamodels.Article {
	article := make([]datamodels.Article, 0)
	orderBy := strings.Join(ids, ",")
	err := ur.db.In("id", ids).In("is_state", state).Where("is_del=?", NotDeleted).OrderBy("field(id," + orderBy + ")").Find(&article)
	if err != nil {
		panic(err)
	}
	return article
}

func (ur *articleRepository) Insert(p *datamodels.Article) (id uint) {
	_, err := ur.db.Insert(p)
	if err != nil {
		panic(err)
	}
	return p.Id
}

func (ur *articleRepository) Update(p *datamodels.Article) (id uint) {
	_, err := ur.db.ID(p.Id).Update(p)
	if err != nil {
		panic(err)
	}
	return p.Id
}

var RArticle ArticleRepository = &articleRepository{mysql.NewMysql()}
