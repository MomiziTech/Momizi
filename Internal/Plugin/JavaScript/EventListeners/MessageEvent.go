/*
 * @Author: NyanCatda
 * @Date: 2022-03-20 21:24:44
 * @LastEditTime: 2022-03-23 14:57:21
 * @LastEditors: NyanCatda
 * @Description: 消息监听器
 * @FilePath: \Momizi\Internal\Plugin\JavaScript\EventListeners\MessageEvent.go
 */
package EventListeners

import (
	"github.com/MomiziTech/Momizi/Internal/MessageReceiving/MessageStruct"
	"github.com/dop251/goja"
)

type MessageListener struct {
	Listener Listener
	FuncName string
	Callback goja.Callable
}

var MessageListenerList []MessageListener

/**
 * @description: 消息监听器注册
 * @param {*goja.Runtime} VM 加载器
 * @param {MessageStruct.MessageStruct} Message 消息结构体
 * @return {*}
 */
func (L Listener) Message(FuncName string, Callback goja.Callable) {
	MessageListenerList = append(MessageListenerList, MessageListener{Listener: L, FuncName: FuncName, Callback: Callback})
}

/**
 * @description: 消息监听器调用
 * @param {MessageStruct.MessageStruct} Message 消息结构体
 * @return {MessageStruct.MessageStruct} 消息结构体
 */
func MessageListenerHandle(Message MessageStruct.MessageStruct) {
	for _, MessageListener := range MessageListenerList {
		Listener := MessageListener.Listener
		if Message.ID != "" {
			switch MessageListener.FuncName {
			case "AllMessage":
				MessageListener.Callback(nil, Listener.VM.ToValue(Message))
			case "UserMessage":
				if Message.Type == "User" {
					MessageListener.Callback(nil, Listener.VM.ToValue(Message))
				}
			case "GroupMessage":
				if Message.Type == "Group" {
					MessageListener.Callback(nil, Listener.VM.ToValue(Message))
				}
			}
		}
	}
}
