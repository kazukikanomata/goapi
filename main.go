package main

import (
	"fmt"

	"local.package/hello"
)

func main(){
	message := hello.Hello("ポカホンタス")
	fmt.Println(message)
}