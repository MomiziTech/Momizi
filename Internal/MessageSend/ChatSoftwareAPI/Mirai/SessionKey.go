/*
 * @Author: NyanCatda
 * @Date: 2022-04-03 16:56:12
 * @LastEditTime: 2022-04-03 18:01:57
 * @LastEditors: NyanCatda
 * @Description: SessionKey处理
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Mirai\SessionKey.go
 */
package Mirai

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"

	"github.com/MomiziTech/Momizi/Internal/Controller"
	"github.com/MomiziTech/Momizi/Tools/File"
	"github.com/MomiziTech/Momizi/Tools/ReadConfig"
	"github.com/nyancatda/HttpRequest"
)

var (
	SessionKeyCacheFile = Controller.BotFilePath + "Mirai/SessionKey"
)

/**
 * @description: 创建SessionKey
 * @param {*}
 * @return {string} SessionKey
 * @return {error} 错误信息
 */
func CreateSessionKey() (string, error) {
	Config := ReadConfig.GetConfig

	//组成请求消息
	MessageBody := map[string]interface{}{
		"verifyKey": Config.ChatSoftware.Mirai.VerifyKey,
	}

	MessageBodyJson, err := json.Marshal(MessageBody)
	if err != nil {
		return "", err
	}

	URL := Config.ChatSoftware.Mirai.APILink + "/verify"
	Body, HttpResponse, err := HttpRequest.PostRequestJson(URL, []string{}, string(MessageBodyJson))
	if err != nil {
		return "", err
	}
	if HttpResponse.StatusCode != 200 {
		return "", errors.New("MiraiAPI请求失败，Status: " + HttpResponse.Status)
	}

	// 解析返回信息
	type VerifyJson struct {
		Code    int    `json:"code"`
		Session string `json:"session"`
	}
	var VerifyInfo VerifyJson
	json.Unmarshal([]byte(Body), &VerifyInfo)
	if VerifyInfo.Code != 0 {
		return "", errors.New("MiraiAPI请求失败，ResponseCode: " + strconv.Itoa(VerifyInfo.Code))
	}

	SessionKey := VerifyInfo.Session

	//绑定SessionKey与QQ号
	MessageBody = map[string]interface{}{
		"sessionKey": SessionKey,
		"qq":         Config.ChatSoftware.Mirai.BotQQNumber,
	}

	MessageBodyJson, err = json.Marshal(MessageBody)
	if err != nil {
		return "", err
	}

	URL = Config.ChatSoftware.Mirai.APILink + "/bind"
	Body, HttpResponse, err = HttpRequest.PostRequestJson(URL, []string{}, string(MessageBodyJson))
	if err != nil {
		return "", err
	}
	if HttpResponse.StatusCode != 200 {
		return "", errors.New("MiraiAPI请求失败，Status: " + HttpResponse.Status)
	}

	//缓存SessionKey
	FileWrite, err := File.NewFileReadWrite(SessionKeyCacheFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE)
	if err != nil {
		return "", err
	}
	defer FileWrite.Close()
	FileWrite.WriteTo(SessionKey)

	return SessionKey, err
}

/**
 * @description: 获取SessionKey
 * @param {*}
 * @return {string} SessionKey
 * @return {error} 错误信息
 */
func GetSessionKey() (string, error) {
	// 判断文件是否存在
	if !File.Exists(SessionKeyCacheFile) {
		// 文件不存在则生成SessionKey并返回
		SessionKey, err := CreateSessionKey()
		if err != nil {
			return "", err
		}
		return SessionKey, nil
	}

	// 文件存在则读取SessionKey并返回
	FileRead, err := File.NewFileReadWrite(SessionKeyCacheFile, os.O_RDONLY)
	if err != nil {
		return "", err
	}
	defer FileRead.Close()
	SessionKey, err := FileRead.Read()
	if err != nil {
		return "", err
	}

	// 判断SessionKey是否过期
	Config := ReadConfig.GetConfig
	MessageBody := map[string]interface{}{
		"sessionKey": SessionKey,
		"qq":         Config.ChatSoftware.Mirai.BotQQNumber,
	}
	MessageBodyJson, err := json.Marshal(MessageBody)
	if err != nil {
		return "", err
	}
	URL := Config.ChatSoftware.Mirai.APILink + "/bind"
	Body, HttpResponse, err := HttpRequest.PostRequestJson(URL, []string{}, string(MessageBodyJson))
	if err != nil {
		return "", err
	}
	if HttpResponse.StatusCode != 200 {
		return "", errors.New("MiraiAPI请求失败，Status: " + HttpResponse.Status)
	}
	// 解析返回
	type Bind struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	var BindInfo Bind
	json.Unmarshal([]byte(Body), &BindInfo)
	if BindInfo.Code != 0 {
		// 重新生成SessionKey并返回
		SessionKey, err := CreateSessionKey()
		if err != nil {
			return "", err
		}
		return SessionKey, nil
	}

	return SessionKey, nil
}
