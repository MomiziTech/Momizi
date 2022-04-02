/*
 * @Author: McPlus
 * @Date: 2022-03-24 20:37:42
 * @LastEditTime: 2022-04-02 21:53:48
 * @LastEditors: NyanCatda
 * @Description: Js插件
 * @FilePath: \Momizi\Internal\Plugin\JavaScriptV8\JavaScript.go
 */
package JavascriptV8

import (
	"io/ioutil"
	"strings"

	"github.com/MomiziTech/Momizi/Internal/Controller"
	"github.com/MomiziTech/Momizi/Internal/MessageReceiving/MessageStruct"
	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScriptV8/Events"
	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScriptV8/Listener"
	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScriptV8/Tools"
	"github.com/MomiziTech/Momizi/Tools/Log"
	"rogchap.com/v8go"

	FileFunc "github.com/MomiziTech/Momizi/Tools/File"
)

/**
 * @description: 执行消息监听器
 * @param {MessageStruct.MessageStruct} Message
 * @return {*}
 */
func ExecutionMessageListener(Message MessageStruct.MessageStruct) error {
	// 执行消息监听器
	err := Events.HandleMessageEvent(Message)
	if err != nil {
		return err
	}
	// 执行命令监听器
	err = Events.HandleCommandEvent(Message)
	if err != nil {
		return err
	}

	return nil
}

/**
 * @description: 初始化JavaScript插件
 * @param {*}
 * @return {error} 错误信息
 */
func InitJavaScriptPlugin() error {
	// 从文件中读取插件
	Files, err := ioutil.ReadDir(Controller.PluginPath + "/")
	if err != nil {
		return err
	}

	// 注册虚拟机
	Isolate, Error := v8go.NewIsolate()
	if Error != nil {
		return Error
	}

	// 遍历插件
	for _, File := range Files {
		FileName := File.Name()
		if strings.HasSuffix(FileName, ".momizi.js") {
			// 注册虚拟机
			Context, err := v8go.NewContext(Isolate)
			if err != nil {
				return Error
			}

			Global := Context.Global()

			// 函数注册
			// 监听器初始化
			Global.Set("Listener", Listener.InitListener(Isolate, Context))

			Tools.Register(Isolate, Context)

			ScriptBuffer, err := ioutil.ReadFile(Controller.PluginPath + "/" + FileName)
			if err != nil {
				return err
			}
			Script := string(ScriptBuffer)

			_, err = Context.RunScript(Script, FileName)

			// 打印插件信息
			PluginName, _ := Context.RunScript("PLUGIN_NAME", FileName)
			PluginVersion, _ := Context.RunScript("PLUGIN_VERSION", FileName)
			PluginAuthor, _ := Context.RunScript("PLUGIN_AUTHOR", FileName)
			Log.Info("Plugin", "Loaded <"+PluginName.String()+">", PluginVersion.String(), PluginAuthor.String())

			if err != nil {
				e := err.(*v8go.JSError)
				Log.Print(PluginName.String(), Log.ERROR, e.Message)
				Log.Print(PluginName.String(), Log.ERROR, e.Location)
				Log.Print(PluginName.String(), Log.ERROR, e.StackTrace)
			}

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
