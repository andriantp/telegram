package main

import (
	"log"
	"os"
	"telegram/telegram"
	"time"
)

func main() {
	token := os.Getenv("token")
	if token == "" {
		log.Fatalf("token was required")
	}
	setting := telegram.Setting{
		Token:  token,
		ChatID: -1, // for filter chatbot
	}

	repo, err := telegram.Newtelegram(setting)
	if err != nil {
		log.Fatalf("Newtelegram:%v", err)
	}
	// log.Println("succesed Authorized account")
	time.Sleep(1 * time.Second)

	log.Println("===== Run  =====")
	if err = Run(repo); err != nil {
		log.Fatalf("Run:%v", err)
	}
}
