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

// // 外部のパーケージから呼び出すときにNEWを使う時が多い
// func New(x, y int) *Vertex {
// 	return &Vertex{x, y}
// }

// ~Embedded~
type Vertex3D struct {
	Vertex
	z int
}

func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z
}

func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

// ~Embedded~

// ~non-struct~

type MyInt int

func (i MyInt) Double() int {
	return int(i * 2)
}

func main() {
	// v := Vertex{3, 4}
	// fmt.Println(Area(v))
	v := New(3, 4, 5)
	v.Scale3D(10)
	fmt.Println(v.Area3D())

	myInt := MyInt(10)
	fmt.Println(myInt.Double())
}
