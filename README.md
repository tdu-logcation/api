# Logcation API

LogcationのAPIです。

- ログデータのクラウド保存
- ログデータの複数端末同期
- オンラインランキング

を提供します。

## API仕様

### アカウント作成

```text
POST https://api.tdu.app/user
Content-Type: application/x-www-form-urlencoded

user_name=[user name]
```

- form
  - user_name: ユーザ名

### ユーザ情報取得

```text
GET https://api.tdu.app/user?id=[id]
```

- query
  - id: アカウント作成時に返ってくるid

### ユーザ名変更

```text
POST https://api.tdu.app/user
Content-Type: application/x-www-form-urlencoded

id=[id]&user_name=[new user name]
```

- form
  - id: アカウント作成時に返ってくるid
  - user_name: 新しいユーザ名

### アカウント（ログ含め）削除

```text
DELETE https://api.tdu.app/user?id=[id]
```

- query
  - id: アカウント作成時に返ってくるid

### ログ取得

```text
GET https://api.tdu.app/log?id=[id]
```

- query
  - id: アカウント作成時に返ってくるid

### ログ追加

```text
POST https://api.tdu.app/log
Content-Type: application/json

{
    id: [id]
    logs: [
        {
            date: [date],
            campus: [campus],
            log_type: [log type],
            label: [label],
            code: code
        }
    ]
}
```

- form
  - id: アカウント作成時に返ってくるid
  - date: ログ取得日時（RFC3339形式）
  - campus: キャンパス
  - log_type: ログの種類
  - label: ラベル
  - code: ログ

### ランキング

```text
GET https://api.tdu.app/rank
```

## Testing

### Local

[Datasotoreエミュレータ](https://cloud.google.com/datastore/docs/tools/datastore-emulator#linux-macos)をインストールしてください。

```bash
# datastoreのエミュレータを実行
gcloud beta emulators datastore start --data-dir=.

# 環境変数設定
$(gcloud beta emulators datastore env-init)

# datasoreをクリーンアップ
rm -rf ./WEB_INF
```

## LICENSE

[MIT](./LICENSE)
