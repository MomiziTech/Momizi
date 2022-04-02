/*
 * @Author: NyanCatda
 * @Date: 2022-04-02 14:25:18
 * @LastEditTime: 2022-04-02 21:54:28
 * @LastEditors: NyanCatda
 * @Description: CommandEvent
 * @FilePath: \Momizi\Internal\Plugin\JavaScriptV8\Events\Command.go
 */
package Events

import (
	"encoding/json"
	"strings"

	"github.com/MomiziTech/Momizi/Internal/MessageReceiving/MessageStruct"
	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScriptV8/Tools/Loader"
	"rogchap.com/v8go"
)

type CommandEventCallback struct {
	Command  string
	CallBack *v8go.Function
	Context  *v8go.Context
}

var (
	CommandEventCallbacks []CommandEventCallback
)

/**
 * @description: 命令监听函数
 * @param {*v8go.Isolate} Isolate 虚拟机
 * @param {*v8go.Context} Context 上下文
 * @return {*}
 */
func InitCommandEvent(Isolate *v8go.Isolate, Context *v8go.Context) *v8go.FunctionTemplate {
	Message, _ := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		Command := Info.Args()[0]
		CallBack := Info.Args()[1]

		if Command.IsString() && CallBack.IsFunction() {
			CallBack, _ := CallBack.AsFunction()
			CommandEventCallbacks = append(CommandEventCallbacks, CommandEventCallback{Command: Command.String(), CallBack: CallBack, Context: Context})
		}
		return nil
	})

	return Message
}

/**
 * @description: 调用命令监听函数
 * @param {MessageStruct.MessageStruct} Message 消息结构体
 * @return {*}
 */
func HandleCommandEvent(Message MessageStruct.MessageStruct) error {
	Json, err := json.Marshal(Message)
	if err != nil {
		return err
	}
	for _, EventCallback := range CommandEventCallbacks {
		if len(Message.MessageChain) > 0 {
			if Message.MessageChain[0].Type == "Text" {
				// 提取命令
				CommandArray := strings.Fields(Message.MessageChain[0].Text)
				if len(CommandArray) > 0 {
					// 如果包含被注册的命令
					if CommandArray[0] == EventCallback.Command {
						// 处理命令
						var CommandParameters []string
						// 修剪数组
						Index := 0
						CommandParameters = append(CommandArray[:Index], CommandArray[Index+1:]...)
						// 转换为v8go.Value
						CommandParametersValue, err := Loader.GoArrayToV8Object(EventCallback.Context, CommandParameters)
						if err != nil {
							return err
						}

						// 处理消息结构体
						Message, err := v8go.JSONParse(EventCallback.Context, string(Json))
						if err != nil {
							return err
						}

						// 传递回调(CommandParametersValue命令参数, Message消息结构体)
						if _, err := EventCallback.CallBack.Call(CommandParametersValue, Message); err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}
