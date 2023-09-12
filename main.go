package main

import (
	"fmt"
	"net/http"
	"os"
)

func getArgument() {
	fmt.Println(os.Args)
}

func main() {
	//  エンドポイントを作成する
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf: 書き込み先を指定する。ResponseWriter
		fmt.Fprintf(w, "Hello World!")
	})
	fmt.Println("サーバー起動！")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
