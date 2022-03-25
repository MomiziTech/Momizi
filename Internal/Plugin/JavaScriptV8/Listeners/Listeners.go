/*
 * @Author: McPlus
 * @Date: 2022-03-24 20:54:07
 * @LastEditTime: 2022-03-24 21:20:25
 * @LastEditors: NyanCatda
 * @Description: 监听器
 * @FilePath: \Internal\Plugin\JavaScriptV8\Listeners\Listeners.go
 */
package Listeners

import (
	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScriptV8/Events"

	"rogchap.com/v8go"
)

func InitListeners(Isolate *v8go.Isolate, Context *v8go.Context) *v8go.Object {
	Listener, _ := v8go.NewObjectTemplate(Isolate)

	MessageEvent := Events.InitMessageEvent(Isolate)

	Listener.Set("Message", MessageEvent)

	ListenerObject, _ := Listener.NewInstance(Context)

	return ListenerObject
}
