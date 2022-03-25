/*
 * @Author: McPlus
 * @Date: 2022-03-24 20:54:07
 * @LastEditTime: 2022-03-25 22:35:04
 * @LastEditors: NyanCatda
 * @Description: 监听器
 * @FilePath: \Momizi\Internal\Plugin\JavaScriptV8\Listeners\Listeners.go
 */
package Listeners

import (
	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScriptV8/Events"

	"rogchap.com/v8go"
)

/**
 * @description: 初始化监听器
 * @param {*v8go.Isolate} Isolate 虚拟机
 * @param {*v8go.Context} Context 上下文
 * @return {*v8go.Object} 监听器
 */
func InitListeners(Isolate *v8go.Isolate, Context *v8go.Context) *v8go.Object {
	Listener, _ := v8go.NewObjectTemplate(Isolate)
	MessageEvent := Events.InitMessageEvent(Isolate)

	// 注册消息监听器
	Listener.Set("Message", MessageEvent)

	ListenerObject, _ := Listener.NewInstance(Context)
	return ListenerObject
}
