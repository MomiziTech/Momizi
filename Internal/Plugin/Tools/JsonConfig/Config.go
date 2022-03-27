/*
 * @Author: NyanCatda
 * @Date: 2022-03-23 20:18:31
 * @LastEditTime: 2022-03-27 21:33:20
 * @LastEditors: NyanCatda
 * @Description: 配置文件操作
 * @FilePath: \Momizi\Internal\Plugin\Tools\JsonConfig\Config.go
 */
package JsonConfig

import (
	"encoding/json"
	"os"

	"github.com/MomiziTech/Momizi/Tools/File"
)

type Config struct {
	Path    string
	Default map[string]any
}

/**
 * @description: 创建配置文件类
 * @param {string} Path 配置文件路径
 * @param {string} Default 默认配置
 * @return {*}
 */
func New(Path string, Default map[string]any) *Config {
	return &Config{Path: Path, Default: Default}
}

/**
 * @description: 初始化配置文件
 * @param {*}
 * @return {error} 错误信息
 */
func (Config Config) Init() error {
	// 判断文件是否存在
	if !File.Exists(Config.Path) {
		// 写入默认配置文件
		DefaultJson, err := json.Marshal(Config.Default)
		if err != nil {
			return err
		}

		File, err := File.NewFileReadWrite(Config.Path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE)
		if err != nil {
			return err
		}
		defer File.Close()

		return File.WriteTo(string(DefaultJson))
	}
	return nil
}

/**
 * @description: 读取配置项
 * @param {string} Name 配置项名称
 * @return {any} 配置项值
 * @return {error} 错误信息
 */
func (Config Config) Get(Name string) (any, error) {
	// 如果配置文件不存在，则返回默认配置
	if !File.Exists(Config.Path) {
		return Config.Default[Name], nil
	}

	File, err := File.NewFileReadWrite(Config.Path, os.O_RDONLY)
	if err != nil {
		return nil, err
	}
	defer File.Close()

	// 读取配置文件
	JsonBody, err := File.Read()
	if err != nil {
		return nil, err
	}
	// 解析配置文件
	JsonMap := make(map[string]any)
	if err = json.Unmarshal([]byte(JsonBody), &JsonMap); err != nil {
		return nil, err
	}

	// 如果配置项不存在，则返回默认配置
	if _, ok := JsonMap[Name]; !ok {
		return Config.Default[Name], nil
	}

	// 获取配置
	return JsonMap[Name], nil
}

/**
 * @description: 写入配置项
 * @param {string} Name 配置项名称
 * @param {any} Value 配置项值
 * @return {error} 错误信息
 */
func (Config Config) Set(Name string, Value any) error {
	// 如果配置文件不存在，则初始化配置文件
	if !File.Exists(Config.Path) {
		if err := Config.Init(); err != nil {
			return err
		}
	}

	// 以只读方式打开
	FileRead, err := File.NewFileReadWrite(Config.Path, os.O_RDONLY)
	if err != nil {
		return err
	}
	defer FileRead.Close()

	// 读取配置文件
	JsonBody, err := FileRead.Read()
	if err != nil {
		return err
	}
	// 解析配置文件
	JsonMap := make(map[string]any)
	if err = json.Unmarshal([]byte(JsonBody), &JsonMap); err != nil {
		return err
	}
	// 设置配置
	JsonMap[Name] = Value
	// 写入配置文件
	NewJsonBody, err := json.Marshal(JsonMap)
	if err != nil {
		return err
	}

	// 以覆盖写入模式打开
	FileWriteTo, err := File.NewFileReadWrite(Config.Path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE)
	if err != nil {
		return err
	}
	defer FileWriteTo.Close()

	return FileWriteTo.WriteTo(string(NewJsonBody))
}

/**
 * @description: 删除配置项
 * @param {string} Name 配置项名称
 * @return {error} 错误信息
 */
func (Config Config) Delete(Name string) error {
	// 如果配置文件不存在，则初始化配置文件
	if !File.Exists(Config.Path) {
		if err := Config.Init(); err != nil {
			return err
		}
	}

	// 以只读方式打开
	FileRead, err := File.NewFileReadWrite(Config.Path, os.O_RDONLY)
	if err != nil {
		return err
	}
	defer FileRead.Close()

	// 读取配置文件
	JsonBody, err := FileRead.Read()
	if err != nil {
		return err
	}
	// 解析配置文件
	JsonMap := make(map[string]any)
	if err = json.Unmarshal([]byte(JsonBody), &JsonMap); err != nil {
		return err
	}
	// 删除配置项
	delete(JsonMap, Name)
	// 写入配置文件
	NewJsonBody, err := json.Marshal(JsonMap)
	if err != nil {
		return err
	}

	// 以覆盖写入模式打开
	FileWriteTo, err := File.NewFileReadWrite(Config.Path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE)
	if err != nil {
		return err
	}
	defer FileWriteTo.Close()

	return FileWriteTo.WriteTo(string(NewJsonBody))
}
