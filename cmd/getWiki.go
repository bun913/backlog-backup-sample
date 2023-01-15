/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cobra"
)

type wikiInfo struct {
	id   string
	name string
}

type wikiAttachment struct {
	id   string
	name string
}

type wikiDetail struct {
	id          string
	content     string
	attachments []wikiAttachment
}

// getWikiCmd represents the getWiki command
var getWikiCmd = &cobra.Command{
	Use:   "wiki",
	Short: "get wiki contents.",
	Long:  `get wiki contents.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := NewRequestClient()
		fmt.Println("Wiki一覧情報の取得開始")
		wikiList := getWikiList(client)
		for _, w := range wikiList {
			// Wikiページの詳細取得
			client = NewRequestClient()
			fmt.Println("ページ情報の取得開始:", w.name)
			detail := getWikiDetail(client, w.id)
			dir := getContentDir(w)
			// Wikiのコンテンツをローカルに保存する
			os.MkdirAll(dir, os.ModePerm)
			cf, err := os.Create(path.Join(dir, "content.md"))
			if err != nil {
				log.Fatalln("Wiki content file create error: ", err)
			}
			defer cf.Close()
			cf.WriteString(detail.content)
			// 添付ファイルのダウンロード
			for _, attachment := range detail.attachments {
				fmt.Println(fmt.Sprintf("    添付ファイルのダウンロード開始:%s", attachment.name))
				client = NewRequestClient()
				downLoadWikiFile(client, detail, attachment, dir)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(getWikiCmd)
	// Project IDは必須
	getWikiCmd.PersistentFlags().
		StringVarP(&ProjectID, "project", "p", "", "(Required) Your Project ID")
	getWikiCmd.MarkPersistentFlagRequired("project")
}

// Wikiページの一覧を取得
func getWikiList(client *resty.Client) []wikiInfo {
	// https://developer.nulab.com/ja/docs/backlog/api/2/get-wiki-page-list/#
	client.SetQueryParam("projectIdOrKey", ProjectID)
	resp, err := client.R().
		Get(client.BaseURL + "/api/v2/wikis")
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode() != http.StatusOK {
		log.Fatalln("Request Fail:" + resp.String())
	}
	fileListResponse := jsoniter.Get(resp.Body())
	fileList := fileListResponse.Get()
	var wikiList []wikiInfo
	for i := 0; i < fileList.Size(); i++ {
		elm := fileList.Get(i)
		wikiID := elm.Get("id").ToString()
		wikiName := elm.Get("name").ToString()
		wi := wikiInfo{id: wikiID, name: wikiName}
		wikiList = append(wikiList, wi)
	}
	return wikiList
}

// Wikiページの詳細情報を取得
func getWikiDetail(client *resty.Client, wikiID string) wikiDetail {
	resp, err := client.R().
		Get(client.BaseURL + fmt.Sprintf("/api/v2/wikis/%s", wikiID))
	// TODO: この辺共通化できそう
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode() != http.StatusOK {
		log.Fatalln("Get WikiDetail Request Fail:" + resp.String())
	}
	fileListResponse := jsoniter.Get(resp.Body())
	elm := fileListResponse.Get()
	wikiContent := elm.Get("content").ToString()
	fileList := elm.Get("attachments")
	// Wikiページの添付ファイル情報を取得
	var attachmentList []wikiAttachment
	for i := 0; i < fileList.Size(); i++ {
		elm := fileList.Get(i)
		attachmenID := elm.Get("id").ToString()
		attachmentName := elm.Get("name").ToString()
		wa := wikiAttachment{id: attachmenID, name: attachmentName}
		attachmentList = append(attachmentList, wa)
	}
	wd := wikiDetail{id: wikiID, content: wikiContent, attachments: attachmentList}
	return wd
}

// Wikiコンテンツを保存するディレクトリパスの作成
func getContentDir(wi wikiInfo) string {
	cwd, _ := os.Getwd()
	baseDir := filepath.Join(cwd, "wiki")
	wikiPathList := strings.Split(wi.name, "/")
	wikiPath := filepath.Join(wikiPathList...)
	dir := filepath.Join(baseDir, wikiPath)
	return dir
}

// Wiki添付ファイルのダウンロード
func downLoadWikiFile(
	client *resty.Client,
	detail wikiDetail,
	attachment wikiAttachment,
	dir string,
) {
	// https://developer.nulab.com/ja/docs/backlog/api/2/get-wiki-page-attachment/#
	url := fmt.Sprintf(
		client.BaseURL+
			"/api/v2/wikis/%s/attachments/%s",
		detail.id,
		attachment.id,
	)
	_, err := client.R().SetOutput(path.Join(dir, attachment.name)).Get(url)
	if err != nil {
		log.Fatalln("DownLoad WikiAttachedFile Fail ", err)
	}

}
