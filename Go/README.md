# Go Echo + sqlx + PostgreSQL

Echo と sqlx を使った PostgreSQL 接続プール付きの最小構成です。

## 起動

```sh
docker compose up --build
```

## 動作確認

```sh
curl http://localhost:3001/health
curl http://localhost:3001/db-health
```

`/db-health` は起動時に作成した `sqlx.DB` の接続プールを使って `select 1` を実行します。

PostgreSQL は Compose 内の `db` サービスとして起動します。ホスト側のアプリから直接接続したい場合は、`compose.yaml` の `db` にポート公開を追加してください。
