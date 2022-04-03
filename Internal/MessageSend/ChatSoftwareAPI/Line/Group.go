/*
 * @Author: NyanCatda
 * @Date: 2022-04-03 16:39:38
 * @LastEditTime: 2022-04-03 16:42:44
 * @LastEditors: NyanCatda
 * @Description: 群组类请求封装
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Line\Group.go
 */
package Line

import (
	"encoding/json"
	"errors"

	"github.com/MomiziTech/Momizi/Tools/ReadConfig"
	"github.com/nyancatda/HttpRequest"
)

type GroupChatSummary struct {
	GroupID    string `json:"groupId"`    // 群组ID
	GroupName  string `json:"groupName"`  // 群组名称
	PictureURL string `json:"pictureUrl"` // 群组头像链接
}

/**
 * @description: 获取群组信息
 * @param {string} GroupID 群组ID
 * @return {GroupChatSummary} 群组信息
 * @return {error} 错误信息
 */
func GetGroupChatSummary(GroupID string) (GroupChatSummary, error) {
	Config := ReadConfig.GetConfig
	ConfigLine := Config.ChatSoftware.Line

	// 组成请求信息
	APIAddress := ConfigLine.DataAPILink + "v2/bot/group/" + GroupID + "/summary"
	Header := []string{
		"Authorization: Bearer " + ConfigLine.APIToken,
	}

	Body, HttpResponse, err := HttpRequest.GetRequest(APIAddress, Header)
	if err != nil {
		return GroupChatSummary{}, err
	}

	if HttpResponse.StatusCode != 200 {
		return GroupChatSummary{}, errors.New("请求失败，Status: " + HttpResponse.Status)
	}

	// 解析返回信息
	var GroupChatSummary GroupChatSummary
	json.Unmarshal(Body, &GroupChatSummary)

	return GroupChatSummary, nil
}
