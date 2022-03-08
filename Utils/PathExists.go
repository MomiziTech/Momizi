/*
 * @Author: NyanCatda
 * @Date: 2022-03-08 21:26:55
 * @LastEditTime: 2022-03-08 21:32:19
 * @LastEditors: NyanCatda
 * @Description: 路径检查
 * @FilePath: \Momizi\Utils\PathExists.go
 */
package Utils

import (
	"fmt"
	"os"
)

/**
 * @description: 判断文件夹是否存在，不存在则创建
 * @param {string} path 文件夹路径
 * @return {*}
 */
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		// 创建文件夹
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		} else {
			return true, nil
		}
	}
	return false, err
}
