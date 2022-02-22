
## Protocol buffer

データシリアライズ機構。この機構は以下を提供する。

- ソースコード中のデータ構造を、バイナリデータにエンコードし、ファイルに書き出すことができる。書き出されたファイルを読み込み、デコードし、データ構造にバインドすることができる。
- エンコード、デコードできるデータは、.protoファイルによって構造化されている。.protoファイルは可読性が高い。
- .protoファイルに従って、言語毎にパーサーを自動生成できる。自動生成ツールをProtocol Buffer Compilerと呼ぶ。

```bash
brew install protobuf
protoc --version

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=internal --go_opt="module=github.com/suzuito/cel-sandbox-go" ./permission.proto ./auth.proto
```

某G社が、「式」を「評価」する機構を提供するライブラリを作っていました
https://github.com/google/cel-go
FAのキーワード検索に自前のクエリ言語作ったり
とか
自前のJWTのパーミションチェック
とかで使えそうかなと思いましたが
ライブラリの使い勝手はまだあんまり良くないという印象。
 Google社内ではgRPC向けに使われれているためか（？）、式中の「変数」にバインディングできるデータのデータ構造が、Protocol bufferのみという点が使いにくさ。 