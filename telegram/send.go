package telegram

import (
	"context"
	"io"
	"time"

	tbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t *teleRepo) SendText(ctx context.Context, data string) error {
	msg := tbot.NewMessage(t.setting.ChatID, data)
	//msg.ReplyToMessageID = update.Message.MessageID // if you want reply
	_, err := t.bot.Send(msg)
	if err != nil {
		return err
	}
	time.Sleep(500 * time.Millisecond)

	return nil
}

func (t *teleRepo) SendImageFromFile(ctx context.Context, caption string, filePath string) error {
	photo := tbot.NewPhoto(t.setting.ChatID, tbot.FilePath(filePath))
	photo.Caption = caption

	_, err := t.bot.Send(photo)
	if err != nil {
		return err
	}
	time.Sleep(500 * time.Millisecond)

	return nil
}

func (t *teleRepo) SendImageFromURL(ctx context.Context, caption string, url string) error {
	photo := tbot.NewPhoto(t.setting.ChatID, tbot.FileURL(url))
	photo.Caption = caption

	_, err := t.bot.Send(photo)
	if err != nil {
		return err
	}
	time.Sleep(500 * time.Millisecond)

	return nil
}

func (t *teleRepo) SendImageFromReader(ctx context.Context, caption string, r io.Reader, filename string) error {
	photo := tbot.NewPhoto(t.setting.ChatID, tbot.FileReader{
		Name:   filename,
		Reader: r,
	})
	photo.Caption = caption

	_, err := t.bot.Send(photo)
	if err != nil {
		return err
	}
	time.Sleep(500 * time.Millisecond)

	return nil
}
