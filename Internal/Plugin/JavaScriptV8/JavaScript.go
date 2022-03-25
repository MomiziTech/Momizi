/*
 * @Author: McPlus
 * @Date: 2022-03-24 20:37:42
 * @LastEditTime: 2022-03-25 19:56:04
 * @LastEditors: McPlus
 * @Description: Js插件
 * @FilePath: \Momizi\Internal\Plugin\JavaScriptV8\JavaScript.go
 */
package JavascriptV8

import (
	"io/ioutil"
	"strings"

	"github.com/MomiziTech/Momizi/Internal/Controller"
	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScriptV8/Listeners"
	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScriptV8/Tools"
	"github.com/MomiziTech/Momizi/Tools/Log"
	"rogchap.com/v8go"

	FileFunc "github.com/MomiziTech/Momizi/Tools/File"
)

func InitJavaScriptPlugin() error {
	// 从文件中读取插件
	Files, err := ioutil.ReadDir(Controller.PluginPath + "/")
	if err != nil {
		return err
	}

	// 注册虚拟机
	Isolate, _ := v8go.NewIsolate()

	// 遍历插件
	for _, File := range Files {
		FileName := File.Name()
		if strings.HasSuffix(FileName, ".momizi.js") {
			// 注册虚拟机
			Context, _ := v8go.NewContext(Isolate)

			Global := Context.Global()

			// 函数注册
			// 监听器初始化
			Global.Set("Listeners", Listeners.InitListeners(Isolate, Context))
			
			Tools.Register(Isolate, Context)

			ScriptBuffer, err := ioutil.ReadFile(Controller.PluginPath + "/" + FileName)
			if err != nil {
				return err
			}
			Script := string(ScriptBuffer)

			Context.RunScript(Script, FileName)

			// 打印插件信息
			PluginName, _ := Context.RunScript("PLUGIN_NAME", FileName)
			PluginVersion, _ := Context.RunScript("PLUGIN_VERSION", FileName)
			PluginAuthor, _ := Context.RunScript("PLUGIN_AUTHOR", FileName)
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
