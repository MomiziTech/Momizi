/*
 * @Author: McPlus
 * @Date: 2022-03-24 20:58:57
 * @LastEditTime: 2022-04-02 14:43:55
 * @LastEditors: NyanCatda
 * @Description: MessageEvent
 * @FilePath: \Momizi\Internal\Plugin\JavaScriptV8\Events\Message.go
 */
package Events

import (
	"encoding/json"

	"github.com/MomiziTech/Momizi/Internal/MessageReceiving/MessageStruct"
	"rogchap.com/v8go"
)

type MessageEventCallback struct {
	FuncName string
	CallBack *v8go.Function
	Context  *v8go.Context
}

var MessageEventCallbacks []MessageEventCallback

/**
 * @description: 初始化消息监听函数
 * @param {*v8go.Isolate} Isolate 虚拟机
 * @return {*v8go.FunctionTemplate}
 */
func InitMessageEvent(Isolate *v8go.Isolate, Context *v8go.Context) *v8go.FunctionTemplate {
	Message, _ := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		FuncName := Info.Args()[0]
		CallBack := Info.Args()[1]

		if FuncName.IsString() && CallBack.IsFunction() {
			CallBack, _ := CallBack.AsFunction()
			MessageEventCallbacks = append(MessageEventCallbacks, MessageEventCallback{FuncName: FuncName.String(), CallBack: CallBack, Context: Context})
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
	for _, EventCallback := range MessageEventCallbacks {
		if Message.ID != "" {
			Value, err := v8go.JSONParse(EventCallback.Context, string(Json))

			if err != nil {
				return err
			}

			switch EventCallback.FuncName {
			case "AllMessage":
				if _, err := EventCallback.CallBack.Call(Value); err != nil {
					return err
				}
			case "UserMessage":
				if Message.Type == "User" {
					if _, err := EventCallback.CallBack.Call(Value); err != nil {
						return err
					}
				}
			case "GroupMessage":
				if Message.Type == "Group" {
					if _, err := EventCallback.CallBack.Call(Value); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
