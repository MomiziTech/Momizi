/*
 * @Author: NyanCatda
 * @Date: 2022-03-20 21:18:09
 * @LastEditTime: 2022-03-21 13:57:18
 * @LastEditors: NyanCatda
 * @Description: 事件监听器
 * @FilePath: \Momizi\Internal\Plugin\JavaScript\EventListeners\EventListeners.go
 */
package EventListeners

import (
	"github.com/dop251/goja"
)

/**
 * @description: 监听器
 */
type Listener struct {
	VM *goja.Runtime
}

/**
 * @description: 监听器注册
 * @param {*goja.Runtime} VM 加载器
 * @return {*}
 */
func Listeners(VM *goja.Runtime) error {
	// 注册消息事件
	err := VM.Set("Listener", Listener{VM})
	if err != nil {
		return err
	}

	return nil
}
