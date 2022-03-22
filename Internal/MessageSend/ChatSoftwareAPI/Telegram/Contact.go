/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:23:42
 * @LastEditTime: 2022-03-14 17:23:43
 * @LastEdit: McPlus
 * @Description: Contact结构体
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Telegram\Contact.go
 */

package Telegram

type Contact struct {
	PhoneNumber string `json:"phone_number"` // Contact's phone number
	FirstName   string `json:"first_name"`   // Contact's first name
	LastName    string `json:"last_name"`    // Optional. Contact's last name
	UserID      int    `json:"user_id"`      // Optional. Contact's user identifier in Telegram. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	VCard       string `json:"vcard"`        // Optional. Additional data about the contact in the form of a vCard
}
