module example.com/myapp

go 1.21.1

replace example.com/calc => ../calc

require example.com/calc v0.0.0-00010101000000-000000000000

require github.com/google/go-cmp v0.5.9 // indirect
