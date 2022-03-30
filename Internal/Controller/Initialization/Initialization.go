/*
 * @Author: NyanCatda
 * @Date: 2022-03-30 19:29:55
 * @LastEditTime: 2022-03-30 19:47:27
 * @LastEditors: NyanCatda
 * @Description: 初始化程序
 * @FilePath: \Momizi\Internal\Controller\Initialization\Initialization.go
 */
package Initialization

import (
	"strconv"

	"github.com/MomiziTech/Momizi/Internal/Controller"
	"github.com/MomiziTech/Momizi/Internal/Plugin"
	"github.com/MomiziTech/Momizi/Tools/File"
	"github.com/MomiziTech/Momizi/Tools/Log"
	"github.com/MomiziTech/Momizi/Tools/ReadConfig"
)

/**
 * @description: 初始化程序
 * @param {*}
 * @return {*}
 */
func Initialization() (FlagConfig, error) {
	// 解析参数
	FlagConfig := Flag()

	// 设置是否输出彩色文本
	Color, err := strconv.ParseBool(FlagConfig.ColorPrint)
	if err != nil {
		return FlagConfig, err
	}
	Log.ColorPrint = Color

	Log.Info("System", "Momizi启动，当前版本："+Controller.Version+"，运行模式："+FlagConfig.RunMode)

	// 设置配置文件路径
	ReadConfig.ConfigPath = FlagConfig.ConfigPath
	// 加载配置文件
	if err := ReadConfig.LoadConfig(); err != nil {
		return FlagConfig, err
	}

	// 初始化日志文件夹
	if _, err := File.MKDir(Controller.LogPath); err != nil {
		return FlagConfig, err
	}
	// 初始化插件文件夹
	if _, err := File.MKDir(Controller.PluginPath); err != nil {
		return FlagConfig, err
	}
	// 初始化数据文件夹
	if _, err := File.MKDir(Controller.DataPath); err != nil {
		return FlagConfig, err
	}

	// 初始化插件
	if err := Plugin.InitPlugin(); err != nil {
		return FlagConfig, err
	}

	return FlagConfig, err
}
