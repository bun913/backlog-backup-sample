# Backlog バックアップサンプル

お金をかけずにBacklogにおける特定のプロジェクトは以下の **課題**と**Wiki**の情報を取得するためのサンプルです。

課題の一覧に関しては、[課題検索結果一覧の出力 – Backlog ヘルプセンター](https://support-ja.backlog.com/hc/ja/articles/360035642534-%E8%AA%B2%E9%A1%8C%E6%A4%9C%E7%B4%A2%E7%B5%90%E6%9E%9C%E4%B8%80%E8%A6%A7%E3%81%AE%E5%87%BA%E5%8A%9B)の機能を利用してcsvとして出力する前提です。

このサンプルプロジェクトでは以下の機能を提供します。

- 課題に添付されたファイルをダウンロードする
- Wikiの内容と添付されたファイルをダウンロードする

詳細・経緯は以下ブログをご参照ください。

TODO: 後でブログへのリンクを貼る

## 必要要件

- Backlog APIを実行するためのAPIキーを取得している

## 使い方

TODO: 後で書き直す。今はテキトーにかいているだけ

1. `go build -ldflags="-s -w" -trimpath -o ` でビルド

