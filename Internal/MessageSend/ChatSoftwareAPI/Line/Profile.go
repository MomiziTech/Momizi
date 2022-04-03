/*
 * @Author: NyanCatda
 * @Date: 2022-04-03 16:27:55
 * @LastEditTime: 2022-04-03 16:38:03
 * @LastEditors: NyanCatda
 * @Description: 个人资料获取
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Line\Profile.go
 */
package Line

import (
	"encoding/json"
	"errors"

	"github.com/MomiziTech/Momizi/Tools/ReadConfig"
	"github.com/nyancatda/HttpRequest"
)

type UserProfile struct {
	DisplayName   string `json:"displayName"`   // 用户名
	UserID        string `json:"userId"`        // 用户ID
	Language      string `json:"language"`      // 语言
	PictureURL    string `json:"pictureUrl"`    // 头像链接
	StatusMessage string `json:"statusMessage"` // 状态信息
}

/**
 * @description: 获取个人资料
 * @param {string} UserID 用户ID
 * @return {UserProfile} 个人资料
 * @return {error} 错误信息
 */
func GetProfile(UserID string) (UserProfile, error) {
	Config := ReadConfig.GetConfig
	ConfigLine := Config.ChatSoftware.Line

	// 组成请求信息
	APIAddress := ConfigLine.DataAPILink + "v2/bot/profile/" + UserID
	Header := []string{
		"Authorization: Bearer " + ConfigLine.APIToken,
	}

	Body, HttpResponse, err := HttpRequest.GetRequest(APIAddress, Header)
	if err != nil {
		return UserProfile{}, err
	}

	if HttpResponse.StatusCode != 200 {
		return UserProfile{}, errors.New("请求失败，Status: " + HttpResponse.Status)
	}

	// 解析返回信息
	var UserProfile UserProfile
	json.Unmarshal(Body, &UserProfile)

	return UserProfile, nil
}
