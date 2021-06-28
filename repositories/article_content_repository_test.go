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
	_ "gitee.com/itsos/golibs/tests/testsdb"
	"gitee.com/itsos/studynotes/datamodels"
	"testing"
)

func TestArticleContentRepository_Select(t *testing.T) {
	t.Log(RArticleContent.Select(&datamodels.ArticleContent{
		Aid: 1,
	}))
}

func TestArticleContentRepository_Update(t *testing.T) {
	t.Log(RArticleContent.Update(&datamodels.ArticleContent{
		Aid: 1,
	}))
}
