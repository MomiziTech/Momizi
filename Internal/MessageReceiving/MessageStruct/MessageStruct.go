/*
 * @Author: NyanCatda
 * @Date: 2022-03-13 00:03:26
 * @LastEditTime: 2022-03-26 00:21:00
 * @LastEditors: NyanCatda
 * @Description: 消息结构体
 * @FilePath: \Momizi\Internal\MessageReceiving\MessageStruct\MessageStruct.go
 */
package MessageStruct

/**
 * @description: 文件信息结构体
 * @param {*}
 * @return {*}
 */
type MessageChainFile struct {
	MimeType string `json:"MimeType"` // 文件类型(MIME type, 如：image/jpeg)
	Path     string `json:"Path"`     // 文件本地路径
	URL      string `json:"URL"`      // 文件网络路径（并不是所有文件都有）
	Name     string `json:"Name"`     // 文件名
	Size     int64  `json:"Size"`     // 文件大小(Byte)
}

/**
 * @description: 消息链结构体
 * @param {*}
 * @return {*}
 */
type MessageChain struct {
	Type string           `json:"Type"` // 消息类型，Text, Image, Audio, File
	Text string           `json:"Text"` // 文本消息内容
	File MessageChainFile `json:"File"` // 文件信息
}

/**
 * @description: 消息发送群聊信息
 * @param {*}
 * @return {*}
 */
type MessageSenderGroup struct {
	ID      string `json:"ID"`      // 群聊ID
	Title   string `json:"Title"`   // 群聊名称
	IsAdmin bool   `json:"IsAdmin"` // 是否为管理员
}

/**
 * @description: 消息发送者信息
 * @param {*}
 * @return {*}
 */
type MessageSender struct {
	ID       string             `json:"ID"`       // 消息发送者ID
	Username string             `json:"Username"` // 消息发送者用户名
	Group    MessageSenderGroup `json:"Group"`    // 群聊信息
}

/**
 * @description: 消息结构体
 * @param {*}
 * @return {*}
 */
type MessageStruct struct {
	ID           string         `json:"ID"`           // 消息ID
	Type         string         `json:"Type"`         // 消息来源类型，User, Group, Other
	ChatSoftware string         `json:"ChatSoftware"` // 消息来源软件，QQ, Telegram, Line
	Time         int            `json:"Time"`         // 消息接收时间戳
	MessageChain []MessageChain `json:"MessageChain"` // 消息链
	Sender       MessageSender  `json:"Sender"`       // 消息发送者信息
}
