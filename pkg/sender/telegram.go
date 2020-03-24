package sender

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gozap/msgsend/pkg/conf"

	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Telegram struct {
	bot *tb.Bot
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

func (tg *Telegram) SendMessage(msg string, recipient []conf.TGRecipient, markdown bool) error {

	opt := &tb.SendOptions{}
	if markdown {
		opt.ParseMode = tb.ModeMarkdown
	}

	send := func(to tb.Recipient) {
		_, err := tg.bot.Send(to, msg, opt)
		if err != nil {
			logrus.Error(err)
		}
	}

	return tg.batchSend(send, recipient)
}

func (tg *Telegram) SendFile(filePath, fileName, mime, caption string, recipient []conf.TGRecipient) error {

	send := func(to tb.Recipient) {
		_, err := tg.bot.Send(to, &tb.Document{
			File:     tb.FromDisk(filePath),
			Caption:  caption,
			MIME:     mime,
			FileName: fileName,
		})
		if err != nil {
			logrus.Error(err)
		}
	}

	return tg.batchSend(send, recipient)
}

func (tg *Telegram) SendImage(imagePath, caption string, recipient []conf.TGRecipient) error {

	send := func(to tb.Recipient) {
		_, err := tg.bot.Send(to, &tb.Photo{
			File:    tb.FromDisk(imagePath),
			Caption: caption,
		})
		if err != nil {
			logrus.Error(err)
		}
	}

	return tg.batchSend(send, recipient)
}

func (tg *Telegram) batchSend(send func(to tb.Recipient), recipient []conf.TGRecipient) error {
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
			send(to)
		}()
	}

	wg.Wait()
	return nil
}

func (tg *Telegram) ID() {
	tg.bot.Handle("/id", func(m *tb.Message) {
		var r tb.Recipient
		if m.Chat != nil {
			r = m.Chat
		} else {
			r = m.Sender
		}
		_, err := tg.bot.Send(r, fmt.Sprintln(m.ID))
		if err != nil {
			logrus.Error(err)
		}
	})
}

func (tg *Telegram) Start() {
	tg.bot.Start()
}
