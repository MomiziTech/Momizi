/*
 * @Author: McPlus
 * @Date: 2022-03-24 20:58:57
 * @LastEditTime: 2022-03-25 23:56:01
 * @LastEditors: McPlus
 * @Description: MessageEvent
 * @FilePath: \Momizi\Internal\Plugin\JavaScriptV8\Events\Message.go
 */
package Events

import (
	"encoding/json"

	"github.com/MomiziTech/Momizi/Internal/MessageReceiving/MessageStruct"
	"rogchap.com/v8go"
)

type EventCallback struct {
	FuncName string
	CallBack *v8go.Function
}

var EventCallbacks []EventCallback

/**
 * @description: 初始化消息监听函数
 * @param {*v8go.Isolate} Isolate 虚拟机
 * @return {*v8go.FunctionTemplate}
 */
func InitMessageEvent(Isolate *v8go.Isolate) *v8go.FunctionTemplate {
	Message, _ := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		FuncName := Info.Args()[0]
		CallBack := Info.Args()[1]

		if FuncName.IsString() && CallBack.IsFunction() {
			CallBack, _ := CallBack.AsFunction()
			EventCallbacks = append(EventCallbacks, EventCallback{FuncName: FuncName.String(), CallBack: CallBack})
		}
		return nil
	})

	return Message
}

/**
 * @description: 调用消息监听函数
 * @param {MessageStruct.MessageStruct} Message 消息结构体
 * @return {*}
 */
func HandleMessageEvent(Message MessageStruct.MessageStruct) error {
	Json, err := json.Marshal(Message)
	if err != nil {
		return err
	}
	Context, err := v8go.NewContext()
	if err != nil {
		return err
	}

	Object, err := v8go.JSONParse(Context, string(Json))
	if err != nil {
		return err
	}
	for _, EventCallback := range EventCallbacks {
		if Message.ID != "" {
			switch EventCallback.FuncName {
			case "AllMessage":
				if _, err := EventCallback.CallBack.Call(Object); err != nil {
					return err
				}
			case "UserMessage":
				if Message.Type == "User" {
					if _, err := EventCallback.CallBack.Call(Object); err != nil {
						return err
					}
				}
			case "GroupMessage":
				if Message.Type == "Group" {
					if _, err := EventCallback.CallBack.Call(Object); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}