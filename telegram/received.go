package telegram

import (
	"context"
	"errors"
	"log"
)

func (t *teleRepo) Received(ctx context.Context) error {
	for {
		select {
		case update, ok := <-t.updates:
			if !ok {
				return errors.New("updates channel closed")
			}
			if update.Message != nil {
				log.Println(" ============ Received ============ ")
				log.Printf("ChatID   :%d", update.Message.Chat.ID)
				log.Printf("MessageID:%d", update.Message.MessageID)
				log.Printf("Message  :%s", update.Message.Text)

				/*if update.Message.Chat.ID == t.setting.ChatID {
					// jika mau message nya dikirim via channel, misal akan diproses di routine yang lain sebagai chatbot
					select {
					case t.msg <- update.Message.Text:
						//send channel
					case <-ctx.Done():
						return ctx.Err()
					}
				}*/

			}

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (t *teleRepo) OnMsg() chan string {
	return t.msg
}
