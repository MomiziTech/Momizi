/*
 * @Author: NyanCatda
 * @Date: 2022-03-20 21:18:09
 * @LastEditTime: 2022-03-21 08:31:53
 * @LastEditors: Please set LastEditors
 * @Description: 事件监听器
 * @FilePath: \Momizi\Controller\Plugin\JavaScript\EventListeners\EventListeners.go
 */
package EventListeners

import (
	"github.com/dop251/goja"
)

/**
 * @description: 监听器
 */
type Listener struct {
	VM            *goja.Runtime
}

/**
 * @description: 监听器注册
 * @param {*goja.Runtime} VM 加载器
 * @param {MessageStruct.MessageStruct} Message 消息结构体
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
