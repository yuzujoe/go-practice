package main

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("world foo")

	fmt.Println("hello foo")
}

func main() {
	foo()

	defer fmt.Println("world")

	fmt.Println("hello")

	fmt.Println("run")
	// 最初に指定したdeferが最後に読み込まれる
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("success")

	// これでファイルの中身のコードを読み取ってくれる
	file, _ := os.Open("./defer.go")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}
