/*
 * @Author: NyanCatda
 * @Date: 2022-03-20 22:20:17
 * @LastEditTime: 2022-03-21 00:13:22
 * @LastEditors: NyanCatda
 * @Description: 工具函数注册
 * @FilePath: \Momizi\Controller\Plugin\JavaScript\Tools\Log.go
 */
package Tools

import (
	"fmt"

	"github.com/dop251/goja"
)

/**
 * @description: 打印日志函数
 * @param {*goja.Runtime} VM
 * @return {*}
 */
func LogPrint(VM *goja.Runtime) error {
	err := VM.Set("log", fmt.Println)
	if err != nil {
		return err
	}
	return nil
}
