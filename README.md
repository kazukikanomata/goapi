# goapi

参考にしている記事
- パッケージを分割する
https://qiita.com/akifumii/items/1eaba6e11caa6b8f688d


##　パッケージを書く

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