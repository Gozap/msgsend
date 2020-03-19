package sender

import (
	"testing"

	"github.com/spf13/viper"
)

func setConfig() {
	viper.Set("qq.token", "msgsend")
	viper.Set("qq.secret", "msgsend")
	viper.Set("qq.address", "http://172.16.10.6:5700")
}

func TestQQ_SendMessage(t *testing.T) {
	setConfig()
	bot, err := NewSenderQQ()
	if err != nil {
		t.Fatal(err)
	}
	msg := `😁😁😁😁😁😁
南城以南不在蓝，北城以北不在美
缘来缘去终会散，花开花败总归尘
二两桃花酿做酒，万杯不及你温柔
本是青灯不归客，却因浊酒留风尘
三里清风三里路，步步风里再无你`
	err = bot.SendMessage(msg, "2222048", "private")
	if err != nil {
		t.Fatal(err)
	}
	err = bot.SendMessage(msg, "1090266", "group")
	if err != nil {
		t.Fatal(err)
	}
}
