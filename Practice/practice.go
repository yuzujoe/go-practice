package main

import "fmt"

func convert() {
	var f float64
	f = 1.11
	ff := int(f)
	fmt.Printf("%T %v %d\n", ff, ff, ff)

}

func maping() {
	m := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
	fmt.Printf("%T %v\n", m, m)
}

func main() {
	convert()
	maping()
}
