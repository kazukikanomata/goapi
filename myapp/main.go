package main

import (
	"fmt"

	"example.com/calc"
)

type Person struct {
	name string
	age  int
}
type Score struct {
	UserID string
	GameID int
	point  int
}

// ... 要素数を値から推測している
var ns = [...]int{10, 20, 30, 40, 50}
var toritasi = ns[1:3]
var lang string = "Go"
var langP *string = &lang // ポインタ変数の宣言

func swap(x, y int) (int, int) {
	return y, x
}
func f(xp *int) {
	*xp = 100
}
func dev2(i int) string {
	if i%2 == 0 {
		return "-偶数"
	} else {
		return "-奇数"
	}
}

func main() {
	fmt.Println(calc.Max(1, 2))

	profile := Person{"Gopher", 10}
	fmt.Println(profile.name, profile.age)
	fmt.Println(toritasi)
	fmt.Println(&lang)
	fmt.Println(*langP)
	fmt.Println(ns)

	x, _ := swap(10, 20)
	fmt.Println(x)

	var xer int
	f(&xer)
	println(xer)

	for i := 1; i <= 100; i++ {
		print(i)
		println(dev2(i))
	}
}
