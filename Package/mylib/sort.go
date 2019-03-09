package mylib

import (
	"fmt"
	"sort"
)

func Sort() {
	i := []int{3, 5, 6, 4, 1}
	s := []string{"s", "d", "r"}
	p := []struct {
		Name string
		Age  int
	}{
		{"Nancy", 20},
		{"Vera", 40},
		{"Mike", 50},
		{"Tom", 30},
	}
	fmt.Println(i, s, p)
	sort.Ints(i)
	sort.Strings(s)
	sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })
	sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age })
	fmt.Println(i, s, p)
}
