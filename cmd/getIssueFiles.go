/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	jsoniter "github.com/json-iterator/go"
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
		client := NewRequestClient()
		resp, err := client.R().Get(client.BaseURL + "/api/v2/issues/EXPORTTEST-215/attachments")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(resp.StatusCode())
		fileListResponse := jsoniter.Get(resp.Body())
		// fmt.Println(string(resp.Body()))
		fileList := fileListResponse.Get()
		for i := 0; i < fileList.Size(); i++ {
			fmt.Println("kitayo")
			elm := fileList.Get(i)
			fmt.Println(elm.Get("id").ToString())
		}
	},
}

func init() {
	rootCmd.AddCommand(getIssueFilesCmd)
	// TODO: SpaceIDもパラメーターとして渡したい
	// まずはIssueIDは引数として渡されていると想定する
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getIssueFilesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getIssueFilesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// func GetAttachedFiles(issueIDList: []string) {

// }
