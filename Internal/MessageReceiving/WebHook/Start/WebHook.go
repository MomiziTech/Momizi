/*
 * @Author: NyanCatda
 * @Date: 2022-03-30 19:40:25
 * @LastEditTime: 2022-03-30 19:52:38
 * @LastEditors: NyanCatda
 * @Description: WebHook处理
 * @FilePath: \Momizi\Internal\MessageReceiving\WebHook\Start\WebHook.go
 */
package Start

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/MomiziTech/Momizi/Internal/MessageReceiving"
	"github.com/MomiziTech/Momizi/Tools/Log"
	"github.com/MomiziTech/Momizi/Tools/ReadConfig"
	"github.com/gin-gonic/gin"
)

/**
 * @description: 启动WebHook接收
 * @param {string} RunMode
 * @return {*}
 */
func Run(RunMode string) error {
	// 初始化Gin
	if RunMode != "Dev" {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	}
	r := gin.Default()

	Config := ReadConfig.GetConfig

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
		Log.Error("System", err)
		return err
	}
	return nil
}
