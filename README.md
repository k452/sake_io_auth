# さけ.io の 認証API

## 使用技術

- Go
  - Echo
  - realize
  - logrus
  - godotenv
- bbolt

## 開発環境

- Docker
- Docker-Compose

## ディレクトリ構成

├── README.md  
├── conf --- 環境変数の読み込み  
│   └── dbConfig.go  
├── db --- db 接続  
│   └── connect.go  
├── docker-compose.yml  
├── Dockerfile  
├── envFile --- 環境変数ファイル  
│   ├── production.env  
│   └── test.env  
├── go.mod  
├── go.sum  
├── main.go --- API サーバー起動  
├── model  
│   └── main.go  
├── routing --- ルーティング  
│   └── main.go 

