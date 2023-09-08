# goapi

## Go Modules

Go言語の依存モジュール管理ツールのこと。
Go.1.11から標準で利用することができるようになった。

### Go.1.11以前は
GOPATHの概念！

[GOPATHモード]
$GOPATH配下で管理する必要があった。

[モジュールモード]
Go Modulesの機能は、`go.sum`と`go.mod`の2つになる！

ライブラリ、パッケージ、モジュールの違いについては、下記のイメージが参考になるのかも

https://osamu-tech-blog.com/wp-content/uploads/2022/05/image.png

### go.sumファイル
必要な理由は、`go.mod`ファイルをもとに取得したモジュールが本当に`go.sum`ファイル生成時のものと一致しているかのチェックのため

### パッケージ
単一のディレクトリを指す。

### モジュール
`go.mod`ファイルのあるディレクトリ以下の全てのパッケージがモジュール



参考にしている記事
- パッケージを分割する
https://qiita.com/akifumii/items/1eaba6e11caa6b8f688d


### パッケージを書く

1. フォルダを作る

2. `go.mod`ファイルを作る
   - `$ go mod init local.package/フォルダ名`

3. フォルダに入ってファイルを作る
   - `$ cd `
   - `$ touch ファイル名`

4. コードを書く。（関数名は大文字で）

5. importさせたいファイルに書いていく。
   -  関数内で記述する。
   -  import文を書く
   
6. `$ go get local.package/フォルダ名`コマンドを打つ

7. ファイルを実行`$ go run .`


### 引っかかったところ

go.modのversionが異なる。

- ./ルートディレクトリは 1.21.1

- /hello と　/greetingsは　1.20

- /disney と /abcは 1.21.1

レンタルMacローカル環境はversion　1.20

### go get
githubにある自分で作った別のリポジトリとか、野良リポジトリ使うときは、go get使って持ってきたら良さそう。バージョン指定するときは、@のあとにlatestとかタグとかコミットのハッシュ値をつける。

`$ go get github.com/hoge/fuga@latest``

### go mod tidy
これ打つと、go.sumってファイルができて、パッケージのインストールをいい感じにしてくれる。importしてるやつをとってきてくれてそう。不要なやつも取り除いてくれる。

