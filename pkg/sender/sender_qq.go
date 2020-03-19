package sender

import (
	"strconv"

	"github.com/spf13/viper"

	qqbotapi "github.com/catsworld/qq-bot-api"
)

type QQ struct {
	bot *qqbotapi.BotAPI
}

func (qq *QQ) SendMessage(msg, recipient, recipientType string, _ ...interface{}) error {

	recipientID, err := strconv.Atoi(recipient)
	if err != nil {
		return err
	}

	_, err = qq.bot.SendMessage(int64(recipientID), recipientType, msg)
	return err
}

func NewSenderQQ() (*QQ, error) {
	token := viper.GetString("qq.token")
	secret := viper.GetString("qq.secret")
	address := viper.GetString("qq.address")

	bot, err := qqbotapi.NewBotAPI(token, address, secret)
	if err != nil {
		return nil, err
	}
	return &QQ{bot: bot}, nil
}
