package main

import "fmt"

func main() {
	t, f := true, false
	fmt.Printf("%T %v %t\n", t, t, 1)
	fmt.Printf("%T %v %t\n", f, f, 0)

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)

	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)

	// 反対の真偽型を呼び出す
	fmt.Println(!true)
	fmt.Println(!false)
}
