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

type ArticleContentRepository interface {
	// Insert 新增
	Insert(p *datamodels.ArticleContent) bool
	// Update 更新
	Update(p *datamodels.ArticleContent) bool
	// Select 查询文章内容
	Select(p *datamodels.ArticleContent) (bool, datamodels.ArticleContent)
}

type articleContentRepository struct {
}

func NewArticleContentRepository() ArticleContentRepository {
	return &articleContentRepository{}
}

// Select 查询文章信息
func (ur *articleContentRepository) Select(p *datamodels.ArticleContent) (bool, datamodels.ArticleContent) {
	has, err := db.Conn.Get(p)
	if err != nil {
		panic(err)
	}
	return has, *p
}

func (ur *articleContentRepository) Insert(p *datamodels.ArticleContent) bool {
	affected, err := db.Conn.Insert(p)
	if err != nil {
		panic(err)
	}
	return affected > 0
}

func (ur *articleContentRepository) Update(p *datamodels.ArticleContent) bool {
	affected, err := db.Conn.ID(p.Aid).Update(p)
	if err != nil {
		panic(err)
	}
	return affected > 0
}
