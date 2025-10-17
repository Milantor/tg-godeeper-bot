package main

import (
	"fmt"
	"os"

	"github.com/milantor/tg-godeeper-bot/internal/ai"
	"github.com/milantor/tg-godeeper-bot/internal/tg"
)

func main() {
	fmt.Println("Hello from tg-godeeper-bot!")

	deepseek_api_key := os.Getenv("DEEPSEEK_API_KEY")
	if deepseek_api_key == "" {
		fmt.Println("Error: DEEPSEEK_API_KEY environment variable is not set.")
		return
	}
	if len(deepseek_api_key) < 10 {
		fmt.Println("Error: DEEPSEEK_API_KEY environment variable is not valid")
		return
	}
	ai.Init(deepseek_api_key)

	tg_api_key := os.Getenv("TELEGRAM_API_KEY")
	if tg_api_key == "" {
		fmt.Println("Error: TELEGRAM_API_KEY environment variable is not set.")
		return
	}
	err := tg.Init(tg_api_key)
	if err != nil {
		fmt.Println("Error: Telegram module initialization failed:", err)
		return
	}
}
