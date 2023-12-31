# backend

## 開発環境

**必要**
- golang 1.20
- docker-compose

**推奨**
- make
- go-migrate
- go-swag


## 起動方法
```bash
# Local環境では自身のマシンのPSQLにつなぐ
make startdb
# サーバの起動
make run

# このように出たら正常
2023/11/11 20:09:46 connected with pq to postgres
2023/11/11 20:09:46 Server listening on port localhost:8080
```

```bash
# dotenvを使う場合、
go run cmd/app/main.go -usedotenv ture -path [任意のパス(初期値はルート)]

# このように出たら正常
2023/11/11 20:09:46 connected with pq to postgres
2023/11/11 20:09:46 Server listening on port localhost:8080
```

## 外部ライブラリ
<details>
<summary>環境変数</summary>

1. github.com/caarlos0/env                      (env読み取り)
2. github.com/joho/godotenv                     (.envを環境変数へ昇格)
</details>

<details>
<summary>DB周り</summary>

1. github.com/lib/pq                            (postgresのドライバ)
</details>

<details>
<summary>HTTP周り</summary>

1. github.com/labstack/echo/v4                  (HTTPサーバ)
</details>

