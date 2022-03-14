/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:30:14
 * @LastEditTime: 2022-03-14 17:32:04
 * @LastEdit: McPlus
 * @Description: Invoice结构体
 * @FilePath: \Momizi\Controller\MessageSend\ChatSoftwareAPI\Telegram\Invoice.go
 */
package Telegram

type Invoice struct {
	Title          string `json:"title"`           // Product name
	Description    string `json:"description"`     // Product description
	StartParameter string `json:"start_parameter"` // Unique bot deep-linking parameter that can be used to generate this invoice
	Currency       string `json:"currency"`        // Three-letter ISO 4217 currency code
	TotalAmount    int    `json:"total_amount"`    // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
}
