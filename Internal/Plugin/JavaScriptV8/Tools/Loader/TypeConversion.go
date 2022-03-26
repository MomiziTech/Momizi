/*
 * @Author: NyanCatda
 * @Date: 2022-03-26 18:48:29
 * @LastEditTime: 2022-03-26 19:03:53
 * @LastEditors: NyanCatda
 * @Description: 常见类型转换封装
 * @FilePath: \Momizi\Internal\Plugin\JavaScriptV8\Tools\Loader\TypeConversion.go
 */
package Loader

import (
	"encoding/json"

	"rogchap.com/v8go"
)

/**
 * @description: V8JavaScript String数组转换为Go String数组
 * @param {*v8go.Context} Context V8上下文
 * @param {*v8go.Value} Array String数组
 * @return {*}
 */
func V8StringArrayToGoStringArray(Array *v8go.Value) ([]string, error) {
	// 将数组转换为对象
	ArrayObject, err := Array.AsObject()
	if err != nil {
		return nil, err
	}
	// 获取数组长度
	ArrayLength, err := ArrayObject.Get("length")
	if err != nil {
		return nil, err
	}
	var Arrays []string
	for i := 0; i < int(ArrayLength.Integer()); i++ {
		Array, err := ArrayObject.GetIdx(uint32(i))
		if err != nil {
			return nil, err
		}
		Arrays = append(Arrays, Array.String())
	}
	return Arrays, nil
}

/**
 * @description: Go结构体转换为V8JavaScript对象(使用Json中转)
 * @param {*v8go.Context} Context V8上下文
 * @param {any} Struct Go结构体
 * @return {*}
 */
func GoStructToV8Object(Context *v8go.Context, Struct any) (*v8go.Value, error) {
	StructJson, err := json.Marshal(Struct)
	if err != nil {
		return nil, err
	}

	ObjectValue, err := v8go.JSONParse(Context, string(StructJson))
	if err != nil {
		return nil, err
	}

	return ObjectValue, nil
}
