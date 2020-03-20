package conf

type Config struct {
	Telegram Telegram `mapstructure:"telegram"`
}

type Telegram struct {
	Api       string        `mapstructure:"api"`
	Token     string        `mapstructure:"token"`
	Recipient []TGRecipient `mapstructure:"recipient"`
}

type TGRecipient struct {
	ID   int64  `mapstructure:"id"`
	Type string `mapstructure:"type"`
}
