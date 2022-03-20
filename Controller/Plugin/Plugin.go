/*
 * @Author: NyanCatda
 * @Date: 2022-03-08 21:58:27
 * @LastEditTime: 2022-03-20 21:13:29
 * @LastEditors: NyanCatda
 * @Description: 插件加载模块
 * @FilePath: \Momizi\Controller\Plugin\Plugin.go
 */
package Plugin

import (
	"github.com/MomiziTech/Momizi/Controller/MessageReceiving/MessageStruct"
	"github.com/MomiziTech/Momizi/Controller/Plugin/JavaScript"
	"github.com/MomiziTech/Momizi/Utils/Log"
)

func RunPlugin(Message MessageStruct.MessageStruct) {
	// 运行JavaScript插件
	JavaScript.RunJavaScriptPlugin(Message)
}

func InitPlugin() error {
	// 初始化JavaScript插件
	err := JavaScript.InitJavaScriptPlugin()
	if err != nil {
		Log.ErrorWrite(err)
		return err
	}

	return nil
}
