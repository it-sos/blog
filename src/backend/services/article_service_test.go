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
	"gitee.com/itsos/golibs/v2/db/mysql"
	"gitee.com/itsos/blog/models/vo"
	"gitee.com/itsos/blog/repositories"
	"testing"
)

func Test_articleService_GetArticleAndContent(t *testing.T) {
	Test_articleService_SaveArticle(t)
	type args struct {
		id uint
	}
	type dataBased struct {
		name    string
		args    args
		wantId  uint
		wantErr bool
	}
	var tests []dataBased
	tests = append(tests, dataBased{
		"查询文章及详情 [后台]",
		args{id: 1},
		1,
		false,
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInfo, err := SArticle.GetArticleAndContent(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArticleAndContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantId != gotInfo.Id {
				t.Errorf("GetArticleAndContent() gotInfo = %v, want id= %d", gotInfo, tt.wantId)
			}
		})
	}
}

func Test_articleService_GetContent(t *testing.T) {
	Test_articleService_SaveArticle(t)
	type args struct {
		isLogin bool
		title   string
	}
	type dataBased struct {
		name    string
		args    args
		wantId  uint
		wantErr bool
	}
	var tests []dataBased
	tests = append(tests, dataBased{
		"查询文章及详情 [前台]",
		args{isLogin: true, title: "新标题"},
		1,
		false,
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SArticle.GetContent(tt.args.isLogin, tt.args.title); got.Article.Article.Id != tt.wantId {
				t.Errorf("GetContent() = %v, wantId %d", got, tt.wantId)
			}
		})
	}
}

func Test_articleService_GetListPage(t *testing.T) {
	Test_articleService_SaveArticle(t)
	type args struct {
		isLogin bool
		page    int
		size    int
		keyword string
	}
	type dataBased struct {
		name    string
		args    args
		wantId  uint
		wantErr bool
	}
	var tests []dataBased
	tests = append(tests, dataBased{
		"查询文章及详情 [前台]",
		args{isLogin: true, page: 1, size: 10, keyword: ""},
		1,
		false,
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SArticle.GetListPage(tt.args.isLogin, tt.args.page, tt.args.size, tt.args.keyword); tt.wantId != got[0].Article.Id {
				t.Errorf("GetListPage() = %v, want %d", got, tt.wantId)
			}
		})
	}
}

func Test_articleService_GetRank(t *testing.T) {
	Test_articleService_SaveArticle(t)
	type args struct {
		isLogin bool
	}
	type dataBased struct {
		name    string
		args    args
		want    string
		wantErr bool
	}
	var tests []dataBased
	tests = append(tests, dataBased{
		"查询排行 [前台]",
		args{isLogin: true},
		"新标题",
		false,
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SArticle.GetRank(tt.args.isLogin); got[0].Title != tt.want {
				t.Errorf("GetRank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_articleService_SaveArticle(t *testing.T) {
	mysql.NewMysql().Exec("truncate table article")
	mysql.NewMysql().Exec("truncate table article_content")
	type args struct {
		vo vo.ArticleParamsVO
	}
	type baseData struct {
		name    string
		args    args
		wantId  uint
		wantErr bool
	}
	var tests []baseData
	tests = append(tests, baseData{
		name: "添加文章",
		args: args{vo.ArticleParamsVO{
			0,
			"新标题",
			"新简介",
			"新内容",
			repositories.IsStatePublic,
			repositories.IsPlaintext,
		}},
		wantId:  1,
		wantErr: false,
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, err := SArticle.SaveArticle(tt.args.vo)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveArticle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("SaveArticle() gotId = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func Test_articleService_DeleteArticle(t *testing.T) {
	Test_articleService_SaveArticle(t)
	type args struct {
		id uint
	}
	type baseData struct {
		name    string
		args    args
		wantErr bool
	}
	var tests []baseData
	tests = append(tests, baseData{
		"文章删除",
		args{1},
		false,
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SArticle.DeleteArticle(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteArticle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
