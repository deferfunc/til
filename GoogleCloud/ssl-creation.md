# SSL作成に使うコマンド

cloud shell の前提

現状のものの確認（紐づいたドメインを見る）

```
gcloud beta compute ssl-certificates describe my-ssl --format="get(managed.domains)"
```

新規作成

```
gcloud beta compute ssl-certificates create new-my-ssl --domains=new.example.com,blog.example.com
```

ロードバランサーの更新

```
gcloud compute target-https-proxies update target-proxy-name --ssl-certificates=new-my-ssl
```

ここで、target-proxy-name は更新するロードバランサーの HTTPS プロキシの名前を指定します。