module local.package/main

go 1.21.1

replace local.package/greetings => ./greetings // ローカルパスへ置き換えるためモジュール名から相対パスを replace で渡す。

require (
	local.package/abc v0.0.0-00010101000000-000000000000
	local.package/hello v0.0.0-00010101000000-000000000000
)

require (
	local.package/disney v0.0.0-00010101000000-000000000000 // indirect
	local.package/greetings v0.0.0-00010101000000-000000000000 // indirect
)

replace local.package/hello => ./hello

replace local.package/abc => ./abc

replace local.package/disney => ./disney
