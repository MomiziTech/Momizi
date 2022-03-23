/*
 * @Author: NyanCatda
 * @Date: 2022-03-23 20:18:31
 * @LastEditTime: 2022-03-23 20:35:48
 * @LastEditors: NyanCatda
 * @Description: 配置文件操作
 * @FilePath: \Momizi\Internal\Plugin\Tools\JsonConfig\Config.go
 */
package JsonConfig

import (

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

/*
func (Config Config) Init() error {
	// 判断文件是否存在
	if !File.Exists(Config.Path) {
		// 写入默认配置文件
		json.Marshal(Config.Default)
	}
	// 读取文件
	ConfigBody, err := File.Read(Config.Path)
	if err != nil {
		return err
	}
	// 如果文件不存在，则创建文件
	// 如果文件存在，则读取文件
	// 如果文件存在，但是内容为空，则读取默认配置

}
*/
