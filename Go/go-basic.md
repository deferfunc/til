## go.mod バージョン更新

```
go mod tidy -go=1.22
```

上記コマンドで go.mod の go ディレクティブのバージョン番号更新する


## cloud function の go ランタイム

[ランタイム サポート  |  Cloud Functions Documentation  |  Google Cloud](https://cloud.google.com/functions/docs/runtime-support?hl=ja#go)

```
Go 1.22	第 2 世代	Ubuntu 22.04	go122	gcr.io/gae-runtimes/buildpacks/go122/run
```

## ログ周り

* Go1.21 から slog が標準ライブラリーに
* lumberjack と組み合わせると良い、例えばログローテーションができたりする
* slog の使い方参考資料: [🪵 Go1.21 log/slogパッケージ超入門](https://zenn.dev/88888888_kota/articles/7e97ff874083cf)
