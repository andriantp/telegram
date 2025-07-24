package telegram

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"testing"
)

var setting = Setting{
	Token:  "",
	ChatID: -1,
}

func Test_Connect(t *testing.T) {
	_, err := Newtelegram(setting)
	if err != nil {
		t.Fatalf("Newtelegram:%v", err)
	}

}

func Test_SendText(t *testing.T) {
	ctx := context.Background()

	repo, err := Newtelegram(setting)
	if err != nil {
		t.Fatalf("Newtelegram:%v", err)
	}

	err = repo.SendText(ctx, "test")
	if err != nil {
		t.Fatalf("SendText:%v", err)
	}
}

func Test_SendImagePathx1A(t *testing.T) {
	ctx := context.Background()

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	log.Printf("basepath:%s", basepath)
	base := basepath[0 : strings.LastIndex(basepath, "/")+1]
	// log.Printf("base:%s", base)

	img := "image/"
	filename := "gotele"
	path := fmt.Sprintf("%s/%s.jpg", filepath.Join(base, img), filename)
	log.Printf("path:%s", path)

	repo, err := Newtelegram(setting)
	if err != nil {
		t.Fatalf("Newtelegram:%v", err)
	}

	if err := repo.SendImageFromFile(ctx, filename, path); err != nil {
		t.Fatalf("SendImageFromFile:%v", err)
	}
}
