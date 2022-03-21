/*
 * @Author: NyanCatda
 * @Date: 2022-03-20 20:40:12
 * @LastEditTime: 2022-03-21 08:32:11
 * @LastEditors: Please set LastEditors
 * @Description: JavaScript插件加载
 * @FilePath: \Momizi\Controller\Plugin\JavaScript\JavaScript.go
 */
package JavaScript

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/MomiziTech/Momizi/Controller/Plugin/JavaScript/EventListeners"
	"github.com/MomiziTech/Momizi/Controller/Plugin/JavaScript/Tools"
	"github.com/dop251/goja"
)

var (
	Programs []*goja.Program // 插件内容
)

/**
 * @description: 运行插件
 * @param {*} MessageStruct 消息结构体
 * @return {*}
 */
func RunJavaScriptPlugin() error {
	// 初始化加载器
	VM := goja.New()

	// 注册函数
	if err := RegistrationFunction(VM); err != nil {
		return err
	}

	// 遍历数组获取预编译的插件
	for _, Program := range Programs {
		// 执行插件
		_, err := VM.RunProgram(Program)
		if err != nil {
			return err
		}
	}

	return nil
}

/**
 * @description: 初始化插件
 * @param {*}
 * @return {error} 错误信息
 */
func InitJavaScriptPlugin() error {
	// 初始化加载器
	VM := goja.New()

	// 注册函数
	if err := RegistrationFunction(VM); err != nil {
		return err
	}

	// 从文件中读取插件
	Files, err := ioutil.ReadDir("./plugins/")
	if err != nil {
		return err
	}
	// 遍历插件
	for _, File := range Files {
		FileName := File.Name()
		if strings.HasSuffix(FileName, ".momizi.js") {
			ScriptBuffer, err := ioutil.ReadFile("./plugins/" + FileName)
			if err != nil {
				return err
			}
			Script := string(ScriptBuffer)

			// 预编译插件
			Program, err := goja.Compile(FileName, Script, false)
			if err != nil {
				return err
			}

			// 将预编译后的插件写入缓存
			Programs = append(Programs, Program)

			// 初始化插件
			_, err = VM.RunProgram(Program)
			if err != nil {
				return err
			}

			// 打印插件信息
			PluginName := VM.Get("PLUGIN_NAME")
			PluginVersion := VM.Get("PLUGIN_VERSION")
			PluginAuthor := VM.Get("PLUGIN_AUTHOR")
			fmt.Println("[插件加载]", PluginName.String(), PluginVersion.String(), PluginAuthor.String())
		}
	}

	fmt.Println("插件加载完毕！")

	return nil
}

func RegistrationFunction(VM *goja.Runtime) error {
	// 注册监听器函数
	if err := EventListeners.Listeners(VM); err != nil {
		return err
	}
	// 注册工具函数
	if err := Tools.Tools(VM); err != nil {
		return err
	}

	return nil
}
