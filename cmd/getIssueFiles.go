/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cobra"
)

type attachmentInfo struct {
	id   string
	name string
}

// getIssueFilesCmd represents the getIssueFiles command
var getIssueFilesCmd = &cobra.Command{
	Use:     "getIssueFiles",
	Aliases: []string{"gif"},
	Short:   "get files attached project's issues",
	Long:    `get files attached project's issues.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Issueを直渡ししている部分を修正する
		issue := "EXPORTTEST-215"
		client := NewRequestClient()
		fileList := getAttachedFileList(client, issue)
		for _, attachment := range fileList {
			downLoadFile(client, issue, attachment)
		}
	},
}

func getAttachedFileList(client *resty.Client, issue string) []attachmentInfo {
	// https://developer.nulab.com/ja/docs/backlog/api/2/get-list-of-issue-attachments/
	resp, err := client.R().
		Get(client.BaseURL + fmt.Sprintf("/api/v2/issues/%s/attachments", issue))
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode() != http.StatusOK {
		log.Fatalln("Request Fail:" + resp.String())
	}
	fileListResponse := jsoniter.Get(resp.Body())
	fileList := fileListResponse.Get()
	var attachmentList []attachmentInfo
	for i := 0; i < fileList.Size(); i++ {
		// ダウンロード
		elm := fileList.Get(i)
		attachmenID := elm.Get("id").ToString()
		attachmentName := elm.Get("name").ToString()
		ai := attachmentInfo{id: attachmenID, name: attachmentName}
		attachmentList = append(attachmentList, ai)
	}
	return attachmentList
}

func downLoadFile(client *resty.Client, issue string, attachement attachmentInfo) {
	// NOTE: 添付ファイルの保存先はユーザーに選んでもらえればなお良い
	baseDir, _ := os.Getwd()
	attachmentFileDir := "attachedFiles"
	outpuDir := path.Join(baseDir, attachmentFileDir, issue)
	os.MkdirAll(outpuDir, os.ModePerm)
	// https://developer.nulab.com/ja/docs/backlog/api/2/get-issue-attachment/#
	url := client.BaseURL + fmt.Sprintf("/api/v2/issues/%s/attachments/%s", issue, attachement.id)
	_, err := client.R().SetOutput(path.Join(outpuDir, attachement.name)).Get(url)
	if err != nil {
		log.Fatalln("DownLoad AttachedFile Fail")
	}
}

func init() {
	rootCmd.AddCommand(getIssueFilesCmd)
}
