package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var version, buildDate, commitID string

var versionTpl = `Name: msgsend
Version: %s
Arch: %s
BuildDate: %s
CommitID: %s
`

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(versionTpl, version, runtime.GOOS+"/"+runtime.GOARCH, buildDate, commitID)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
