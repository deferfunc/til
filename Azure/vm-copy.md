## Azure でのVM作成について

できるだけ自動化したいけど・・・以下目的で調査する

* Azure Portal でできることをコマンドラインで
* 自動化の仕組みを使ってワンクリックかPRマージ等のワンアクションでデプロイまで完了させたい
   * まずは、入れ替えようのVMの作成までを目標に

### az コマンドで

スナップショットがある状態で

指定は個人的に使用するものをチョイス

```bash
# 管理ディスクを作成する
az disk create --resource-group RESOURCE_GROUP --tags environment=VALUE --name DISK_NAME --zone 1 --public-network-access Disabled --source SNAPSHOT_ID
```

管理ディスクからVMイメージを作成する

* 正常性モニターは設定できない→あとから手動で
* 推奨される監視項目も設定できなし→あとから手動で
* パブリックIP無し
* 削除時のNIC削除をON
* ネットワークセキュリティグループは無し（サブネットにセキュリティグループを設定している前提）
* `--custom-data` オプションで hostname 指定を試みたが期待通りいかず

```bash
az vm create --resource-group RESOURCE_GROUP --name VM_NAME --attach-os-disk DISK_NAME  --os-type linux --size Standard_F8s_v2 --license-type None   --nic-delete-option Delete --tags environment=VALUE --public-ip-address "" --nsg "" --vnet-name VNET_NAME --subnet SUBNET_NAME --zone 1
```
