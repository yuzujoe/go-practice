package main

import "fmt"

// var変数は関数外からでも呼び出せる
var (
	// このように()でも実行可能
	i   int     = 1
	f64 float64 = 1.2
	s   string  = "test"
	t   bool    = true
	f   bool    = false
)

// xでの変数は関数内でしか呼び出すことができない
func foo() {
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	// Printfを使用する事で中身を見ることができる、表示の際に\nを入れる
	fmt.Printf("%T\n", xf64)
	fmt.Printf("%T\n", xi)
}

func main() {
	// var i int = 1
	// var f64 float64 = 1.2
	// var s string = "test"
	// var t bool = true
	// var f bool = false
	// var t, f bool = true, falseでも可能

	fmt.Println(i, f64, s, t, f)
	foo()
}
