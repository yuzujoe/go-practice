package main

import "fmt"

const Pi = 3.14

const (
	Username = "test_user"
	Password = "test_pass"
)

// var big int = 9898989898989898999 + 1
// constはコンパイル時に型を判別して実行してくれるのでこのような形で実行できる
const big = 989898989898989899 + 1

func main() {
	fmt.Println(Pi, Username, Password)
	fmt.Println(big - 1)
}
