/*
 * @Author: NyanCatda
 * @Date: 2022-03-23 21:01:14
 * @LastEditTime: 2022-03-23 21:30:17
 * @LastEditors: NyanCatda
 * @Description: 读取配置项
 * @FilePath: \Momizi\Internal\Plugin\JavaScript\Tools\Config\Config.go
 */
package Config

import (
	"github.com/MomiziTech/Momizi/Internal/Controller"
	"github.com/MomiziTech/Momizi/Internal/Plugin/Tools/JsonConfig"
	"github.com/MomiziTech/Momizi/Tools/Log"
	"github.com/dop251/goja"
)

var PluginPath = Controller.PluginPath + "/"

type Config struct {
	VM *goja.Runtime
}

/**
 * @description: 注册配置文件函数
 * @param {*goja.Runtime} VM 实例
 * @return {error} 错误信息
 */
func RegistrationFunction(VM *goja.Runtime) error {
	return VM.Set("Config", Config{VM: VM})
}

/**
 * @description: 初始化配置文件
 * @param {string} Path 配置文件路径
 * @param {map[string]any} Default 默认配置
 * @return {bool} 是否初始化成功
 */
func (Config Config) Init(Path string, Default map[string]any) bool {
	PluginName := Config.VM.Get("PLUGIN_NAME").String()
	conf := JsonConfig.New(PluginPath+PluginName+"/"+Path, Default)
	if err := conf.Init(); err != nil {
		Log.Error("Plugin", err)
		return false
	}
	return true
}

/**
 * @description: 读取配置项
 * @param {string} Path 配置文件路径
 * @param {map[string]any} Default 默认配置
 * @param {string} Name 配置项名称
 * @return {*}
 */
func (Config Config) Get(Path string, Default map[string]any, Name string) any {
	PluginName := Config.VM.Get("PLUGIN_NAME").String()
	conf := JsonConfig.New(PluginPath+PluginName+"/"+Path, Default)
	Body, err := conf.Get(Name)
	if err != nil {
		Log.Error("Plugin", err)
		return Default[Name]
	}
	return Body
}

/**
 * @description: 设置配置项
 * @param {string} Path 配置文件路径
 * @param {map[string]any} Default 默认配置
 * @param {string} Name 配置项名称
 * @param {any} Value 配置项值
 * @return {bool} 是否设置成功
 */
func (Config Config) Set(Path string, Default map[string]any, Name string, Value any) bool {
	PluginName := Config.VM.Get("PLUGIN_NAME").String()
	conf := JsonConfig.New(PluginPath+PluginName+"/"+Path, Default)
	if err := conf.Set(Name, Value); err != nil {
		Log.Error("Plugin", err)
		return false
	}
	return true
}

/**
 * @description: 删除配置项
 * @param {string} Path 配置文件路径
 * @param {map[string]any} Default 默认配置
 * @param {string} Name 配置项名称
 * @return {bool} 是否删除成功
 */
func (Config Config) Delete(Path string, Default map[string]any, Name string) bool {
	PluginName := Config.VM.Get("PLUGIN_NAME").String()
	conf := JsonConfig.New(PluginPath+PluginName+"/"+Path, Default)
	if err := conf.Delete(Name); err != nil {
		Log.Error("Plugin", err)
		return false
	}
	return true
}
