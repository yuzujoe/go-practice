package main

import "fmt"

// this type declaration Vertex
type Vertex struct {
	x, y int
}

func (v Vertex) Area() int {
	return v.x * v.y
}

// *をつけることでVertexに紐づいたScaleというメソッドを呼び出すことができる
func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

// looking for coordinates with this function
// func Area(v Vertex) int {
// 	return v.X * v.Y
// }

// 外部のパーケージから呼び出すときにNEWを使う時が多い
func New(x, y int) *Vertex {
	return &Vertex{x, y}
}

func main() {
	// v := Vertex{3, 4}
	// fmt.Println(Area(v))
	v := New(3, 4)
	v.Scale(10)
	fmt.Println(v.Area())
}
