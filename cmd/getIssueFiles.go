/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getIssueFilesCmd represents the getIssueFiles command
var getIssueFilesCmd = &cobra.Command{
	Use:     "getIssueFiles",
	Aliases: []string{"gif"},
	Short:   "get files attached project's issues",
	Long:    `get files attached project's issues.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getIssueFiles called")
	},
}

func init() {
	rootCmd.AddCommand(getIssueFilesCmd)
	getIssueFilesCmd.Flags().
		StringVarP(&ProjectID, "project", "p", "", "(Required) Your project id")
	getIssueFilesCmd.MarkFlagRequired("project")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getIssueFilesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getIssueFilesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
