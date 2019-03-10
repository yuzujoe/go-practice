package mylib

import (
	"fmt"
	"regexp"
)

func Match() {
	match, _ := regexp.MatchString("a([a-z])+e", "apple")
	fmt.Println(match)

	// 再利用したい場合はこちらの方を汎用的に使う
	r := regexp.MustCompile("a([a-z])+e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fs := r2.FindString("/view/test")
	fmt.Println(fs)
	fss := r2.FindStringSubmatch("/view/test")
	fmt.Println(fss)
	fss2 := r2.FindStringSubmatch("/view/edit")
	fmt.Println(fss, fss2[0], fss[1], fss[2])
}
