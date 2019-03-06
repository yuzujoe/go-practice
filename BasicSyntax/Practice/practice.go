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

func searchmin() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

	var min int
	for i, num := range l {
		if i == 0 {
			min = num
			continue
		} else if min >= num {
			min = num
		}
		fmt.Println(min)
	}
}

func addFruits() {
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}

	sum := 0
	for _, v := range m {
		sum += v
	}

	fmt.Println(sum)
}

func main() {
	convert()
	maping()
	searchmin()
	addFruits()
}
