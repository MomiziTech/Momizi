/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:34:18
 * @LastEditTime: 2022-03-14 17:34:19
 * @LastEdit: McPlus
 * @Description: PassportData结构体
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Telegram\PassportData.go
 */
package Telegram

type PassportData struct {
	Data        []EncryptedPassportElement `json:"data"`        // Array with information about documents and other Telegram Passport elements that was shared with the bot
	Credentials EncryptedCredentials       `json:"credentials"` // Encrypted credentials required to decrypt the data
}
