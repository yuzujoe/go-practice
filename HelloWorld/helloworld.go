package main

import "fmt"

// init関数はmain関数よりも先に呼び出される特別な関数
func init() {
	fmt.Println("Init")
}

// 別の関数を呼び出したい場合はfunc関数を作ってからmain関数の中で呼び出しを行う
func bazz() {
	fmt.Println("Bazz")
}

// golangは基本的にmain関数の中を読み込みに行く
func main() {
	bazz()
	// カンマ区切りで次の文字列も呼び出せる
	fmt.Println("Hello World", "TEST TEST")
}
