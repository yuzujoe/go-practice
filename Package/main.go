package main

import (
	// 基本的に氷刃ライブラリから読み込んでいく
	"fmt"
	// 階層フォルダにファイルを作成してmainでimportすると使える
	"./mylib"
	"./mylib/under"
	"./talib"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))
	mylib.Say()
	under.Hello()

	talib.Talib()

	person := mylib.Person{Name: "Mike", Age: 20}
	fmt.Println(person)

	fmt.Println(mylib.Public)
}
