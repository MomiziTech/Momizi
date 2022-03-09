/*
 * @Author: NyanCatda
 * @Date: 2022-03-09 15:53:21
 * @LastEditTime: 2022-03-09 15:57:52
 * @LastEditors: NyanCatda
 * @Description: WebHook解析结构体
 * @FilePath: \Momizi\Controller\MessageReceiving\WebHookStruct\WebHook.go
 */
package WebHookStruct

type WebHook struct {
	// 注册TelegramWebHook接收
	Telegram
	// 注册MiraiWebHook接收
	Mirai
	// 注册LineWebHook接收
	Line
}
