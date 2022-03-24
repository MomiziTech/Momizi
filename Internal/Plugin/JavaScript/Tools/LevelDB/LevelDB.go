/*
 * @Author: NyanCatda
 * @Date: 2022-03-24 14:28:26
 * @LastEditTime: 2022-03-24 15:11:16
 * @LastEditors: NyanCatda
 * @Description: LevelDB函数封装
 * @FilePath: \Momizi\Internal\Plugin\JavaScript\Tools\LevelDB\LevelDB.go
 */
package LevelDB

import (
	"github.com/dop251/goja"

	"github.com/MomiziTech/Momizi/Internal/Controller"
	LevelDBFunc "github.com/MomiziTech/Momizi/Internal/Plugin/Tools/LevelDB"
	"github.com/MomiziTech/Momizi/Tools/Log"
)

var DataPath = Controller.DataPath + "/"

type LevelDB struct {
	VM *goja.Runtime
}

/**
 * @description: 注册配置文件函数
 * @param {*goja.Runtime} VM 实例
 * @return {error} 错误信息
 */
func RegistrationFunction(VM *goja.Runtime) error {
	return VM.Set("LevelDB", LevelDB{VM: VM})
}

/**
 * @description: 创建数据库连接
 * @param {string} Path 数据库路径
 * @return {*LevelDBFunc.DB} 数据库连接
 */
func (LevelDB LevelDB) Open(Path string) *LevelDBFunc.DB {
	PluginName := LevelDB.VM.Get("PLUGIN_NAME").String()
	Path = DataPath + PluginName + "/" + Path
	DB := LevelDBFunc.Open(Path)
	return DB
}

/**
 * @description: 写入值
 * @param {*LevelDBFunc.DB} DB 数据库连接
 * @param {*} Key 键
 * @param {string} Value 值
 * @return {bool} 是否成功
 */
func (LevelDB LevelDB) Set(DB *LevelDBFunc.DB, Key, Value string) bool {
	PluginName := LevelDB.VM.Get("PLUGIN_NAME").String()
	if err := DB.Set(Key, Value); err != nil {
		Log.Error(PluginName, err)
		return false
	}
	return true
}

/**
 * @description: 读取值
 * @param {*LevelDBFunc.DB} DB 数据库连接
 * @param {string} Key 键
 * @return {any} 值
 */
func (LevelDB LevelDB) Get(DB *LevelDBFunc.DB, Key string) any {
	PluginName := LevelDB.VM.Get("PLUGIN_NAME").String()
	Value, err := DB.Get(Key)
	if err != nil {
		Log.Error(PluginName, err)
		return nil
	}
	return Value
}

/**
 * @description: 删除值
 * @param {*LevelDBFunc.DB} DB 数据库连接
 * @param {string} Key 键
 * @return {bool} 是否成功
 */
func (LevelDB LevelDB) Delete(DB *LevelDBFunc.DB, Key string) bool {
	PluginName := LevelDB.VM.Get("PLUGIN_NAME").String()
	if err := DB.Delete(Key); err != nil {
		Log.Error(PluginName, err)
		return false
	}
	return true
}

/**
 * @description: 获取键列表
 * @param {*LevelDBFunc.DB} DB 数据库连接
 * @param {string} Prefix 键前缀
 * @param {string} Suffix 键后缀
 * @return {[]string} 键列表
 */
func (LevelDB LevelDB) Key(DB *LevelDBFunc.DB, Prefix, Suffix string) []string {
	PluginName := LevelDB.VM.Get("PLUGIN_NAME").String()
	Value, err := DB.Key(Prefix, Suffix)
	if err != nil {
		Log.Error(PluginName, err)
		return nil
	}
	return Value
}

/**
 * @description: 关闭连接
 * @param {*LevelDBFunc.DB} DB 数据库连接
 * @return {*}
 */
func (LevelDB LevelDB) Close(DB *LevelDBFunc.DB) bool {
	if DB.Close() != nil {
		return false
	}
	return true
}
