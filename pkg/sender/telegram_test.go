package sender

import (
	"os"
	"strconv"
	"testing"

	"github.com/gozap/msgsend/pkg/conf"

	"github.com/spf13/viper"
)

func initConfig() {
	viper.Set("telegram.api", os.Getenv("TELEGRAM_API"))
	viper.Set("telegram.token", os.Getenv("TELEGRAM_TOKEN"))
}

func TestNewTelegram(t *testing.T) {
	initConfig()
	_, err := NewTelegram()
	if err != nil {
		t.Fatal(err)
	}
}

func TestTelegram_SendMessage(t *testing.T) {
	initConfig()
	bot, err := NewTelegram()
	if err != nil {
		t.Fatal(err)
	}
	userID, _ := strconv.Atoi(os.Getenv("telegram_user"))
	groupID, _ := strconv.Atoi(os.Getenv("telegram_group"))
	err = bot.SendMessage("Hello World", []conf.TGRecipient{
		{
			ID:   int64(userID),
			Type: "user",
		},
		{
			ID:   int64(groupID),
			Type: "group",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestTelegram_SendFile(t *testing.T) {
	initConfig()
	bot, err := NewTelegram()
	if err != nil {
		t.Fatal(err)
	}

	userID, _ := strconv.Atoi(os.Getenv("telegram_user"))
	groupID, _ := strconv.Atoi(os.Getenv("telegram_group"))
	err = bot.SendFile("telegram_test.go", "telegram_test.go", "text/plain", "This is a test file", []conf.TGRecipient{
		{
			ID:   int64(userID),
			Type: "user",
		},
		{
			ID:   int64(groupID),
			Type: "group",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestTelegram_SendImage(t *testing.T) {
	initConfig()
	bot, err := NewTelegram()
	if err != nil {
		t.Fatal(err)
	}

	userID, _ := strconv.Atoi(os.Getenv("telegram_user"))
	groupID, _ := strconv.Atoi(os.Getenv("telegram_group"))
	err = bot.SendImage("/tmp/test.png", "This is a test image", []conf.TGRecipient{
		{
			ID:   int64(userID),
			Type: "user",
		},
		{
			ID:   int64(groupID),
			Type: "group",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestTelegram_ID(t *testing.T) {
	initConfig()
	bot, err := NewTelegram()
	if err != nil {
		t.Fatal(err)
	}

	bot.ID()
	bot.Start()
}
