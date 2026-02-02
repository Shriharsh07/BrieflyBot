package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func SendToTelegram(subject, summary string) error {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")

	if botToken == "" || chatID == "" {
		return fmt.Errorf("telegram bot token or chat id missing")
	}

	message := fmt.Sprintf(
		"📧 *Email Subject:*\n%s\n\n🧠 *Simplified Summary:*\n%s",
		subject,
		summary,
	)

	payload := map[string]interface{}{
		"chat_id":    chatID,
		"text":       message,
		"parse_mode": "Markdown",
	}

	data, _ := json.Marshal(payload)

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("telegram send failed with status %d", resp.StatusCode)
	}

	return nil
}
