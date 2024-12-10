## 経緯

2024/12/8 頃、Azure Monitor Agent から Log Analytics Workspace に VM Insights のデータが到達していないことがわかった。

1. Azure Monitor のパフォーマンスグラフに何も表示されない
2. Log Analytics Workspace の InsightsMetrics テーブルにデータが無い

ことがわかり、AMAについてもろもろ調査した。

前提:

* VM 拡張機能 AzureMonitorLinuxAgent のバージョン 1.33.1
* VM のOS ubuntu 24.04 LTS
    * [Azure Monitor Agent supported operating systems - Azure Monitor | Microsoft Learn](https://learn.microsoft.com/en-us/azure/azure-monitor/agents/azure-monitor-agent-supported-operating-systems) ではこのバージョンはサポートされている

調査内容:

以下ページにそってトラブルシューティングを実施

[Linux オペレーティング システム (OS) の Azure Monitor エージェント トラブルシューティング ツールの使用方法 - Azure Monitor | Microsoft Learn](https://learn.microsoft.com/ja-jp/azure/azure-monitor/agents/troubleshooter-ama-linux?tabs=redhat%2CInteractive#run-the-troubleshooter)

root の状態で

```bash
cd /var/lib/waagent/Microsoft.Azure.Monitor.AzureMonitorLinuxAgent-1.33.1/ama_tst/
sudo sh ama_troubleshooter.sh -A
```

したら以下出力で、OSバージョンがサポートされていない表示 :sob: 困った


```
Python version being used is:
Python 3.12.3

Starting AMA Troubleshooting Tool v.1.3...

CHECKING INSTALLATION...
Checking if running a supported OS version...
ERROR(S) FOUND.
================================================================================
================================================================================
ALL ERRORS/WARNINGS ENCOUNTERED:
  ERROR FOUND: This version of ubuntu (24.04) is not supported. Please download 16.04, 18.04, 20.04 or 22.04. To see all supported Operating Systems, please go to:

   https://docs.microsoft.com/en-us/azure/azure-monitor/agents/agents-overview#linux

--------------------------------------------------------------------------------
Please review the errors found above.
================================================================================
If you still have an issue, please run the troubleshooter again and collect the logs for AMA.
In addition, please include the following information:
  - Azure Subscription ID where the Log Analytics Workspace is located
  - Workspace ID the agent has been onboarded to
  - Workspace Name
  - Region Workspace is located
  - Pricing Tier assigned to the Workspace
  - Linux Distribution on the VM
  - Azure Monitor Agent Version
================================================================================
Restarting AMA can solve some of the problems. If you need to restart Azure Monitor Agent on this machine, please execute the following commands as the root user:
  $ cd /var/lib/waagent/Microsoft.Azure.Monitor.AzureMonitorLinuxAgent-<agent version number>/
  $ ./shim.sh -disable
  $ ./shim.sh -enable
```

### AMA について

VMInsightsを有効化した際の収集されたデータの種類と行き先

Azure MonitorのVM Insightsは、仮想マシン（VM）を監視する際にさまざまなデータを収集します。収集されるデータの種類には以下が含まれます：

* パフォーマンスデータ: CPU、メモリ、ディスク、およびネットワークパフォーマンスのメトリックが `Perf` テーブルに格納されます。
* InsightsMetrics: VM Insightsは一般的なパフォーマンスカウンターを収集し、それらをLog Analyticsワークスペースの `InsightsMetrics` テーブルに送信します。
* VMProcess: 仮想マシン上で実行されているプロセスに関する情報。
* VMConnection: 仮想マシンへの入出力トラフィック接続に関するデータ。
* VMComputer: 仮想マシンのプロパティなどのインベントリデータ。
* VMBoundPort: 仮想マシン上のオープンサーバーポートに対するトラフィックの詳細。
