/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/package cmd

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cobra"
)

// Command Args
var ProjectID string
var ApiKey string
var SpaceID string

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

// Return New RestyClient setted apikey
func NewRequestClient() *resty.Client {
	client := resty.New()
	client.BaseURL = getBaseUrl()
	client.SetQueryParam("apiKey", ApiKey)
	// json-iteratorをデフォルトのJSOnクライアントに設定
	client.JSONMarshal = jsoniter.Marshal
	return client
}

// Get Backlog API BaseURL
func getBaseUrl() string {
	baseURL := fmt.Sprintf("https://%s.backlog.com", SpaceID)
	return baseURL
}

func init() {
	// 必須フラグのチェック
	// SpaceID
	rootCmd.PersistentFlags().StringVarP(&SpaceID, "space", "s", "", "(Required) Your space id")
	rootCmd.MarkPersistentFlagRequired("space")
	// ProjectID
	rootCmd.PersistentFlags().
		StringVarP(&ProjectID, "project", "p", "", "(Required) Your project id")
	rootCmd.MarkPersistentFlagRequired("project")
	ProjectID, _ = rootCmd.PersistentFlags().GetString("project")
	// apiKey
	rootCmd.PersistentFlags().StringVarP(&ApiKey, "apikey", "a", "", "(Required) Your api key")
	rootCmd.MarkPersistentFlagRequired("apikey")
}
