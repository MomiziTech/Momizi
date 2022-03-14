/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:36:28
 * @LastEditTime: 2022-03-14 17:36:29
 * @LastEdit: McPlus
 * @Description: EncryptedCredentials结构体
 * @FilePath: \Momizi\Controller\MessageSend\ChatSoftwareAPI\Telegram\Passport\EncryptedCredentials.go
 */
package Telegram

type EncryptedCredentials struct {
	Data   string `json:"data"`   // Base64-encoded encrypted JSON-serialized data with unique user's payload, data hashes and secrets required for EncryptedPassportElement decryption and authentication
	Hash   string `json:"hash"`   // Base64-encoded data hash for data authentication
	Secret string `json:"secret"` // Base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption
}
