/*
 * @Author: NyanCatda
 * @Date: 2022-04-03 16:49:40
 * @LastEditTime: 2022-04-03 18:03:49
 * @LastEditors: NyanCatda
 * @Description: 文件操作API
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Mirai\File.go
 */
package Mirai

import (
	"encoding/json"
	"errors"

	"github.com/MomiziTech/Momizi/Tools/ReadConfig"
	"github.com/nyancatda/HttpRequest"
)

type FileInfo struct {
	ResponseTemplate
	Data struct {
		Name    string      `json:"name"`   // 文件名
		ID      string      `json:"id"`     // 文件ID
		Path    string      `json:"path"`   // 文件路径
		Parent  interface{} `json:"parent"` // 文件对象, 递归类型. null 为存在根目录
		Contact struct {
			ID         int    `json:"id"`         // 文件来源聊天ID
			Name       string `json:"name"`       // 文件来源聊天名称
			Permission string `json:"permission"` // 文件来源聊天权限
		} `json:"contact"`
		IsFile       bool `json:"isFile"`      // 是否为文件
		IsDirectory  bool `json:"isDirectory"` // 是否为目录
		DownloadInfo struct {
			Sha1           string `json:"sha1"`           // 文件SHA1
			Md5            string `json:"md5"`            // 文件MD5
			DownloadTimes  int    `json:"downloadTimes"`  // 下载次数
			UploaderID     int    `json:"uploaderId"`     // 上传者ID
			UploadTime     int    `json:"uploadTime"`     // 上传时间
			LastModifyTime int    `json:"lastModifyTime"` // 最后修改时间
			URL            string `json:"url"`            // 文件下载地址
		} `json:"downloadInfo"`
	} `json:"data"`
}

/**
 * @description: 获取文件信息
 * @param {string} FileID 文件ID
 * @return {FileInfo} 文件信息
 * @return {error} 错误信息
 */
func GetFileInfo(FileID string) (FileInfo, error) {
	Config := ReadConfig.GetConfig
	ConfigMirai := Config.ChatSoftware.Mirai

	// 组成请求信息
	SessionKey, err := GetSessionKey()
	if err != nil {
		return FileInfo{}, err
	}
	RequestBody := map[string]string{
		"sessionKey": SessionKey,
		"id":         FileID,
	}
	RequestBodyJson, err := json.Marshal(RequestBody)
	if err != nil {
		return FileInfo{}, err
	}
	APIAddress := ConfigMirai.APILink + "/file/info"
	Body, HttpResponse, err := HttpRequest.PostRequestJson(APIAddress, []string{}, string(RequestBodyJson))
	if err != nil {
		return FileInfo{}, err
	}
	if HttpResponse.StatusCode != 200 {
		return FileInfo{}, errors.New("MiraiAPI请求失败，Status: " + HttpResponse.Status)
	}

	// 解析返回信息
	var FileInfo FileInfo

	err = json.Unmarshal(Body, &FileInfo)
	if err != nil {
		return FileInfo, err
	}
	return FileInfo, nil
}
