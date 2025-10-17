package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/milantor/tg-godeeper-bot/internal/ai"
	"github.com/milantor/tg-godeeper-bot/internal/tg"
)

func main() {
	fmt.Println("Hello from tg-godeeper-bot!")

	tg_api_key := os.Getenv("TELEGRAM_API_KEY")
	if tg_api_key == "" {
		fmt.Println("Error: TELEGRAM_API_KEY environment variable is not set.")
		return
	}
	if !strings.Contains(tg_api_key, ":") || len(tg_api_key) < 10 {
		fmt.Println("Error: TELEGRAM_API_KEY environment variable is not valid.")
		return
	}
	tg.Init(tg_api_key)

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
}
