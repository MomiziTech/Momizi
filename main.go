/*
 * @Author: NyanCatda
 * @Date: 2022-03-08 21:19:51
 * @LastEditTime: 2022-03-28 15:45:28
 * @LastEditors: NyanCatda
 * @Description:
 * @FilePath: \Momizi\main.go
 */
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/MomiziTech/Momizi/Internal/Controller"
	"github.com/MomiziTech/Momizi/Internal/MessageReceiving"
	"github.com/MomiziTech/Momizi/Internal/Plugin"
	"github.com/MomiziTech/Momizi/Tools/File"
	"github.com/MomiziTech/Momizi/Tools/Log"
	"github.com/MomiziTech/Momizi/Tools/ReadConfig"
)

/**
 * @description: 主函数错误处理
 * @param {error} Error 错误信息
 * @return {*}
 */
func Error(Error error) {
	fmt.Println(Error.Error())
	key := make([]byte, 1)
	os.Stdin.Read(key)
	os.Exit(1)
}

/**
 * @description: 初始化程序
 * @param {*}
 * @return {*}
 */
func Initialization() error {
	// 初始化日志文件夹
	if _, err := File.MKDir(Controller.LogPath); err != nil {
		return err
	}
	// 初始化插件文件夹
	if _, err := File.MKDir(Controller.PluginPath); err != nil {
		return err
	}
	// 初始化数据文件夹
	if _, err := File.MKDir(Controller.DataPath); err != nil {
		return err
	}
	return nil
}

func main() {
	// 参数解析
	RunMode := flag.String("Mode", "Release", "运行模式")
	ConfigPath := flag.String("config", Controller.ConfigPath, "指定配置文件路径")
	flag.Parse()

	// 初始化程序
	if err := Initialization(); err != nil {
		Error(err)
	}

	// 设置配置文件路径
	ReadConfig.ConfigPath = *ConfigPath
	// 加载配置文件
	if err := ReadConfig.LoadConfig(); err != nil {
		Error(err)
	}
	Config := ReadConfig.GetConfig

	// 初始化插件
	if err := Plugin.InitPlugin(); err != nil {
		Error(err)
	}

	// 初始化Gin
	if *RunMode != "Dev" {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	}
	r := gin.Default()

	Port := Config.Run.WebHook.Port
	WebHookKey := Config.Run.WebHook.Key

	// 注册WebHook接收地址
	r.POST("/"+WebHookKey, func(c *gin.Context) {
		if err := MessageReceiving.MessageReceiving(c); err != nil {
			Log.Error("System", err)
			c.JSONP(http.StatusInternalServerError, gin.H{"success": false, "time": time.Now().Unix()})
		}
		c.JSONP(http.StatusOK, gin.H{"success": true, "time": time.Now().Unix()})
	})

	// 启动WebHook接收
	Log.Info("System", "WebHook接收已启动，地址：http://0.0.0.0:"+Port+"/"+WebHookKey)
	if err := r.Run(":" + Port); err != nil {
		Error(err)
	}
}
