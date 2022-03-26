/*
 * @Author: NyanCatda
 * @Date: 2022-03-26 10:21:35
 * @LastEditTime: 2022-03-26 19:04:26
 * @LastEditors: NyanCatda
 * @Description: HttpRequest函数注册
 * @FilePath: \Momizi\Internal\Plugin\JavaScriptV8\Tools\HttpRequest\HttpRequest.go
 */
package HttpRequest

import "rogchap.com/v8go"

func Register(Isolate *v8go.Isolate, Context *v8go.Context) *v8go.Object {
	HttpRequest, _ := v8go.NewObjectTemplate(Isolate)

	Get := Get(Isolate, Context)
	HttpRequest.Set("Get", Get)

	ConsoleObject, _ := HttpRequest.NewInstance(Context)
	return ConsoleObject
}
