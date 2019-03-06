package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "softbank"
}

func main() {

	switch os := getOsName(); os {
	case "mac":
		fmt.Println("Mac!!")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Println("Default", os)
	}

	// timeというパッケージで時間を取得できる
	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")
	}
}
