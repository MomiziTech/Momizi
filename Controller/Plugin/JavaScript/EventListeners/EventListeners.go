/*
 * @Author: NyanCatda
 * @Date: 2022-03-20 21:18:09
 * @LastEditTime: 2022-03-20 21:25:01
 * @LastEditors: NyanCatda
 * @Description: 事件监听器
 * @FilePath: \Momizi\Controller\Plugin\JavaScript\EventListeners\EventListeners.go
 */
package EventListeners

import (
	"github.com/MomiziTech/Momizi/Controller/MessageReceiving/MessageStruct"
	"github.com/dop251/goja"
)

/**
 * @description: 监听器注册
 * @param {*goja.Runtime} VM 加载器
 * @param {MessageStruct.MessageStruct} Message 消息结构体
 * @return {*}
 */
func Listeners(VM *goja.Runtime, Message MessageStruct.MessageStruct) error {
	// 注册消息事件
	err := MessageEvent(VM, Message)
	if err != nil {
		return err
	}

	return nil
}
