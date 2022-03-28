/*
 * @Author: McPlus
 * @Date: 2022-03-23 20:22:16
 * @LastEditTime: 2022-03-28 13:54:57
 * @LastEditors: NyanCatda
 * @Description: Message
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Line\Message.go
 */
package Line

import (
	"strconv"
	"time"

	"github.com/MomiziTech/Momizi/Tools"
	"github.com/MomiziTech/Momizi/Tools/File"
	"github.com/MomiziTech/Momizi/Tools/ReadConfig"
)

func GetContent(MessageID string) (string, error) {
	Config := ReadConfig.GetConfig

	ConfigTelegram := Config.ChatSoftware.Line

	APIAddress := "https://api-data.line.me/v2/bot/message/" + MessageID + "/content"

	Header := []string{
		"Authorization: Bearer " + ConfigTelegram.APIToken,
	}

	Path, _, Error := Tools.DownloadFile(APIAddress, Header, "data/File/Line/"+strconv.FormatInt(time.Now().Unix(), 10)+"/", true, 120)

	// 修改文件类型
	Path, err := File.CorrectFileType(Path)
	if err != nil {
		return "", err
	}

	return Path, Error

}
