/*
 * @Author: NyanCatda
 * @Date: 2022-03-20 20:40:12
 * @LastEditTime: 2022-03-22 23:47:03
 * @LastEditors: NyanCatda
 * @Description: JavaScript插件加载
 * @FilePath: \Momizi\Internal\Plugin\JavaScript\JavaScript.go
 */
package JavaScript

import (
	"io/ioutil"
	"strings"

	"github.com/MomiziTech/Momizi/Internal/Controller"
	"github.com/MomiziTech/Momizi/Internal/MessageReceiving/MessageStruct"
	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScript/EventListeners"
	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScript/Tools"
	"github.com/MomiziTech/Momizi/Tools/Log"
	"github.com/dop251/goja"

	FileFunc "github.com/MomiziTech/Momizi/Tools/File"
)

var (
	Programs []*goja.Program // 插件内容
)

/**
 * @description: 执行消息监听器
 * @param {MessageStruct.MessageStruct} Message
 * @return {*}
 */
func ExecutionMessageListener(Message MessageStruct.MessageStruct) error {
	// 执行消息监听器
	EventListeners.MessageListenerHandle(Message)

	return nil
}

/**
 * @description: 初始化插件
 * @param {*}
 * @return {error} 错误信息
 */
func InitJavaScriptPlugin() error {
	// 从文件中读取插件
	Files, err := ioutil.ReadDir(Controller.PluginPath + "/")
	if err != nil {
		return err
	}
	// 遍历插件
	for _, File := range Files {
		FileName := File.Name()
		if strings.HasSuffix(FileName, ".momizi.js") {
			// 注册虚拟机
			VM, err := RegistrationVM()
			if err != nil {
				return err
			}

			ScriptBuffer, err := ioutil.ReadFile(Controller.PluginPath + "/" + FileName)
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
			Log.Print("Plugin", Log.INFO, "Loaded <"+PluginName.String()+">", PluginVersion.String(), PluginAuthor.String())

			// 创建插件数据文件夹与配置文件夹
			if _, err := FileFunc.MKDir(Controller.DataPath + "/" + PluginName.String() + "/"); err != nil {
				return err
			}
			if _, err := FileFunc.MKDir(Controller.PluginPath + "/" + PluginName.String() + "/"); err != nil {
				return err
			}
		}
	}

	return nil
}

/**
 * @description: 注册虚拟机
 * @param {*}
 * @return {*}
 */
func RegistrationVM() (*goja.Runtime, error) {
	// 初始化加载器
	VM := goja.New()

	// 注册监听器函数
	if err := EventListeners.Listeners(VM); err != nil {
		return nil, err
	}

	// 注册工具函数
	if err := Tools.Tools(VM); err != nil {
		return nil, err
	}

	return VM, nil
}
