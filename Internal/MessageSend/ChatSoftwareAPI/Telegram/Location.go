/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:10:43
 * @LastEditTime: 2022-03-14 17:12:14
 * @LastEdit: McPlus
 * @Description:
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Telegram\Location.go
 */
package Telegram

type Location struct {
	Longitude            float64 `json:"longitude"`              // Longitude as defined by sender
	Latitude             float64 `json:"latitude"`               // Latitude as defined by sender
	HorizontalAccuracy   float64 `json:"horizontal_accuracy"`    // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod           int     `json:"live_period"`            // Optional. Time relative to the message sending date, during which the location can be updated; in seconds. For active live locations only.
	Heading              int     `json:"heading"`                // Optional. The direction in which user is moving, in degrees; 1-360. For active live locations only.
	ProximityAlertRadius int     `json:"proximity_alert_radius"` // Optional. Maximum distance for proximity alerts about approaching another chat member, in meters. For sent live locations only.
}
