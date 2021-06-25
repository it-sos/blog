/*
   Copyright (c) [2021] IT.SOS
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package services

import (
	"gitee.com/itsos/golibs/utils"
	"gitee.com/itsos/golibs/utils/validate"
	"gitee.com/itsos/studynotes/caches"
	"gitee.com/itsos/studynotes/models/vo"
	"gitee.com/itsos/studynotes/repositories"
)

type ArticleService interface {
	// GetRank 获取前访问前50的文章列表
	GetRank() []vo.ArticleVO
	// GetListPage 获取最新文章列表
	GetListPage(isLogin bool, page int, size int) []vo.ArticlesVO
	// GetContent 获取文章详情
	GetContent(id string) vo.ArticleContentVO
}

func NewArticleService() ArticleService {
	return &articleService{repositories.NewArticleRepository(), caches.NAccessTimes}
}

type articleService struct {
	repo  repositories.ArticleRepository
	times caches.AccessTimes
}

func (a articleService) GetRank() []vo.ArticleVO {
	// 获取访问量的前50条
	rank := a.times.Rank(50)
	article := a.repo.SelectManyByIds(rank)
	articleVO := make([]vo.ArticleVO, 0)
	if len(article) > 0 {
		for _, v := range article {
			articleVO = append(articleVO, vo.ArticleVO{Article: v, AccessTimes: a.times.Id(v.Id).Get()})
		}
	}
	return articleVO
}

func (a articleService) GetListPage(isLogin bool, page int, size int) []vo.ArticlesVO {
	page = validate.IntRange(page, 0, 100000)
	size = validate.IntRange(size, 0, 100000)
	var state = []uint8{repositories.IsStatePublic}
	if isLogin {
		state = append(state, repositories.IsStatePrivate)
	}
	offset := (page - 1) * size
	article := a.repo.SelectMany(state, offset, size)
	articleVO := make([]vo.ArticlesVO, 0)
	if len(article) > 0 {
		for _, v := range article {
			articleVO = append(articleVO, vo.ArticlesVO{Article: v, Duration: utils.TimeDuration(v.Utime)})
		}
	}
	return articleVO
}

func (a articleService) GetContent(id string) vo.ArticleContentVO {
	panic("implement me")
}
