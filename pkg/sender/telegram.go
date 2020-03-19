package sender

import (
	"errors"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Telegram struct {
	bot *tb.Bot
}

type TGRecipient struct {
	ID   int64
	Type string
}

func NewTelegram() (*Telegram, error) {
	bot, err := tb.NewBot(tb.Settings{
		URL:    viper.GetString("telegram.api"),
		Token:  viper.GetString("telegram.token"),
		Poller: &tb.LongPoller{Timeout: 5 * time.Second},
	})
	if err != nil {
		return nil, err
	} else {
		return &Telegram{bot: bot}, nil
	}
}

func (tg *Telegram) SendMessage(msg string, recipient []TGRecipient) error {

	if len(recipient) == 0 {
		err := viper.UnmarshalKey("telegram.recipient", &recipient)
		if err != nil {
			return err
		}
	}

	var wg sync.WaitGroup
	wg.Add(len(recipient))

	for _, r := range recipient {
		var to tb.Recipient
		switch r.Type {
		case "user":
			to = &tb.User{ID: int(r.ID)}
		case "group":
			to = &tb.Chat{ID: r.ID}
		default:
			return errors.New("invalid recipient type")
		}
		go func() {
			defer wg.Done()
			_, err := tg.bot.Send(to, msg)
			if err != nil {
				logrus.Error(err)
			}
		}()
	}

	wg.Wait()

	return nil
}
