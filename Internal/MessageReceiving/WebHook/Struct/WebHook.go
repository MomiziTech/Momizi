/*
 * @Author: NyanCatda
 * @Date: 2022-03-09 15:53:21
 * @LastEditTime: 2022-03-12 22:43:42
 * @LastEditors: NyanCatda
 * @Description: WebHook解析结构体
 * @FilePath: \Momizi\Internal\MessageReceiving\WebHook\Struct\WebHook.go
 */
package Struct

type WebHook struct {
	// 注册TelegramWebHook接收
	Telegram
	// 注册MiraiWebHook接收
	Mirai
	// 注册LineWebHook接收
	Line
}
