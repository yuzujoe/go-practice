package main

import "fmt"

func one(x *int) {
	*x = 1
}

func main() {
	var n int = 1000
	one(&n)
	fmt.Println(n)

	fmt.Println(&n)

	// var p *int = &n
	// fmt.Println(p)
	// fmt.Println(*p)

	// newを入れることによってアドレスを取得できる
	// makeとの違いはポインタを返すか返さないかの違い
	var p *int = new(int)
	fmt.Println(*p)
	*p++
	fmt.Println(*p)

	// こちらはポインタの宣言はされているがメモリが確保されてないのでnilが帰ってくる
	var p2 *int
	fmt.Println(p2)
	*p++
	fmt.Println(p2)
}
