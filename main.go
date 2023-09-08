package main

import (
	"fmt"
	"os"

	"local.package/abc"
	"local.package/disney"
	"local.package/hello"
)
func getArgument() {
	fmt.Println(os.Args)
}

func main() {
	message := hello.Hello("ポカホンタス")
	fmt.Println(message)
	abc.Hello()
	disney.Disney()
	getArgument()
	fmt.Println(os.Args)
}