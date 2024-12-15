# ARM = Azure Resource Manager について

## 調査目的

* VMの作成を対象
* コンソール上の作業のカバー範囲
* 実行方法

## 結果

* ARM テンプレートを作る。ARMテンプレートはJSONファイル。
* JSON 内でパラメーターの置き換えが可能。これはテンプレート式と言う。
    * 例） `"location": "[parameters('location')]",` 
    * 式は大括弧 `[]` で囲む


### ARM テンプレートの作り方

* [Visual Studio Code を使用して ARM テンプレートを作成する](https://learn.microsoft.com/ja-jp/azure/azure-resource-manager/templates/quickstart-create-templates-use-visual-studio-code?tabs=CLI)
* [Azure portal を使用して ARM テンプレートを作成およびデプロイする](https://learn.microsoft.com/ja-jp/azure/azure-resource-manager/templates/quickstart-create-templates-use-the-portal)
* 複数のリソースをひとつのARMテンプレート内に作成できる。`resources` プロパティ内に配列を指定する。
    * [Create multiple resource instances - Azure Resource Manager | Microsoft Learn](https://learn.microsoft.com/en-us/azure/azure-resource-manager/templates/template-tutorial-create-multiple-instances?tabs=CLI%2Cazure-cli)
* スナップショットからCopyしてディスクを作成といった指示もできる
  ```json
      {  
      "type": "Microsoft.Compute/disks",  
      "apiVersion": "2022-03-02",  
      "name": "[variables('managedDiskName')]",  
      "location": "[resourceGroup().location]",  
      "properties": {  
        "creationData": {  
          "createOption": "Copy",  
          "sourceResourceId": "[parameters('snapshotResourceId')]"  
        }  
      }  
    },  
  ```


### 実行方法

Cloud Shell から使う場合、テンプレートをどこに置くか？

> リモートに格納されている ARM テンプレートをデプロイするか、Cloud Shell のローカル ストレージ アカウントに格納されている ARM テンプレートをデプロイできます。

リモートに配置したら `--template-uri` で対象ファイルのURIを渡す。

```bash
az deployment group create \
  --name ExampleDeployment \
  --resource-group ExampleGroup \
  --template-uri "https://raw.githubusercontent.com/Azure/azure-quickstart-templates/master/quickstarts/microsoft.storage/storage-account-create/azuredeploy.json" \
  --parameters storageAccountType=Standard_GRS
  ```

* github.com に配置するなら公開可能な情報のみをARMに含めるべき
* ストレージにアップロードして SAS URL を発行する方法もある
* あるいは普通に CloudShell にアップロードしてローカルファイルを指定して実行する

```bash
az group create --name ExampleGroup --location "South Central US"
az deployment group create --resource-group ExampleGroup --template-file azuredeploy.json --parameters storageAccountType=Standard_GRS
```


参考）[Cloud Shell を使用したテンプレートのデプロイ - Azure Resource Manager | Microsoft Learn](https://learn.microsoft.com/ja-jp/azure/azure-resource-manager/templates/deploy-cloud-shell?tabs=azure-cli)
