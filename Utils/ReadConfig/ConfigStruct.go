/*
 * @Author: NyanCatda
 * @Date: 2022-03-08 21:25:31
 * @LastEditTime: 2022-03-08 21:55:24
 * @LastEditors: NyanCatda
 * @Description: 配置文件结构体
 * @FilePath: \Momizi\Utils\ReadConfig\ConfigStruct.go
 */

package ReadConfig

type Config struct {
	Run struct {
		WebHook struct {
			Port string `yaml:"Port"` // WebHook接收端口
			Key  string `yaml:"Key"`  // WebHook接收密钥
		} `yaml:"WebHook"`
	} `yaml:"Run"`
	ChatSoftware struct {
		Mirai struct {
			Switch      bool   `yaml:"Switch"`      // Mirai处理开关
			APILink     string `yaml:"APILink"`     // Mirai API地址
			BotQQNumber int    `yaml:"BotQQNumber"` // Mirai机器人QQ号
			VerifyKey   string `yaml:"VerifyKey"`   // Mirai验证密钥
		} `yaml:"QQ"`
		Telegram struct {
			Switch     bool   `yaml:"Switch"`     // Telegram处理开关
			APIToken   string `yaml:"APIToken"`   // Telegram Bot API Token
			BotAPILink string `yaml:"BotAPILink"` // Telegram Bot API地址
		} `yaml:"Telegram"`
		Line struct {
			Switch     bool   `yaml:"Switch"`     // Line处理开关
			APIToken   string `yaml:"APIToken"`   // Line Bot API Token
			BotAPILink string `yaml:"BotAPILink"` // Line Bot API地址
		} `yaml:"Line"`
		KaiHeiLa struct {
			Switch   bool   `yaml:"Switch"`   // 开黑啦处理开关
			APIToken string `yaml:"APIToken"` // 开黑啦Bot API Token
		} `yaml:"KaiHeiLa"`
	} `yaml:"ChatSoftware"`
}
