package main

import "fmt"

func main() {
	l := []string{"python", "go", "java"}

	for i, v := range l {
		fmt.Println(i, v)
	}
	// これでindex番号を使わなくてよくなる
	for _, v := range l {
		fmt.Println(v)
	}

	m := map[string]int{"apple": 100, "banana": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}

	for k := range m {
		fmt.Println(k)
	}

	for _, v := range m {
		fmt.Println(v)
	}
}
