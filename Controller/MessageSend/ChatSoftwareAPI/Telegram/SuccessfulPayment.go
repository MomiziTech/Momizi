/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:32:16
 * @LastEditTime: 2022-03-14 17:32:18
 * @LastEdit: McPlus
 * @Description: SuccessfulPayment结构体
 * @FilePath: \Momizi\Controller\MessageSend\ChatSoftwareAPI\Telegram\SuccessfulPayment.go
 */
package Telegram

type SuccessfulPayment struct {
	Currency                string    `json:"currency"`                   // Three-letter ISO 4217 currency code
	TotalAmount             int       `json:"total_amount"`               // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	InvoicePayload          string    `json:"invoice_payload"`            // Bot specified invoice payload
	ShippingOptionID        string    `json:"shipping_option_id"`         // Optional. Identifier of the shipping option chosen by the user
	OrderInfo               OrderInfo `json:"order_info"`                 // Optional. Order info provided by the user
	TelegramPaymentChargeID string    `json:"telegram_payment_charge_id"` // Telegram payment identifier
	ProviderPaymentChargeID string    `json:"provider_payment_charge_id"` // Provider payment identifier
}
