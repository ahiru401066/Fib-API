# Fib_api
2025.04 ~ フィボナッチ数を返すAPIサーバーを作成

### 処理概要説明
　このAPIサーバーはクエリパラメータとして指定された整数nに対してフィボナッチ数を計算し、JSON形式でクライアントに返却するものです。主な処理の流れは以下の通りです。また、今回は引数によっては返り値がかなり大きくなることが予想されるので文字列で返却するという判断をしています。

1\. /fib?n={number}に対するHTTPリクエストを受け取る  
2\. リクエストのメゾットの確認、クエリパラメータへのバリデーション処理を実施  
3\. {number}に対するフィボナッチ数を計算  
4\. 計算結果をJSON形式でクライアントに返却  


#### エンドポイントURL
https://fib-api-spds.onrender.com/  
renderの無料枠を利用しているため最初のアクセスは時間がかかります。


### ディレクトリ構成
<details><summary>構成を見る</summary>
<pre>
.
├── backend
│   ├── cmd
│   │   └── api-server
│   │       └── main.go
│   ├── Dockerfile             本番環境用
│   ├── go.mod
│   └── internal
│       ├── handler
│       │   ├── fib_test.go    テストコード
│       │   └── fib.go
│       ├── model
│       │   ├── error.go
│       │   └── fib.go
│       ├── service
│       │   ├── fib_test.go　　テストコード
│       │   └── fib.go
│       └── utils
│           └── error.go
├── docker-compose.yml         開発環境用
├── Dockerfile　　　　　　　　　　開発環境用
└── README.md
</pre>
</details>

### フォルダ・ファイルの説明

#### backend/
　本番環境用のDockerfileとAPIサーバーのソースコードが存在

#### backend/cmd/api-server/main.go
　エントリーポイント、HTTPルーティング、サーバーの起動処理を行う。

#### backend/internal/handler/fib.go
"/fib"に対応する処理の実装  
　/fib?n={number}に対して、リクエストメゾット、クエリパラメータの確認、numberに対してバリデーションを行い、service/fib.goへ適切な変数を与える。fib.goでの処理が終了後、計算結果となるデータをクライアントへjson形式で返す。　　　　　　　　　　　　　　　　　　　　　　　　　

#### internal/model/
- error.go：エラーレスポンスの構造体を定義
- fib.go:fib計算結果のレスポンスの構造体を定義

#### service/fib.go
　フィボナッチ数の計算ロジック処理の実装

### utils
　errorレスポンス処理の実装
