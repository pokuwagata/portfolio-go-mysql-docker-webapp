# bff-todo

## 構成

- Front - (GraphQL) - BFF - REST API Server
- BFF : TypeScript + Apollo Server

## 進め方

### MVP

- BFF 用コンテナを立ち上げ
- サーバーをセットアップ
- docker-compose に組み入れて疎通確認
- 1 リクエスト分のスキーマを実装
- フロントに Apollo Client を導入
- モックに対してクエリを発行
- リゾルバを実装して API とつなぎこみ

### 置き換え

- 認証
- ページネーション
- その他

### モバイル対応

- モバイル用 UI を作成

## 気になる点

- Docker では、ts-node でサーバーを立ち上げるべきなのか？ tsc + node で立ち上げるべきなのか？
  - コンパイル結果をキャッシュできる
- 開発時のコード監視＋反映はどう組むべきか？
  - nodemon + ts-node
    - nodemon の 1 プロセスで済むので楽
    - <https://stackoverflow.com/questions/37979489/how-to-watch-and-reload-ts-node-when-typescript-files-change>
  - tsc watch → nodemon
    - 2 プロセスの起動が必要
- nodemon はなぜ必要？　変更監視＋再起動がこのライブラリでしかできないから
- `express()`関数は`import express from 'express'`のデフォルトインポートでしか読み込めない
- class 構文で Express をラップして書く流儀があるっぽい
  - <https://github.com/mwanago/express-typescript/blob/master/src/app.ts>
- これが一番参考になるのかも？ <https://github.com/microsoft/TypeScript-Node-Starter/blob/master/src/app.ts>
- .graphql ファイルを ES2015 import で読み込むことはできるのか？ → このままでは無理
  - Node.js(Common.js)の readfileSync なら可能っぽいが、import では無理
  - .graphql ファイルでシンタックスハイライトしたい、それを import して Apollo Server に渡したい
    - <https://github.com/ardatan/graphql-import> が解決する
      - 解決しなかった。そもそも tsc 単体では規定外のファイルをモジュールとして扱えないのではないか？
      - 型定義ファイルがあれば何とかなる説 → 試したが駄目だった
        - <https://github.com/apollographql/graphql-tag/issues/59>
      - 型定義ファイル＋ webpack で解決する説
        - <https://dev.to/open-graphql/how-to-resolve-import-for-the-graphql-file-with-typescript-and-webpack-35lf>
      - `graphql-import-node`を試す
        - <https://github.com/ardatan/graphql-import-node>
        - **TypeScript上で動作した！！**
        - 結局ソースを読むとreadfilesyncでラップしてるだけ？　でも単純にreadfilesyncするだけだとtscがモジュールを解決しない
        - [ ] なぜこれを使うと動くんだろう？
      - `graphql-import`を試す
        - <https://github.com/ardatan/graphql-import/issues/267>
          - なんか駄目っぽい？
      - `babel-plugin-import-graphql`
        - babelの利用が前提
      - 
- tsconfig.json の module, target は変換先の設定
  - <https://stackoverflow.com/questions/41993811/understanding-target-and-module-in-tsconfig>
