package main

import (
	"os"

	bot "codaea.com/kalathia/Bot"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

bot.BotToken = os.Getenv("BOT_TOKEN")
bot.Run()
}
