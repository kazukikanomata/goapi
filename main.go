package main

import (
	"fmt"

	"local.package/abc"
	"local.package/disney"
	"local.package/hello"
)

func main() {
	message := hello.Hello("ポカホンタス")
	fmt.Println(message)
	abc.Hello()
	disney.Disney()
}
