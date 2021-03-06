/*
 * @Author: NyanCatda
 * @Date: 2022-03-26 18:48:29
 * @LastEditTime: 2022-04-03 13:29:27
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

/**
 * @description: V8JavaScript对象转换为Go String Map(使用Json中转)
 * @param {*v8go.Context} Context V8上下文
 * @param {v8go.Valuer} Object V8JavaScript对象
 * @return {map[any]any} Map
 * @return {error} 错误
 */
func V8ObjectToGoStringMap(Context *v8go.Context, Object v8go.Valuer) (map[string]string, error) {
	ObjectValue, err := v8go.JSONStringify(Context, Object)
	if err != nil {
		return nil, err
	}

	var Map map[string]string
	err = json.Unmarshal([]byte(ObjectValue), &Map)
	if err != nil {
		return nil, err
	}
	return Map, nil
}

/**
 * @description: Go数组转V8对象(使用Json中转)
 * @param {*v8go.Context} Context V8上下文
 * @param {[]string|[]int} Array Go数组
 * @return {*}
 */
func GoArrayToV8Object[T []string | []int](Context *v8go.Context, Array T) (*v8go.Value, error) {
	ArrayJson, err := json.Marshal(Array)
	if err != nil {
		return nil, err
	}

	ObjectValue, err := v8go.JSONParse(Context, string(ArrayJson))

	return ObjectValue, err
}

/**
 * @description: V8JavaScript对象转换为Go String Any Map(使用Json中转)
 * @param {*v8go.Context} Context V8上下文
 * @param {v8go.Valuer} Object V8JavaScript对象
 * @return {map[any]any} Map
 * @return {error} 错误
 */
func V8ObjectToGoStringAnyMap(Context *v8go.Context, Object v8go.Valuer) (map[string]any, error) {
	ObjectValue, err := v8go.JSONStringify(Context, Object)
	if err != nil {
		return nil, err
	}

	var Map map[string]any
	err = json.Unmarshal([]byte(ObjectValue), &Map)
	if err != nil {
		return nil, err
	}
	return Map, nil
}
