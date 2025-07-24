package telegram

import (
	"context"
	"io"
	"log"

	tbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Setting struct {
	Token  string
	ChatID int64
}

type teleRepo struct {
	setting Setting
	bot     *tbot.BotAPI
	updates tbot.UpdatesChannel
	msg     chan string
}

type RepositoryI interface {
	SendText(ctx context.Context, data string) error
	SendImageFromFile(ctx context.Context, caption string, filePath string) error
	SendImageFromURL(ctx context.Context, caption string, url string) error
	SendImageFromReader(ctx context.Context, caption string, r io.Reader, filename string) error

	Received(ctx context.Context) error
	OnMsg() chan string
}

func Newtelegram(setting Setting) (RepositoryI, error) {
	var updates tbot.UpdatesChannel

	bot, err := tbot.NewBotAPI(setting.Token)
	if err != nil {
		return &teleRepo{
			setting: setting,
			bot:     bot,
			updates: updates,
			msg:     nil,
		}, err

	}
	bot.Debug = false

	u := tbot.UpdateConfig{
		Offset:  0,
		Timeout: 60,
	}
	log.Printf("succesed Authorized account:%s", bot.Self.UserName)

	updates = bot.GetUpdatesChan(u)

	return &teleRepo{
		setting: setting,
		bot:     bot,
		updates: updates,
		msg:     make(chan string),
	}, nil
}
