/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:25:38
 * @LastEditTime: 2022-03-14 17:26:40
 * @LastEdit: McPlus
 * @Description: Poll结构体
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Telegram\Poll.go
 */
package Telegram

type Poll struct {
	ID                    string          `json:"id"`                      // Unique poll identifier
	Question              string          `json:"question"`                // Poll question, 1-300 characters
	Options               []PollOption    `json:"options"`                 // List of poll options
	TotalVoterCount       int             `json:"total_voter_count"`       // Total number of users that voted in the poll
	IsClosed              bool            `json:"is_closed"`               // True, if the poll is closed
	IsAnonymous           bool            `json:"is_anonymous"`            // True, if the poll is anonymous
	Type                  string          `json:"type"`                    // Poll type, currently can be "regular" or "quiz"
	AllowsMultipleAnswers bool            `json:"allows_multiple_answers"` // True, if the poll allows multiple answers
	CorrectOptionID       int             `json:"correct_option_id"`       // Optional. 0-based identifier of the correct answer option. Available only for polls in the quiz mode, which are closed, or was sent (not forwarded) by the bot or to the private chat with the bot.
	Explanation           string          `json:"explanation"`             // Optional. Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters
	ExplanationEntities   []MessageEntity `json:"explanation_entities"`    // Optional. Special entities like usernames, URLs, bot commands, etc. that appear in the explanation
	OpenPeriod            int             `json:"open_period"`             // Optional. Amount of time in seconds the poll will be active after creation
	CloseDate             int             `json:"close_date"`              // Optional. Point in time (Unix timestamp) when the poll will be automatically closed
}
