package cmd

import (
	"io/ioutil"
	"strings"

	"github.com/gozap/msgsend/pkg/conf"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/gozap/msgsend/pkg/sender"
	"github.com/spf13/cobra"
)

var txtFile string

var txtCmd = &cobra.Command{
	Use:    "txt",
	Short:  "A brief description of your command",
	PreRun: func(_ *cobra.Command, _ []string) { initConfig() },
	Run: func(cmd *cobra.Command, args []string) {

		var msg string
		if len(args) > 0 {
			msg = strings.Join(args, " ")
		} else if txtFile != "" {
			bs, err := ioutil.ReadFile(txtFile)
			if err != nil {
				logrus.Fatal(err)
			}
			msg = string(bs)
		} else {
			_ = cmd.Help()
			return
		}

		bot, err := sender.NewTelegram()
		if err != nil {
			logrus.Fatal(err)
		}

		var to []conf.TGRecipient
		err = viper.UnmarshalKey("telegram.recipient", &to)
		if err != nil {
			logrus.Fatal(err)
		}

		err = bot.SendMessage(msg, to)
		if err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	txtCmd.PersistentFlags().StringVarP(&txtFile, "file", "f", "", "text message file")
	rootCmd.AddCommand(txtCmd)
}
