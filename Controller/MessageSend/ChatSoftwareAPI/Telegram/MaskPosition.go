/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:20:40
 * @LastEditTime: 2022-03-14 17:49:30
 * @LastEdit: McPlus
 * @Description: MaskPosition结构体
 * @FilePath: \Momizi\Controller\MessageSend\ChatSoftwareAPI\Telegram\MaskPosition.go
 */
package Telegram

type MaskPosition struct {
	Point  string  `json:"point"`   // The part of the face relative to which the mask should be placed. One of "forehead", "eyes", "mouth", or "chin".
	XShift float64 `json:"x_shift"` // Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. For example, choosing -1.0 will place mask just to the left of the default mask position.
	YShift float64 `json:"y_shift"` // Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. For example, 1.0 will place the mask just below the default mask position.
	Scale  float64 `json:"scale"`   // Mask scaling coefficient. For example, 2.0 means double size.
}
