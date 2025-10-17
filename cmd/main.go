package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/milantor/tg-godeeper-bot/internal/ai"
	"github.com/milantor/tg-godeeper-bot/internal/tg"
	"github.com/milantor/tg-godeeper-bot/logger"
)

func main() {
	fmt.Println("\033[33mHello from tg-godeeper-bot!\033[0m")

	_logger := logger.Init(true)
	slog.SetDefault(_logger)
	slog.Info("Logger successfully initialized!")

	tg_api_key := os.Getenv("TELEGRAM_API_KEY")
	if tg_api_key == "" {
		slog.Error("Error: TELEGRAM_API_KEY environment variable is not set.")
		return
	}
	if !strings.Contains(tg_api_key, ":") || len(tg_api_key) < 10 {
		slog.Error("Error: TELEGRAM_API_KEY environment variable is not valid.")
		return
	}
	tg.Init(tg_api_key)

	deepseek_api_key := os.Getenv("DEEPSEEK_API_KEY")
	if deepseek_api_key == "" {
		slog.Error("Error: DEEPSEEK_API_KEY environment variable is not set.")
		return
	}
	if len(deepseek_api_key) < 10 {
		slog.Error("Error: DEEPSEEK_API_KEY environment variable is not valid")
		return
	}
	ai.Init(deepseek_api_key)
}
