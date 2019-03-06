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

// ~non-struct~

// interface

// 指定したメソッドを必ず使うときにinterfaceを使う
type Human interface {
	Say() string
}

type Person struct {
	Name string
}

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get out")
	}
}

// interface

// switch type

// interfaceはどんな型の引数でも受け付けてくれる
func do(i interface{}) {
	// ii := i.(int)
	// ii *= 2
	// fmt.Println(ii)

	// ss := i.(string)
	// fmt.Println(ss + "!")

	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know %T\n", v)
	}
}

// switch type

// Stringer

type Person2 struct {
	Name string
	Age  int
}

func (p Person2) String() string {
	return fmt.Sprintf("My name is %v.", p.Name)
}

// Stringer

func main() {
	// v := Vertex{3, 4}
	// fmt.Println(Area(v))
	v := New(3, 4, 5)
	v.Scale3D(10)
	fmt.Println(v.Area3D())

	myInt := MyInt(10)
	fmt.Println(myInt.Double())

	var mike Human = &Person{"Mike"}
	var x Human = &Person{"x"}
	DriveCar(mike)
	DriveCar(x)

	do(10)
	do("mike")
	do(true)

	ken := Person2{"Ken", 22}
	fmt.Println(ken)
}
