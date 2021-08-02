/*
   Copyright (c) [2021] IT.SOS
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package main

import (
	"gitee.com/itsos/golibs/v2/db/mysql"
	"gitee.com/itsos/studynotes/datamodels"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if viper.Get("use-driver") == "sqlite3" {
		if _, e := os.Stat("sqlite3.db"); !os.IsExist(e) {
			log.Print("remove sqlite3.db")
			os.Remove("sqlite3.db")
		}
	}
	tables := []interface{}{
		new(datamodels.Article),
		new(datamodels.ArticleContent),
		new(datamodels.Category),
		new(datamodels.CategoryRel),
	}
	err := mysql.NewMysql().DropTables(tables...)
	if err != nil {
		panic(err)
		return
	}
	err = mysql.NewMysql().Sync2(tables...)
	if err != nil {
		panic(err)
		return
	}
}
