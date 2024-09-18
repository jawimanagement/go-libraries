package models

type TelegramMessage struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode,omitempty"` // Optional: Use "HTML" or "Markdown" if needed
}
