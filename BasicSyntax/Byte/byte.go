package main

import "fmt"

func main() {
	b := []byte{72, 73}
	fmt.Println(b)
	// stringでキャストすることによって期待された結果が返ってくる。
	fmt.Println(string(b))

	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))
}
