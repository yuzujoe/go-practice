package main

import "fmt"

func one(x *int) {
	*x = 1
}

// Vertex is export function
type Vertex struct {
	// この中身を小文字で表記してしまうとprivateになるのでアクセスできない
	X int
	Y int
	S string
}

func changeVertex(ver Vertex) {
	ver.X = 1000
}

func changeVertex2(ver *Vertex) {
	ver.X = 1000
}

func main() {

	ver := Vertex{1, 2, "test"}
	changeVertex(ver)
	fmt.Println(ver)

	ver2 := Vertex{1, 2, "test2"}
	changeVertex2(&ver2)
	fmt.Println(ver2)

	var n int
	n = 100
	one(&n)
	fmt.Println(n)

	fmt.Println(&n)

	// var p *int = &n
	// fmt.Println(p)
	// fmt.Println(*p)

	// newを入れることによってアドレスを取得できる
	// makeとの違いはポインタを返すか返さないかの違い
	var p *int
	p = new(int)
	fmt.Println(*p)
	*p++
	fmt.Println(*p)

	// こちらはポインタの宣言はされているがメモリが確保されてないのでnilが帰ってくる
	var p2 *int
	fmt.Println(p2)
	*p++
	fmt.Println(p2)

	// v := Vertex{X: 1, Y: 2}
	// fmt.Println(v)
	// fmt.Println(v.X, v.Y)

	// v.X = 100
	// fmt.Println(v.X, v.Y)

	// // {1 0 }この結果が返ってくる
	// v2 := Vertex{X: 1}
	// fmt.Println(v2)

	// // structの順番で入れていっても結果が反映される
	// v3 := Vertex{1, 2, "test"}
	// fmt.Println(v3)

	// v4 := Vertex{}
	// fmt.Println(v4)

	// // {0 0 }が返ってくる
	// var v5 Vertex
	// fmt.Println("%T %v", v5, v5)

	// v6 := new(Vertex)
	// fmt.Println("%T %v", v6, v6)

	// // こちらの方が使用頻度が高い
	// v7 := &Vertex{}
	// fmt.Println("%T %v", v7, v7)
}
