/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:33:03
 * @LastEditTime: 2022-03-14 17:33:04
 * @LastEdit: McPlus
 * @Description: OrderInfo结构体
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Telegram\OrderInfo.go
 */
package Telegram

type OrderInfo struct {
	Name            string          `json:"name"`             // Optional. User name
	PhoneNumber     string          `json:"phone_number"`     // Optional. User's phone number
	Email           string          `json:"email"`            // Optional. User email
	ShippingAddress ShippingAddress `json:"shipping_address"` // Optional. User shipping address
}
