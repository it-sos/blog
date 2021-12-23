/*
   Copyright (c) [2021] itsos
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
               http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package config

import (
	"gitee.com/itsos/golibs/v2/config"
	"gitee.com/itsos/golibs/v2/utils/reflects"
	"github.com/spf13/viper"
	"reflect"
)

type ConfigurationReadOnly interface {
	config.ConfigurationReadOnly
	GetDemo() int
}

type Configuration struct {
	*config.Configuration
	Demo string `yaml:"demo.demo"`
}

func (c Configuration) GetDemo() int {
	return viper.GetInt(c.Demo)
}

func covertConfiguration() *Configuration {
	c := &Configuration{
		Configuration: config.Config,
	}
	t := reflect.ValueOf(c).Elem()
	reflects.TagToValueFlip(t, reflects.YAML)
	return c
}

var C ConfigurationReadOnly = covertConfiguration()
