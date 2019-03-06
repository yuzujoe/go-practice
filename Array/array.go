package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	// var b [2]int
	// b = [2]int{100, 200}
	// fmt.Println(b)

	// 配列自体はリサイズができないが[]のスライスであれば変更できる
	var b []int
	b = []int{100, 200}
	b = append(b, 300)
	fmt.Println(b)
}
