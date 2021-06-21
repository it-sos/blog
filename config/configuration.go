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
	"gitee.com/itsos/golibs/config"
	"github.com/spf13/viper"
	"reflect"
)

type ConfigurationReadOnly interface {
	config.ConfigurationReadOnly
	GetQWidth() int
	GetQHeight() int
	GetQBuffer() int64
	GetQUrl() string
}

type Configuration struct {
	*config.Configuration
	QWidth  string `yaml:"qrcode.width"`
	QHeight string `yaml:"qrcode.height"`
	QBuffer string `yaml:"qrcode.buffer"`
	QUrl    string `yaml:"qrcode.url"`
}

func (c Configuration) GetQWidth() int {
	return viper.GetInt(c.QWidth)
}

func (c Configuration) GetQHeight() int {
	return viper.GetInt(c.QHeight)
}

func (c Configuration) GetQBuffer() int64 {
	return viper.GetInt64(c.QBuffer)
}

func (c Configuration) GetQUrl() string {
	return viper.GetString(c.QUrl)
}

func covertConfiguration() *Configuration {
	c := &Configuration{
		Configuration: config.CovertConfiguration(),
	}
	t := reflect.ValueOf(c).Elem()
	for i := 0; i < t.NumField(); i++ {
		f := t.Type().Field(i)
		if f.Type.Name() != "string" {
			continue
		}
		tag := f.Tag.Get("yaml")
		s := t.FieldByName(f.Name)
		if !s.CanSet() {
			panic("not set value.")
		}
		s.SetString(tag)
	}
	return c
}

var C ConfigurationReadOnly = covertConfiguration()
