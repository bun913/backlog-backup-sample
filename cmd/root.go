/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var ProjectID string
var ApiKey string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bl-backup",
	Short: "Sample CLI tool for backup backlog data",
	Long: `Sample CLI tool for backup backlog data
Targets
- Issues attached files
- Wiki
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// 必須フラグの指定
	rootCmd.PersistentFlags().
		StringVarP(&ProjectID, "project", "p", "", "(Required) Your project id")
	rootCmd.MarkPersistentFlagRequired("project")
	rootCmd.PersistentFlags().StringVarP(&ApiKey, "apikey", "a", "", "(Required) Your api key")
	rootCmd.MarkPersistentFlagRequired("apikey")
}
