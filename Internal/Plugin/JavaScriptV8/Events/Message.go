/*
 * @Author: McPlus
 * @Date: 2022-03-24 20:58:57
 * @LastEditTime: 2022-03-24 21:45:52
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

func HandleMessageEvent(Message MessageStruct.MessageStruct) {
	Json, _ := json.Marshal(Message)
	Object, _ := v8go.JSONParse(nil, string(Json))
	for _, EventCallback := range EventCallbacks {
		if Message.ID != "" {
			switch EventCallback.FuncName {
			case "AllMessage":
				EventCallback.CallBack.Call(Object)
			case "UserMessage":
				if Message.Type == "User" {
					EventCallback.CallBack.Call(Object)
				}
			case "GroupMessage":
				if Message.Type == "Group" {
					EventCallback.CallBack.Call(Object)
				}
			}
		}
	}
}
