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
	"gitee.com/itsos/studynotes/models/vo"
)

type ArticleService interface {
	// GetTop50 获取前访问前50的文章列表
	GetTop50(id string) []vo.ArticleVO
	// GetListPage 获取最新文章列表
	GetListPage(isLogin bool, offset int, limit int) []vo.ArticlesVO
	// GetContent 获取文章详情
	GetContent(id string) vo.ArticleContentVO
}

func NewArticleService() ArticleService {
	return &articleService{}
}

type articleService struct{}

func (a articleService) GetTop50(id string) []vo.ArticleVO {
	panic("implement me")
}

func (a articleService) GetListPage(isLogin bool, offset int, limit int) []vo.ArticlesVO {
	panic("implement me")
}

func (a articleService) GetContent(id string) vo.ArticleContentVO {
	panic("implement me")
}
