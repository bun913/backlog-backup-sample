package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RootCmd is root command
var IssueFileCmd = &cobra.Command{
	Use:     "issue-files",
	Short:   "get issue attached files",
	Aliases: []string{"if"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("issueのFileを一覧で取得するためのコマンド")
	},
}
