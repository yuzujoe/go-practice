package mylib

import (
	"fmt"
	"io/ioutil"
	"log"
)

func Ioutil() {
	content, err := ioutil.ReadFile("main.go")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(content))
	// WriteFileは、filenameで指定したファイルへdataを書き込みます。ファイルが存在しないときは、そのファイルをパーミッションpermで作成します。ファイルが存在したときは、書き込み前にトランケートを行います。
	if err2 := ioutil.WriteFile("ioutil_tenp.go", content, 0666); err != nil {
		log.Fatalln(err2)
	}

	// r := bytes.NewBuffer([]byte("abc"))
	// ReadAllは、エラーまたはEOFに達するまでrから読み込み、読み込んだデータを返します。
	// content, _ := ioutil.ReadAll(r)
	// fmt.Println(string(content))
}
