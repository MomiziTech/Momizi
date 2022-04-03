/*
 * @Author: NyanCatda
 * @Date: 2022-04-03 17:49:08
 * @LastEditTime: 2022-04-03 17:51:11
 * @LastEditors: NyanCatda
 * @Description: Mirai APi封装
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Mirai\Mirai.go
 */
package Mirai

type ResponseTemplate struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
