package net

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	// omitemptyで指定したものが空の場合隠すことができる
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	Nicknames []string `json:"nicknames"`
}

// Unmarshalが呼ばれた際のカスタム
func (p *Person) UnmarshalJSON(b []byte) error {
	type Person2 struct {
		Name string
	}
	var p2 Person2
	err := json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Println(err)
	}
	p.Name = p2.Name + "!"
	return err
}

// Marshalが呼ばれた際に中身をカスタムする
func (p Person) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr." + "p.Name",
	})
	return v, err
}

func Json() {
	b := []byte(`{"name":"mike","age":20,"nicknames":["a","b","c"]}`)
	var p Person
	// Unmarshalはstructのキーを見てそのまま返してくれる
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)
	// Marshalすると好き名前で指定できる
	v, _ := json.Marshal(p)
	fmt.Println(string(v))
}
