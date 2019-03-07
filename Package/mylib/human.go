package mylib

import "fmt"

// Public privete

var Public string = "Public"

// 頭文字を小文字にするとprivateになって出力できない
type Person struct {
	Name string
	Age  int
}

// Public privete

func Say() {
	fmt.Println("Human!")
}
