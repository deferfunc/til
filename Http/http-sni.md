ロードバランサーがどの証明書をクライアントに対して使うのか？の仕組みについて

参考資料: [AFD + AppGW の構成で考える アクセス先 と SNI と ホストヘッダー の関係](https://zenn.dev/microsoft/articles/qiita-20230414-f55693d62af85e6f9cbd)

## まとめ

* SNI は TLS の時点で使用されるもの
* SNI は Https 接続で使用されるもの
* Http リクエストの Host ヘッダーとは役割が異なる
* クライアントから SNI で対象のホスト名が送信されるので、ロードバランサーは使用すべき証明書を判断できる(ロードバランサーに複数の証明書が設定されている場合)

# copilot とのやり取り

## SNIについて詳しく教えてください

SNI（Server Name Indication）は、TLS（Transport Layer Security）プロトコルの拡張機能であり、クライアントが接続を開始する際に、どのホスト名（ドメイン名）にアクセスしようとしているかをサーバーに伝えるために使用されます。これにより、同一のIPアドレスとポート番号を共有する複数のドメインが、それぞれ異なるSSL/TLS証明書を使用してセキュアな通信を行うことが可能になります。

### SNIの動作プロセス

1. **クライアントのリクエスト**: クライアント（例えばWebブラウザ）がサーバーにHTTPS接続を試みる際、TLSハンドシェイクの一環として、アクセスしようとしているサーバーのホスト名をSNIフィールドに含めて送信します。
2. **サーバーの応答**: サーバーは受け取ったSNI情報を基に、適切なSSL/TLS証明書を選択し、その証明書を使用してクライアントとのセキュアな通信を確立します。
3. **通信の確立**: 適切な証明書が選択され、クライアントがその証明書を検証することに成功すれば、クライアントとサーバー間で暗号化された通信が確立されます。

### SNIの利点

- **複数ドメインのサポート**: 単一のサーバーまたはロードバランサーで複数のドメインをホストし、それぞれに異なるSSL/TLS証明書を割り当てることができます。
- **コスト削減**: 複数のIPアドレスを購入する必要がなくなり、ホスティングコストを削減できます。
- **柔軟性の向上**: 新しいドメインを追加する際の設定が容易になり、運用の柔軟性が向上します。

### 注意点

- **古いクライアントの非対応**: SNIは比較的新しい技術であるため、古いブラウザやクライアントではサポートされていない場合があります。これらのクライアントからのアクセスを受け入れる必要がある場合は、対応策を検討する必要があります。
- **プライバシーの懸念**: SNIは暗号化されていないため、ネットワーク上でホスト名が露出する可能性があります。これにより、ユーザーのプライバシーに関する懸念が生じることがあります。TLS 1.3では、Encrypted SNI（ESNI）という技術が導入されており、この問題に対処しています。

SNIは、現代のインターネット通信において重要な役割を果たしており、複数のドメインを効率的に運用するための鍵となっています。

## SNIとHTTPリクエストのHostヘッダーの違いを教えて

SNI（Server Name Indication）とHTTPリクエストのHostヘッダーは、どちらもクライアントがサーバーに対して特定のリソースを要求する際に使用される情報ですが、彼らが機能するレイヤー（プロトコルの階層）と目的に違いがあります。

### SNI（Server Name Indication）

- **レイヤー**: SNIはTLS（Transport Layer Security）プロトコルの一部であり、SSL/TLSハンドシェイクの初期段階で使用されます。
- **目的**: SNIの主な目的は、同一のIPアドレスとポート番号を共有するサーバー上で複数のドメインをホストする際に、クライアントが接続しようとしている特定のドメイン名をサーバーに伝えることです。これにより、サーバーは適切なSSL/TLS証明書を選択してクライアントに提示できます。
- **使用例**: HTTPS接続を確立する際に、クライアントがサーバーに対してどのドメインの証明書を使用するかを指定するために使用されます。

### HTTPリクエストのHostヘッダー

- **レイヤー**: HostヘッダーはHTTPプロトコルの一部であり、HTTPリクエストのヘッダー情報として送信されます。
- **目的**: Hostヘッダーの目的は、クライアントが要求している特定のドメインまたはホスト名をサーバーに伝えることです。これにより、同一のIPアドレスで複数のドメインやウェブサイトをホストしているサーバーが、どのウェブサイトのリソースをクライアントに提供すべきかを判断できます。
- **使用例**: HTTPリクエストを送信する際に、クライアントがアクセスしようとしているウェブサイトのドメイン名を指定するために使用されます。

### 違いの要点

- **プロトコルのレイヤー**: SNIはTLSレベルで動作し、SSL/TLS証明書の選択に使用されます。一方、HostヘッダーはHTTPレベルで動作し、サーバー上の正しいリソースを特定するために使用されます。
- **通信の段階**: SNIは接続の確立段階（SSL/TLSハンドシェイク）で使用され、Hostヘッダーは接続が確立された後のHTTPリクエストで使用されます。
- **目的の違い**: SNIは主に証明書の選択に関連し、Hostヘッダーはリソースの特定に関連します。

これらの違いにより、SNIとHostヘッダーは、セキュアなウェブ通信を提供するために互いに補完的な役割を果たします。