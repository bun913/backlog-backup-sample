# Backlog バックアップサンプル

お金をかけずにBacklogにおける特定のプロジェクトは以下の **課題**と**Wiki**の情報を取得するためのサンプルです。

現段階では以下の機能を実装しています。

- Backlogから出力した課題一覧のCSVから課題に添付されたファイルを順次ダウンロードしていく

課題の一覧に関しては、[課題検索結果一覧の出力 – Backlog ヘルプセンター](https://support-ja.backlog.com/hc/ja/articles/360035642534-%E8%AA%B2%E9%A1%8C%E6%A4%9C%E7%B4%A2%E7%B5%90%E6%9E%9C%E4%B8%80%E8%A6%A7%E3%81%AE%E5%87%BA%E5%8A%9B)の機能を利用してcsvとして出力する前提です。

このサンプルプロジェクトでは以下の機能を提供します。

- 課題に添付されたファイルをダウンロードする
- Wikiの内容と添付されたファイルをダウンロードする

詳細・経緯は以下ブログをご参照ください。

TODO: 後でブログへのリンクを貼る

## 必要要件

- Backlog APIを実行するためのAPIキーを取得している
- 課題の一覧をcsvとしてダウンロード済み
- スペースのIDを把握している

## 使い方

先に本プロジェクトを git cloneしたうえで `go build`により実行ファイルをビルドする必要があります。（余裕があればreleaseパッケージに自動追加する処理を作りたい）

```
go build -ldflags="-s -w" -trimpath
```

## 共通の注意事項

[BacklogAPIのレート制限](https://developer.nulab.com/ja/docs/backlog/rate-limit/#)による制限を受けた際、65秒ほど待機する処理を実装しています。

そのため処理が途中で止まる場面があるかと思いますが、そのまま待機してください。

### 課題の添付ファイル取得処理

```
./backlog-backup-sample gif -a ${APIKey} --space ${SpaceID} -c ${BackLogから出力したCSVファイルのパス}
```

以下のように処理を開始します

```
CSVの読み込み終了
課題:EXPORTTEST-215の添付ファイルを確認開始
    添付ファイル1m7wlVHg_400x400.jpegのダウンロード開始
    添付ファイル1m7wlVHg_400x400.jpegのダウンロード終了
課題:EXPORTTEST-213の添付ファイルを確認開始
    添付ファイルcomment.csvのダウンロード開始
    添付ファイルcomment.csvのダウンロード終了
課題:EXPORTTEST-214の添付ファイルを確認開始
終了しました
```


