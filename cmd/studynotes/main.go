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
	"flag"
	"gitee.com/itsos/golibs/global/variable"
	"gitee.com/itsos/studynotes"
)

var workdir = flag.String("w", "", "指定工作目录 -w /workdir")

func main() {
	flag.Parse()
	if *workdir != "" {
		variable.BasePath = *workdir
	}
	println(variable.BasePath)
	studynotes.Listen()
}
