package sender

import (
	"os"
	"testing"

	"github.com/gozap/msgsend/pkg/conf"

	"github.com/spf13/viper"
)

func initConfig() {
	viper.Set("telegram.api", os.Getenv("telegram_api"))
	viper.Set("telegram.token", os.Getenv("telegram_token"))
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
	err = bot.SendMessage("Hello World", []conf.TGRecipient{
		{
			ID:   111111111,
			Type: "user",
		},
		{
			ID:   -111111111,
			Type: "group",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
}
