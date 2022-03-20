package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/MomiziTech/Momizi/Controller/MessageReceiving"
	"github.com/MomiziTech/Momizi/Controller/Plugin"
	"github.com/MomiziTech/Momizi/Utils"
	"github.com/MomiziTech/Momizi/Utils/Log"
	"github.com/MomiziTech/Momizi/Utils/ReadConfig"
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
	if _, err := Utils.PathExists("./logs"); err != nil {
		return err
	}
	// 初始化插件文件夹
	if _, err := Utils.PathExists("./plugins"); err != nil {
		return err
	}
	// 初始化数据文件夹
	if _, err := Utils.PathExists("./data"); err != nil {
		return err
	}
	return nil
}

func main() {
	// 参数解析
	ConfigPath := flag.String("config", "./config.yml", "指定配置文件路径")
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
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	Port := Config.Run.WebHook.Port
	WebHookKey := Config.Run.WebHook.Key

	// 注册WebHook接收地址
	r.POST("/"+WebHookKey, func(c *gin.Context) {
		if err := MessageReceiving.MessageReceiving(c); err != nil {
			Log.ErrorWrite(err)
		}
	})

	// 启动WebHook接收
	fmt.Println("WebHook接收已启动，地址：http://0.0.0.0:" + Port + "/" + WebHookKey)
	if err := r.Run(":" + Port); err != nil {
		Error(err)
	}
}
