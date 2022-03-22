/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:27:25
 * @LastEditTime: 2022-03-14 17:27:26
 * @LastEdit: McPlus
 * @Description: Venue结构体
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Telegram\Venue.go
 */
package Telegram

type Venue struct {
	Location        Location `json:"location"`          // Venue location. Can't be a live location
	Title           string   `json:"title"`             // Name of the venue
	Address         string   `json:"address"`           // Address of the venue
	FoursquareID    string   `json:"foursquare_id"`     // Optional. Foursquare identifier of the venue
	FoursquareType  string   `json:"foursquare_type"`   // Optional. Foursquare type of the venue. (For example, "arts_entertainment/default", "arts_entertainment/aquarium" or "food/icecream".)
	GooglePlaceID   string   `json:"google_place_id"`   // Optional. Google Places identifier of the venue
	GooglePlaceType string   `json:"google_place_type"` // Optional. Google Places type of the venue. (See supported types.)
}
