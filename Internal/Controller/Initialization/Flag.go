/*
 * @Author: NyanCatda
 * @Date: 2022-03-30 19:31:35
 * @LastEditTime: 2022-03-30 19:36:49
 * @LastEditors: NyanCatda
 * @Description: 参数解析
 * @FilePath: \Momizi\Internal\Controller\Initialization\Flag.go
 */
package Initialization

import (
	"flag"

	"github.com/MomiziTech/Momizi/Internal/Controller"
)

type FlagConfig struct {
	RunMode    string
	ConfigPath string
	ColorPrint string
}

/**
 * @description: 参数解析
 * @param {*}
 * @return {*}
 */
func Flag() FlagConfig {
	RunMode := flag.String("Mode", "Release", "运行模式")
	ConfigPath := flag.String("Config", Controller.ConfigPath, "指定配置文件路径")
	ColorPrint := flag.String("ColorPrint", "true", "是否输出彩色文本") // 这个地方如果使用flag.Bool()会出现异常，如果默认值为true则无论如何无法修改为false，原因不明
	flag.Parse()

	FlagConfig := FlagConfig{
		RunMode:    *RunMode,
		ConfigPath: *ConfigPath,
		ColorPrint: *ColorPrint,
	}
	return FlagConfig
}
