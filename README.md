# Logcation API

LogcationのAPIです。

- ログデータのクラウド保存
- ログデータの複数端末同期
- オンラインランキング

を提供します。

## Routes

- /user
  - POST
    - アカウント作成
    - ユーザ名設定（変更）
  - GET
    - ユーザ情報取得
      - 名前とかログ数とか？
  - DELETE
    - アカウント削除
- /log
  - POST
    - ログ追加
  - GET
    - ログ取得
- /rank
  - GET
    - ランキング取得

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
