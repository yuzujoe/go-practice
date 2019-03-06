package main

import "fmt"

// this type declaration Vertex
type Vertex struct {
	X, Y int
}

func (v Vertex) Area() int {
	return v.X * v.Y
}

// *をつけることでVertexに紐づいたScaleというメソッドを呼び出すことができる
func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

// looking for coordinates with this function
func Area(v Vertex) int {
	return v.X * v.Y
}

func main() {
	v := Vertex{3, 4}
	// fmt.Println(Area(v))
	v.Scale(10)
	fmt.Println(v.Area())
}
