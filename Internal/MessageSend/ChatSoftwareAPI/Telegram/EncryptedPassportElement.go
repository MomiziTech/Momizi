/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:35:20
 * @LastEditTime: 2022-03-14 17:35:21
 * @LastEdit: McPlus
 * @Description: EncryptedPassportElement结构体
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Telegram\Passport\EncryptedPassportElement.go
 */
package Telegram

type EncryptedPassportElement struct {
	Type        string         `json:"type"`         // Element type. One of "personal_details", "passport", "driver_license", "identity_card", "internal_passport", "address", "utility_bill", "bank_statement", "rental_agreement", "passport_registration", "temporary_registration", "phone_number", "email".
	Data        string         `json:"data"`         // Optional. Base64-encoded encrypted Telegram Passport element data provided by the user, available for "personal_details", "passport", "driver_license", "identity_card", "internal_passport" and "address" types. Can be decrypted and verified using the accompanying EncryptedCredentials.
	PhoneNumber string         `json:"phone_number"` // Optional. User's verified phone number, available only for "phone_number" type
	Email       string         `json:"email"`        // Optional. User's verified email address, available only for "email" type
	Files       []PassportFile `json:"files"`        // Optional. Array of encrypted files with documents provided by the user, available for "utility_bill", "bank_statement", "rental_agreement", "passport_registration" and "temporary_registration" types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
	FrontSide   PassportFile   `json:"front_side"`   // Optional. Encrypted file with the front side of the document, provided by the user. Available for "passport", "driver_license", "identity_card" and "internal_passport". The file can be decrypted and verified using the accompanying EncryptedCredentials.
	ReverseSide PassportFile   `json:"reverse_side"` // Optional. Encrypted file with the reverse side of the document, provided by the user. Available for "driver_license" and "identity_card". The file can be decrypted and verified using the accompanying EncryptedCredentials.
	Selfie      PassportFile   `json:"selfie"`       // Optional. Encrypted file with the selfie of the user holding a document, provided by the user; available for "passport", "driver_license", "identity_card" and "internal_passport". The file can be decrypted and verified using the accompanying EncryptedCredentials.
	Translation []PassportFile `json:"translation"`  // Optional. Array of encrypted files with translated versions of documents provided by the user. Available if requested for "passport", "driver_license", "identity_card", "internal_passport", "utility_bill", "bank_statement", "rental_agreement", "passport_registration" and "temporary_registration" types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
	Hash        string         `json:"hash"`         // Base64-encoded element hash for using in PassportElementErrorUnspecified
}
