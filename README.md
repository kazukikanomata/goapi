# Go Modules

Go言語の依存モジュール管理ツールのこと。
Go.1.11から標準で利用することができるようになった。

## Go.1.11以前はGOPATHの概念！

[GOPATHモード]
- $GOPATH配下で管理する必要があった。

[モジュールモード]
- Go Modulesの機能は、`go.sum`と`go.mod`の2つになる！

`$ go mod init`
を実行してカレントディレクトリにファイルが存在する時点でmoduleモードが発動する。

ライブラリ、パッケージ、モジュールの違いについては、下記のイメージが参考になるのかも
- https://osamu-tech-blog.com/wp-content/uploads/2022/05/image.png

## go.sumファイル
- 必要な理由は、`go.mod`ファイルをもとに取得したモジュールが本当に`go.sum`ファイル生成時のものと一致しているかのチェックのため

### パッケージ
- 単一のディレクトリを指す。

### モジュール
`go.mod`ファイルのあるディレクトリ以下の全てのパッケージがモジュール

参考にしている記事
- パッケージを分割する
- https://qiita.com/akifumii/items/1eaba6e11caa6b8f688d


### ファイルを分割して、パッケージをimportする手順

ファイル構成
```
api
│   ├── README.md
│   ├── calc
│   │   ├── calc.go
│   │   ├── calc_test.go
│   │   ├── go.mod
│   │   └── go.sum
│   └── myapp
│       ├── go.mod
│       ├── go.sum
│       └── main.go
```

やること
`main.go`で`/calc/calc.go`を取り入れる。

これを自作モジュールの利用という。（ローカル）

1. ディレクトリの作成

```
$ mkdir myapp
```

2. ディレクトリに入る

```
$ cd myapp/
```

3. モジュールファイルを作成

```
$ go mod init example.com/myapp
```

4.  ぶち込みたいファイルに記述


`main.go`

```
package main

import (
   "fmt"
   "example.com/calc"
)

func main() {
	fmt.Println(calc.Max(1, 2)) # ファイル名.関数名がポイント
}
```

5. ローカルマシン上のモジュールの場所を知らせる準備

以下のコマンドを使用し、モジュールの場所を知らせる。
確認する

`$ go mod edit -replace "example.com/calc=../calc"`

`$ cat go.mod`

6. 依存関係にファイルを追加する

`$ go mod tidy`

or

`$ go get example.com/calc`

7. ディレクトリを移動し、`main.go`を実行

`$ go run main.go`

引用：https://www.twihike.dev/docs/golang-primer/using-modules


### パッケージには次の種類があります。

- 標準パッケージ：Goが最初から用意しているもの（一覧は公式ドキュメントに公開）
- サードパーティパッケージ：第三者が用意したもの
- 自作パッケージ：自分で用意したもの

### go get
githubにある自分で作った別のリポジトリとか、野良リポジトリ使うときは、go get使って持ってきたら良さそう。バージョン指定するときは、@のあとにlatestとかタグとかコミットのハッシュ値をつける。

`$ go get github.com/hoge/fuga@latest``

### go mod tidy
これ打つと、go.sumってファイルができて、パッケージのインストールをいい感じにしてくれる。importしてるやつをとってきてくれてそう。不要なやつも取り除いてくれる。

## API

- GoのHTTPサーバー
> https://tutuz-tech.hatenablog.com/entry/2020/03/23/162831

- Goで作るWebAPIサーバー
> https://kyutechlc.blog.jp/archives/1079246765.html