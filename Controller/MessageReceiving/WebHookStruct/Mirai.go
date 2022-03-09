/*
 * @Author: NyanCatda
 * @Date: 2022-03-09 14:21:12
 * @LastEditTime: 2022-03-09 16:10:06
 * @LastEditors: NyanCatda
 * @Description: Telegram WebHook结构体
 * @FilePath: \Momizi\Controller\MessageReceiving\WebHookStruct\Mirai.go
 */
package WebHookStruct

type Mirai struct {
	Type         string `json:"type"` // 聊天类型，可能类型："FriendMessage"，"GroupMessage"，"TempMessage"，"OtherClientMessage"
	MessageChain []struct {
		Type string `json:"type"` // 消息类型，可能类型："Source"，"Quote"，"At"，"AtAll"，"Plain"，"Image，"Voice"，"Xml"，"Json"，"App"，"Poke"，"Forward"
		ID   int    `json:"id"`   // 消息ID
		Time int    `json:"time"` // 消息时间戳

		// 文字消息
		Text string `json:"text"` // 文字消息内容

		// 图片消息
		ImageID string `json:"imageId"` // 图片消息ID

		// 语音消息
		VoiceID string `json:"voiceId"` // 语音消息ID
		Length  int    `json:"length"`  // 语音消息长度(ms)

		// 回复消息
		SenderID int `json:"senderId"` // 被引用回复的原消息的发送者的QQ号
		TargetID int `json:"targetId"` // 被引用回复的原消息的发送者的QQ号
		GroupID  int `json:"groupId"`
		Origin   []struct {
			Type string `json:"type"` // 消息类型，可能类型："Source"，"Quote"，"At"，"AtAll"，"Plain"，"Image，"Voice"，"Xml"，"Json"，"App"，"Poke"，"Forward"
			ID   int    `json:"id"`   // 消息ID
			Time int    `json:"time"` // 消息时间戳

			// 文字消息
			Text string `json:"text"` // 文字消息内容

			// 图片消息
			ImageID string `json:"imageId"` // 图片消息ID

			// 语音消息
			VoiceID string `json:"voiceId"` // 语音消息ID
			Length  int    `json:"length"`  // 语音消息长度(ms)

			// At消息
			Target  int    `json:"target"`  // 被@的QQ号
			Display string `json:"display"` // At时显示的文字

			URL    string `json:"url"`    //文件链接
			Path   string `json:"path"`   //文件路径
			Base64 string `json:"base64"` //文件base64编码
		} `json:"origin"` // 原消息链

		// At消息
		Target  int    `json:"target"`  // 被@的QQ号
		Display string `json:"display"` // At时显示的文字

		URL    string `json:"url"`    //文件链接
		Path   string `json:"path"`   //文件路径
		Base64 string `json:"base64"` //文件base64编码
	} `json:"messageChain"` // 消息链
	Sender struct {
		ID int `json:"id"` // 发送者QQ号
		// 私聊
		Nickname string `json:"nickname"` // 发送者昵称
		Remark   string `json:"remark"`   // 发送者备注

		// 群聊
		MemberName         string `json:"memberName"`         // 发送者群昵称
		SpecialTitle       string `json:"specialTitle"`       // 发送者群头衔
		Permission         string `json:"permission"`         // 发送者群权限
		JoinTimestamp      int    `json:"joinTimestamp"`      // 发送者入群时间戳
		LastSpeakTimestamp int    `json:"lastSpeakTimestamp"` // 发送者最后发言时间戳
		MuteTimeRemaining  int    `json:"muteTimeRemaining"`  // 发送者禁言剩余时间
		Group              struct {
			ID         int    `json:"id"`         // 群号
			Name       string `json:"name"`       // 群名
			Permission string `json:"permission"` // 群权限，"ADMINISTRATOR"，"MEMBER"
		} `json:"group"`
	} `json:"sender"`
}
