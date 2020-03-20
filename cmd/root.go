package cmd

import (
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string
var debug bool

var rootCmd = &cobra.Command{
	Use:   "msgsend",
	Short: "Message send tool",
	Run:   func(cmd *cobra.Command, args []string) { _ = cmd.Help() },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig, initLog)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "/etc/msgsend.yaml", "config file (default is /etc/msgsend.yaml)")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug mode")
}

func initConfig() {
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal(err)
	}
	logrus.Debug("Using config file:", viper.ConfigFileUsed())
}

func initLog() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}
