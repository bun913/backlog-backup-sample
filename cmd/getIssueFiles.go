package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cobra"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type attachmentInfo struct {
	id   string
	name string
}

// getIssueFilesCmd represents the getIssueFiles command
var getIssueFilesCmd = &cobra.Command{
	Use:     "getIssueFiles",
	Aliases: []string{"gif"},
	Short:   "get files attached space's issues",
	Long:    `get files attached space's issues.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CSVの読み込み開始")
		f := openExportedCsv()
		defer f.Close()
		r := csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))
		// 最初の行は無視して読み込み
		fmt.Println("CSVの読み込み終了")
		r.Read()
		for {
			records, err := r.Read()
			if err == io.EOF {
				fmt.Println("終了しました")
				break
			} else if err != nil {
				log.Fatal(err)
			}
			issue := records[4]
			fmt.Printf("課題:%sの添付ファイルを確認開始\n", issue)
			client := NewRequestClient()
			fileList := getAttachedFileList(client, issue)
			for _, attachment := range fileList {
				fmt.Printf("    添付ファイル%sのダウンロード開始\n", attachment.name)
				downLoadFile(client, issue, attachment)
				fmt.Printf("    添付ファイル%sのダウンロード終了\n", attachment.name)
			}
		}
	},
}

func openExportedCsv() *os.File {
	cf, err := os.OpenFile(CsvFile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return cf
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
		log.Fatalln("DownLoad IssueAttachedFile Fail", err)
	}
}

func init() {
	rootCmd.AddCommand(getIssueFilesCmd)
	// csvFile
	getIssueFilesCmd.PersistentFlags().
		StringVarP(&CsvFile, "csv", "c", "", "(Required) exported csv(https://support-ja.backlog.com/hc/ja/articles/360035642534-%E8%AA%B2%E9%A1%8C%E6%A4%9C%E7%B4%A2%E7%B5%90%E6%9E%9C%E4%B8%80%E8%A6%A7%E3%81%AE%E5%87%BA%E5%8A%9B)")
	getIssueFilesCmd.MarkPersistentFlagRequired("csv")
}
