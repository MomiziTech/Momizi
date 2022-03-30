/*
 * @Author: NyanCatda
 * @Date: 2022-03-30 19:49:57
 * @LastEditTime: 2022-03-30 19:50:22
 * @LastEditors: NyanCatda
 * @Description: 启动消息接收器
 * @FilePath: \Momizi\Internal\MessageReceiving\ReceivingStart\ReceivingStart.go
 */
package ReceivingStart

import "github.com/MomiziTech/Momizi/Internal/MessageReceiving/WebHook/Start"

/**
 * @description: 启动消息接收器
 * @param {string} RunMode 运行模式
 * @return {*}
 */
func Run(RunMode string) {
	// 启动WebHook接收
	Start.Run(RunMode)
}
