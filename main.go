package main

import (
	"fmt"
	"os"

	"github.com/bun913/backlog-backup-sample/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{}
	rootCmd.AddCommand(cmd.IssueFileCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}
}
