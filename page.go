package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("page.html"))

type Handler struct {
	Name string
}

func (hand *Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	page(rw, r)
}

func page(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		//page, err := ioutil.ReadFile("page.html")
		err := tpl.Execute(rw, nil)
		if err != nil {
			panic(err)
		}
		//rw.Write(page)
		return
	}

	r.ParseForm()
	nick := r.FormValue("nickname")
	fmt.Println(r.Form)
	fmt.Fprintf(rw, fmt.Sprintf("Searching nickname: %s\n", nick))
}
