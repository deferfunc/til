# Redshift Serverless の初期設定とか使い方

## 用語とか前提知識とか

- 名前空間 namespace
    - 名前空間は、データベースオブジェクトとユーザーのコレクション
    - ストレージ関連の名前空間は、スキーマ、テーブル、ユーザー、またはデータを暗号化するための AWS Key Management Service キーをグループ化
- ワークグループ workgroup
    - ワークグループは、コンピューティングリソース
    - コンピューティング関連のワークグループは、 **RPU** 、VPC サブネットグループ、セキュリティグループなどのコンピューティングリソースをグループ化
- [ワークグループと名前空間 - Amazon Redshift](https://docs.aws.amazon.com/ja_jp/redshift/latest/mgmt/serverless-workgroup-namespace.html)
- RPU
    - 東京リージョンは USD **0.494/h**
        - 8RPU単位で指定できる、8,16,,, という感じ
        - 最低金額が USD **3.952/h**  633円ぐらい(160円換算)
        - 月額31日だと USD **2,940/month** 47万円/月、フルで使うとなるとちょっと高いか
        - [料金 - Amazon Redshift | AWS](https://aws.amazon.com/jp/redshift/pricing/)
    - 最低60秒で課金
        - You pay for the workloads you run in RPU-hours on a per-second basis, with a 60-second minimum charge.
    - [Billing for Amazon Redshift Serverless - Amazon Redshift](https://docs.aws.amazon.com/redshift/latest/mgmt/serverless-billing.html)
    - 課金額の確認SQL `select trunc(start_time) "Day", (sum(charged_seconds)/3600::double precision) * 0.494 as cost_incurred from sys_serverless_usage group by 1 order by 1`
- RPUを変更するのは時間がかかる、即時反映ではない
    - 一度変更したRPUはすぐには戻せなかった、「スナップショットがないと駄目」みたいなエラー出た

## 参考資料

- [\[アップデート\] Amazon Redshift Serverless の最小ベースキャパシティが32RPUから8RPUに削減されました！ | DevelopersIO](https://dev.classmethod.jp/articles/20230310-amazon-redshift-rpu-8/)
- [Amazon Redshift の継続的なコストパフォーマンス最適化 | Amazon Web Services ブログ](https://aws.amazon.com/jp/blogs/news/big-data-amazon-redshift-continues-its-price-performance-leadership/)
