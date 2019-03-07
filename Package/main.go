package main

import (
	// 基本的に氷刃ライブラリから読み込んでいく
	"fmt"
	// 階層フォルダにファイルを作成してmainでimportすると使える
	"./mylib"
	"./mylib/under"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))
	mylib.Say()
	under.Hello()
}
