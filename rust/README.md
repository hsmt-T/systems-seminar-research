# Rust Axum + PostgreSQL

Go Echo との比較用に、ORM なしで Axum と PostgreSQL だけを Docker Compose で起動する最小構成です。

## 起動

```sh
docker compose up --build
```

## 動作確認

```sh
curl http://localhost:3000/health
curl http://localhost:3000/db-health
```

`/db-health` は `tokio-postgres` で `select 1` を直接実行します。ORM やマイグレーションツールは入れていません。
PostgreSQL は Compose 内部ネットワークだけに公開しています。ホスト側から直接接続したい場合は、`compose.yaml` の `db` に空いているポートを追加してください。

## 構成

- `compose.yaml`: Axum アプリと PostgreSQL の起動設定
- `Dockerfile`: Rust アプリのビルドと実行
- `src/main.rs`: 最小限の HTTP サーバーと DB 接続確認
