package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello + World")
	fmt.Println(string("Hello World"[0]))

	var s string = "Hello World"
	// strings.Replaceで文字列の置き換えができる。この場合は変数sのHをXに置き換える１つだけの意味になる
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)
	// Containsで含んでいるかどうかを判断する
	fmt.Println(strings.Contains(s, "World"))
	// ``で囲むと囲った中で改行してくれる
	fmt.Println(`Test
						Test
Test`)
	// 特殊文字を囲いたい場合も``で対応できるもしくは/で対応
	fmt.Println(`"`)
}
