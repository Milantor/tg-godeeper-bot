package tg

import "fmt"

func Init(api_key string) {
	if api_key == "" {
		fmt.Println("Error: API key is empty.")
		return
	}
	fmt.Println("Telegram bot init succesfully!")
}
