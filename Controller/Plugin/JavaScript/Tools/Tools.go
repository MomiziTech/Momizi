/*
 * @Author: NyanCatda
 * @Date: 2022-03-20 22:20:17
 * @LastEditTime: 2022-03-21 00:13:49
 * @LastEditors: NyanCatda
 * @Description: 工具函数注册
 * @FilePath: \Momizi\Controller\Plugin\JavaScript\Tools\Tools.go
 */
package Tools

import (
	"github.com/dop251/goja"
)

func Tools(VM *goja.Runtime) error {
	// 打印日志函数
	if err := LogPrint(VM); err != nil {
		return err
	}

	return nil
}
