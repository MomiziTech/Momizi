/*
 * @Author: NyanCatda
 * @Date: 2022-04-04 12:35:10
 * @LastEditTime: 2022-04-04 12:37:12
 * @LastEditors: NyanCatda
 * @Description: 插件列表模块
 * @FilePath: \Momizi\Internal\Plugin\Tools\PluginList\PluginList.go
 */
package PluginList

type PluginInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Author  string `json:"author"`
}

var PluginList []PluginInfo

/**
 * @description: 添加信息进入插件列表
 * @param {string} Name 插件名称
 * @param {string} Version 插件版本
 * @param {string} Author 插件作者
 * @return {*}
 */
func AddPluginInfo(Name string, Version string, Author string) {
	PluginList = append(PluginList, PluginInfo{
		Name:    Name,
		Version: Version,
		Author:  Author,
	})
}

/**
 * @description: 清理插件列表
 * @param {*}
 * @return {*}
 */
func ClearPluginList() {
	PluginList = nil
}
