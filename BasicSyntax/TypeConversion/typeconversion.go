package main

import (
	"fmt"
	"strconv"
)

func main() {
	// intからfloat64への型変換
	var x int
	x = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	// float64からintへの型変換
	var y float64
	y = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	var s string
	s = "14"
	// _にはエラーが入っているのだがこれを定義するとエラーが出てしまうので_で対処
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v\n", i, i)

	h := "Hello World"
	fmt.Println(string(h[0]))
}
