/*
 * @Author: NyanCatda
 * @Date: 2022-03-08 21:25:31
 * @LastEditTime: 2022-03-08 21:31:34
 * @LastEditors: NyanCatda
 * @Description: 配置文件模块
 * @FilePath: \Momizi\Tools\ReadConfig\ReadConfig.go
 */
package ReadConfig

import (
	"errors"
	"io/ioutil"
	"os"
	"reflect"

	"gopkg.in/yaml.v2"
)

var (
	GetConfig  *Config
	ConfigPath string
)

/**
 * @description: 加载配置文件
 * @param {*}
 * @return {*}
 */
func LoadConfig() error {
	// 检查配置文件是否存在
	_, err := os.Lstat(ConfigPath)
	if err != nil {
		return err
	}

	content, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		return err
	}
	newStu := &Config{}
	err = yaml.Unmarshal(content, &newStu)
	if err != nil {
		return err
	}
	GetConfig = newStu

	if err := GetConfig.CheckConfig(); err != nil {
		return err
	}

	return nil
}

/**
 * @description: 检查配置文件字段是否为空
 * @param {*}
 * @return {error}
 */
func (value *Config) CheckConfig() error {
	val := reflect.ValueOf(value).Elem() // 获取字段值
	typ := reflect.TypeOf(value).Elem()  // 获取字段类型
	// 遍历struct中的字段
	for i := 0; i < typ.NumField(); i++ {
		// 当字段出现空时，输出错误
		if val.Field(i).IsZero() {
			return errors.New(typ.Field(i).Name + "字段为空，请按照说明填写配置文件")
		}
	}
	return nil
}
