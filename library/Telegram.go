package library

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jawimanagement/go-libraries/models"
)

func sendTelegramMessage(chatID, message string) error {

	var telegramStruct models.TelegramMessage

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", os.Getenv("telegramBotToken"))

	telegramStruct.ChatID = chatID
	telegramStruct.Text = message
	telegramStruct.ParseMode = "HTML"

	telegramMessage := telegramStruct

	body, err := json.Marshal(telegramMessage)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("telegram API request failed with status code: %d", resp.StatusCode)
	}

	return nil
}
