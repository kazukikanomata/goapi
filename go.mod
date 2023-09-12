module local.package/main

go 1.21.1

replace local.package/greetings => ./greetings // ローカルパスへ置き換えるためモジュール名から相対パスを replace で渡す。

replace local.package/hello => ./hello

replace local.package/abc => ./abc

replace local.package/disney => ./disney
