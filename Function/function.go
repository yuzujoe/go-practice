package main

import "fmt"

func add(x int, y int) (int, int) {
	return x - y, x + y
}

// 外から指定してあげることでresutはintで使うことができるようになる
// ここで指定することで何をしているのわかりやすくなる
func cal(price, item int) (result int) {
	result = price * item
	// あらかじめresultを宣言している為returnで記載する必要がある
	return
}

func main() {
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 2)
	fmt.Println(r3)

	f := func(x int) {
		fmt.Println("inner funtと言う", x)
	}
	f(1)

	func(x int) {
		fmt.Println("inner funtと言う", x)
	}(2)
}
