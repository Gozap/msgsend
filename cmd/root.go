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
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "msgsend.yaml", "config file (default is ./msgsend.yaml)")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug mod")
}

func initConfig() {
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal(err)
	}
	logrus.Debug("Using config file:", viper.ConfigFileUsed())
}
