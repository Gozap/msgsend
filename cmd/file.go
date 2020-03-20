package cmd

import (
	"path/filepath"

	"github.com/gozap/msgsend/pkg/conf"
	"github.com/gozap/msgsend/pkg/sender"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var fileName string
var mime string
var caption string

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
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

		if fileName == "" {
			fileName = filepath.Base(args[0])
		}
		err = bot.SendFile(args[0], fileName, mime, caption, to)
		if err != nil {
			logrus.Fatal(err)
		}

	},
}

func init() {
	fileCmd.PersistentFlags().StringVarP(&fileName, "name", "n", "", "filename")
	fileCmd.PersistentFlags().StringVarP(&mime, "mime", "m", "", "file mime type")
	fileCmd.PersistentFlags().StringVarP(&caption, "caption", "c", "", "file caption")
	rootCmd.AddCommand(fileCmd)
}
