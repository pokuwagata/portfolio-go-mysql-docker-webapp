# bff-todo

## 構成

- Front - (GraphQL) - BFF - REST API Server
- BFF : TypeScript + Apollo Server

## 進め方

### MVP

- BFF用コンテナを立ち上げ
- サーバーをセットアップ
- docker-composeに組み入れて疎通確認
- 1リクエスト分のスキーマを実装
- フロントにApollo Clientを導入
- モックに対してクエリを発行
- リゾルバを実装してAPIとつなぎこみ

### 置き換え

- 認証
- ページネーション
- その他

### モバイル対応

- モバイル用UIを作成

## 気になる点

- Dockerでは、ts-nodeでサーバーを立ち上げるべきなのか？ tsc + nodeで立ち上げるべきなのか？
  - コンパイル結果をキャッシュできる
- 開発時のコード監視＋反映はどう組むべきか？
  - nodemon + ts-node
    - nodemonの1プロセスで済むので楽
    - <https://stackoverflow.com/questions/37979489/how-to-watch-and-reload-ts-node-when-typescript-files-change>
  - tsc watch → nodemon
    - 2プロセスの起動が必要
- nodemonはなぜ必要？　変更監視＋再起動がこのライブラリでしかできないから
- `express()`関数は`import express from 'express'`のデフォルトインポートでしか読み込めない
