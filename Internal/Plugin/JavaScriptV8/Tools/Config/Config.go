/*
 * @Author: NyanCatda
 * @Date: 2022-04-03 13:23:50
 * @LastEditTime: 2022-04-04 12:55:53
 * @LastEditors: NyanCatda
 * @Description: 配置文件操作类
 * @FilePath: \Momizi\Internal\Plugin\JavaScriptV8\Tools\Config\Config.go
 */
package Config

import (
	"errors"

	"github.com/MomiziTech/Momizi/Internal/Controller"
	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScriptV8/Tools/Loader"
	"github.com/MomiziTech/Momizi/Internal/Plugin/Tools/JsonConfig"
	"github.com/MomiziTech/Momizi/Tools/Log"
	"rogchap.com/v8go"
)

var PluginPath = Controller.PluginPath + "/"

/**
 * @description: 注册配置文件函数
 * @param {*v8go.Isolate} Isolate v8实例
 * @param {*v8go.Context} Context v8上下文
 * @return {*}
 */
func Register(Isolate *v8go.Isolate, Context *v8go.Context) *v8go.Object {
	Config, err := v8go.NewObjectTemplate(Isolate)
	if err != nil {
		return nil
	}

	ConfigConstructor, err := RegisterConstructor(Isolate, Context)
	if err != nil {
		return nil
	}
	Config.Set("Config", ConfigConstructor)

	ConfigInstance, err := Config.NewInstance(Context)
	if err != nil {
		return nil
	}

	return ConfigInstance
}

/**
 * @description: 注册类构造器
 * @param {*v8go.Isolate} Isolate v8实例
 * @param {*v8go.Context} Context v8上下文
 * @return {*}
 */
func RegisterConstructor(Isolate *v8go.Isolate, Context *v8go.Context) (*v8go.FunctionTemplate, error) {
	Constructor, err := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		PluginName, err := Context.Global().Get("PLUGIN_NAME")
		if err != nil {
			Log.Error("Plugin", err)
			return nil
		}
		if len(Info.Args()) < 2 {
			Log.Error(PluginName.String(), errors.New("参数不足"))
			return nil
		}
		Path := Info.Args()[0].String()
		Default := Info.Args()[1].Object()

		// 获取默认配置
		DefaultMap, err := Loader.V8ObjectToGoStringAnyMap(Context, Default)
		if err != nil {
			Log.Error(PluginName.String(), err)
		}

		Config := JsonConfig.New(PluginPath+PluginName.String()+"/"+Path, DefaultMap)
		ConfigObject, err := v8go.NewObjectTemplate(Isolate)
		if err != nil {
			Log.Error(PluginName.String(), err)
		}
		Instance, err := ConfigObject.NewInstance(Info.Context())
		if err != nil {
			Log.Error(PluginName.String(), err)
		}
		Instance.Set("path", Path)
		Instance.Set("default", Default)

		RegisterObjectFunction(Isolate, Context, Instance, Info, Config)

		return Instance.Value
	})

	return Constructor, err
}

/**
 * @description: 注册配置文件对象函数
 * @param {*v8go.Isolate} Isolate v8实例
 * @param {*v8go.Context} Context v8上下文
 * @param {*v8go.Object} Instance v8对象实例
 * @param {*v8go.FunctionCallbackInfo} ObjectInfo v8函数回调信息
 * @param {*JsonConfig.Config} Config 配置文件
 * @return {*}
 */
func RegisterObjectFunction(Isolate *v8go.Isolate, Context *v8go.Context, Instance *v8go.Object, ObjectInfo *v8go.FunctionCallbackInfo, Config *JsonConfig.Config) error {
	PluginName, err := Context.Global().Get("PLUGIN_NAME")
	if err != nil {
		Log.Error("Plugin", err)
		return nil
	}
	/**
	* @description: 初始化配置文件函数
	* @param {*}
	* @return {bool} 是否初始化成功
	 */
	Init, err := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		err := Config.Init()
		var OK bool
		if err == nil {
			OK = true
		} else {
			Log.Error(PluginName.String(), err)
		}

		Value, err := v8go.NewValue(Isolate, OK)
		if err != nil {
			Log.Error(PluginName.String(), err)
		}

		return Value
	})
	if err != nil {
		return err
	}
	Instance.Set("Init", Init.GetFunction(Context))

	/**
	* @description:读取配置项函数
	* @param {string} Key 配置项名称
	* @return {any} Value 配置项值
	 */
	Get, err := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		Key := Info.Args()[0].String()

		ConfigValue, err := Config.Get(Key)
		if err != nil {
			Log.Error(PluginName.String(), err)
		}

		Value, err := v8go.NewValue(Isolate, ConfigValue)
		if err != nil {
			Log.Error(PluginName.String(), err)
		}

		return Value
	})
	if err != nil {
		return err
	}
	Instance.Set("Get", Get.GetFunction(Context))

	/**
	* @description: 写入配置项函数
	* @param {string} Key 配置项名称
	* @param {any} Value 配置项值
	* @return {bool} 是否写入成功
	 */
	Set, err := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		Key := Info.Args()[0].String()
		Value := Info.Args()[1].String()

		err := Config.Set(Key, Value)
		var OK bool
		if err == nil {
			OK = true
		} else {
			Log.Error(PluginName.String(), err)
		}

		v8Value, err := v8go.NewValue(Isolate, OK)
		if err != nil {
			Log.Error(PluginName.String(), err)
		}

		return v8Value
	})
	if err != nil {
		return err
	}
	Instance.Set("Set", Set.GetFunction(Context))

	/**
	* @description: 删除配置项函数
	* @param {string} Key 配置项名称
	* @return {bool} 是否删除成功
	 */
	Delete, err := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		Key := Info.Args()[0].String()

		err := Config.Delete(Key)
		var OK bool
		if err == nil {
			OK = true
		} else {
			Log.Error(PluginName.String(), err)
		}

		v8Value, err := v8go.NewValue(Isolate, OK)
		if err != nil {
			Log.Error(PluginName.String(), err)
		}

		return v8Value
	})
	if err != nil {
		return err
	}
	Instance.Set("Delete", Delete.GetFunction(Context))

	return nil
}
