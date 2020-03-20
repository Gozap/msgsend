package conf

type Telegram struct {
	Api       string        `mapstructure:"api"`
	Token     string        `mapstructure:"token"`
	Recipient []TGRecipient `mapstructure:"recipient"`
}

type TGRecipient struct {
	ID   int64  `mapstructure:"id"`
	Type string `mapstructure:"type"`
}

func Example() string {
	return `
telegram:
  api: https://api.telegram.org
  token: asdadsasddfgdfgdf
  recipient:
    - id: 11111
      type: user
    - id: -2222
      type: group
`

}
