/*
 * @Author: NyanCatda
 * @Date: 2022-03-13 00:03:26
 * @LastEditTime: 2022-03-19 16:43:11
 * @LastEditors: NyanCatda
 * @Description: 消息结构体
 * @FilePath: \Momizi\Controller\MessageReceiving\MessageStruct\MessageStruct.go
 */
package MessageStruct

/**
 * @description: 文件信息结构体
 * @param {*}
 * @return {*}
 */
type MessageChainFile struct {
	MimeType string `json:"mime_type"` // 文件类型(MIME type, 如：image/jpeg)
	Path     string `json:"path"`      // 文件本地路径
	URL      string `json:"url"`       // 文件网络路径（并不是所有文件都有）
	Name     string `json:"name"`      // 文件名
	Size     int64  `json:"size"`      // 文件大小(Byte)
}

/**
 * @description: 消息链结构体
 * @param {*}
 * @return {*}
 */
type MessageChain struct {
	Type string           `json:"type"` // 消息类型，Text, Image, Audio, File
	Text string           `json:"text"` // 文本消息内容
	File MessageChainFile `json:"file"` // 文件信息
}

/**
 * @description: 消息发送群聊信息
 * @param {*}
 * @return {*}
 */
type MessageSenderGroup struct {
	ID      string `json:"id"`       // 群聊ID
	Title   string `json:"title"`    // 群聊名称
	IsAdmin bool   `json:"is_admin"` // 是否为管理员
}

/**
 * @description: 消息发送者信息
 * @param {*}
 * @return {*}
 */
type MessageSender struct {
	ID       string             `json:"id"`       // 消息发送者ID
	Username string             `json:"username"` // 消息发送者用户名
	Group    MessageSenderGroup `json:"group"`    // 群聊信息
}

/**
 * @description: 消息结构体
 * @param {*}
 * @return {*}
 */
type MessageStruct struct {
	ID           string         `json:"id"`            // 消息ID
	Type         string         `json:"type"`          // 消息来源类型，User, Group, Other
	ChatSoftware string         `json:"chat_software"` // 消息来源软件，QQ, Telegram, Line
	Time         int            `json:"time"`          // 消息接收时间戳
	MessageChain []MessageChain `json:"message_chain"` // 消息链
	Sender       MessageSender  `json:"sender"`        // 消息发送者信息
}
