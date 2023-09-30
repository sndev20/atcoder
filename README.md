# atcoder

### 参考ページ

atcoder-cli: http://tatamo.81.la/blog/2018/12/07/atcoder-cli-tutorial

### 利用するコマンド群

```shell
./script/new.sh <コンテストID>
./script/add.sh <コンテストID>
./script/test.sh <コンテストID> <問題ID>
./script/submit.sh <コンテストID> <問題ID>
```

### テンプレートの再利用

汎用性の高い関数（処理）は`templates/main.go`ファイルにまとめる。