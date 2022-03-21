/*
 * @Author: NyanCatda
 * @Date: 2022-03-20 21:24:44
 * @LastEditTime: 2022-03-21 08:47:39
 * @LastEditors: Please set LastEditors
 * @Description: 消息监听器
 * @FilePath: \Momizi\Controller\Plugin\JavaScript\EventListeners\MessageEvent.go
 */
package EventListeners

import (
	"github.com/MomiziTech/Momizi/Controller/MessageReceiving/MessageStruct"
	"github.com/dop251/goja"
)

type MessageListener struct {
	Listener Listener
	FuncName string
	Func     goja.Callable
}

var MessageListenerList []MessageListener

/**
 * @description: 消息监听器
 * @param {*goja.Runtime} VM 加载器
 * @param {MessageStruct.MessageStruct} Message 消息结构体
 * @return {*}
 */
func (L Listener) Message(FuncName string, Func goja.Callable) {
	MessageListenerList = append(MessageListenerList, MessageListener{Listener: L, FuncName: FuncName, Func: Func})
}

func MessageListenerHandle(Message MessageStruct.MessageStruct) {
	for i := 0; i < len(MessageListenerList); i++ {
		MessageListener := MessageListenerList[i]
		Listener := MessageListener.Listener
		if Message.ID != "" {
			switch MessageListener.FuncName {
			case "AllMessage":
				MessageListener.Func(nil, Listener.VM.ToValue(Message))
			case "UserMessage":
				if Message.Type == "User" {
					MessageListener.Func(nil, Listener.VM.ToValue(Message))
				}
			case "GroupMessage":
				if Message.Type == "Group" {
					MessageListener.Func(nil, Listener.VM.ToValue(Message))
				}
			}
		}
	}
}
