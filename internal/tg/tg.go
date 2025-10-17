package tg

import (
	"fmt"
	"regexp"
)

func Init(api_key string) error {
	if !isValidTelegramAPIKey(api_key) {
		return fmt.Errorf("invalid API key")
	}
	fmt.Println("Telegram bot init successfully!")
	return nil
}

func isValidTelegramAPIKey(api_key string) bool {
	match, _ := regexp.MatchString(`^\d+:[A-Za-z0-9_-]+$`, api_key)
	return match
}
