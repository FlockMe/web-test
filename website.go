package main

import (
 "fmt"
 "log"
 "net/http"
 "io/ioutil"
)

// Save, loadPage function is copied from golang.com
// This code is view source for learning.
// ~ FlockMe Club

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

func home(w http.ResponseWriter, r *http.Request) {
 title := r.URL.Path[len("/"):]
 p, _ := loadPage(title)
 

 fmt.Fprintf(w, "<h1>Ini judul <b>%s</b></h1><br /><p>Ini badan %s</p>", p.Title, p.Body)
}

func main() {
 http.HandleFunc("/", home)
 log.Fatal(http.ListenAndServe(":8080", nil))
}
