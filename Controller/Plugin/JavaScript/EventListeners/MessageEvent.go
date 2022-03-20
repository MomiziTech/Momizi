/*
 * @Author: NyanCatda
 * @Date: 2022-03-20 21:24:44
 * @LastEditTime: 2022-03-21 00:28:03
 * @LastEditors: NyanCatda
 * @Description: 消息监听器
 * @FilePath: \Momizi\Controller\Plugin\JavaScript\EventListeners\MessageEvent.go
 */
package EventListeners

import (
	"github.com/dop251/goja"
)

/**
 * @description: 消息监听器
 * @param {*goja.Runtime} VM 加载器
 * @param {MessageStruct.MessageStruct} Message 消息结构体
 * @return {*}
 */
func (L Listener) Message(FuncName string, Func goja.Callable) error {
	if L.MessageStruct.ID != "" {
		switch FuncName {
		case "AllMessage":
			_, err := Func(nil, L.VM.ToValue(L.MessageStruct))
			return err
		case "UserMessage":
			if L.MessageStruct.Type == "User" {
				_, err := Func(nil, L.VM.ToValue(L.MessageStruct))
				return err
			}
		case "GroupMessage":
			if L.MessageStruct.Type == "Group" {
				_, err := Func(nil, L.VM.ToValue(L.MessageStruct))
				return err
			}
		}
	}
	return nil
}
