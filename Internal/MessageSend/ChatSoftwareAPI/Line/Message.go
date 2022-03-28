/*
 * @Author: McPlus
 * @Date: 2022-03-23 20:22:16
 * @LastEditTime: 2022-03-28 15:16:09
 * @LastEditors: NyanCatda
 * @Description: Message
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Line\Message.go
 */
package Line

import (
	"strconv"
	"time"

	"github.com/MomiziTech/Momizi/Internal/Controller"
	"github.com/MomiziTech/Momizi/Tools"
	"github.com/MomiziTech/Momizi/Tools/File"
	"github.com/MomiziTech/Momizi/Tools/ReadConfig"
)

/**
 * @description: 获取文件内容
 * @param {string} MessageID 消息ID
 * @return {string} 文件路径
 * @return {int64} 文件大小
 * @return {error} 错误信息
 */
func GetContent(MessageID string) (string, int64, error) {
	Config := ReadConfig.GetConfig

	ConfigTelegram := Config.ChatSoftware.Line

	APIAddress := ConfigTelegram.DataAPILink + "v2/bot/message/" + MessageID + "/content"

	Header := []string{
		"Authorization: Bearer " + ConfigTelegram.APIToken,
	}

	Path, FileSize, Error := Tools.DownloadFile(APIAddress, Header, Controller.BotFilePath+"Line/"+strconv.FormatInt(time.Now().Unix(), 10)+"/", true, 120)

	// 修改文件类型
	Path, err := File.CorrectFileType(Path)
	if err != nil {
		return "", 0, err
	}

	return Path, FileSize, Error

}
