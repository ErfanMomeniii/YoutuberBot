package YoutuberBot

import (
	"github.com/ErfanMomeniii/YoutuberBot/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotFactory interface {
	ConnectToBot() (*tgbotapi.BotAPI, error)
}

func ConnectToBot() (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(config.C.YoutuberBotConfig.BotConfig.Token)
	return bot, err
}

func main() {
}
