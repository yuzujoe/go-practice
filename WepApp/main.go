package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func PageInit() {
	page1 := &Page{Title: "test", Body: []byte("this is a sample page.")}
	page1.save()

	page2, _ := loadPage(page1.Title)
	fmt.Println(string(page2.Body))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// /view/test
	title := r.URL.Path[len("/view/"):]
	page, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s<h1/><div>%s</div>", page.Title, page.Body)
}

func main() {
	// PageInit()
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
