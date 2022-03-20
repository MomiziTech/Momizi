/*
 * @Author: NyanCatda
 * @Date: 2022-03-20 21:24:44
 * @LastEditTime: 2022-03-20 22:41:23
 * @LastEditors: NyanCatda
 * @Description: 消息监听器
 * @FilePath: \Momizi\Controller\Plugin\JavaScript\EventListeners\MessageEvent.go
 */
package EventListeners

import (
	"github.com/MomiziTech/Momizi/Controller/MessageReceiving/MessageStruct"
	"github.com/dop251/goja"
)

/**
 * @description: 消息监听器
 * @param {*goja.Runtime} VM 加载器
 * @param {MessageStruct.MessageStruct} Message 消息结构体
 * @return {*}
 */
func MessageEvent(VM *goja.Runtime, Message MessageStruct.MessageStruct) error {
	// 全部消息监听
	err := VM.Set("MessageListeners", func(Func goja.Callable) {
		Func(nil, VM.ToValue(Message))
	})
	if err != nil {
		return err
	}

	// 用户消息监听
	err = VM.Set("UserMessageListeners", func(Func goja.Callable) {
		if Message.Type == "User" {
			Func(nil, VM.ToValue(Message))
		}
	})
	if err != nil {
		return err
	}

	// 群消息监听
	err = VM.Set("GroupMessageListeners", func(Func goja.Callable) {
		if Message.Type == "Group" {
			Func(nil, VM.ToValue(Message))
		}
	})
	if err != nil {
		return err
	}

	return nil
}
