package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Handler struct {
	Name string
}

func (hand *Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	page(rw, r)
}

func page(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		page, err := ioutil.ReadFile("page.html")
		if err != nil {
			panic(err)
		}
		rw.Write(page)
	} else {
		buildAnswerPage(rw, r)
	}
}

func buildAnswerPage(rw http.ResponseWriter, r *http.Request) {
	var ansPage string = `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8"/>
	<title>Answer</title>
	<link rel="stylesheet" href="assets/answer.css">
</head>
	
<body>
<div class="links">
`
	log.Println(getUsedLinks(r))
	for _, user := range getUsedLinks(r) {
		if user.IsAvailable {
			ansPage += fmt.Sprintf("\t<a name=\"%s\" href=\"%s\">%s: %s</a>\n\t<br/>\n", user.SocialNetwork, user.Link, user.SocialNetwork, user.Name)
		} else {
			ansPage += fmt.Sprintf("\t<a name=\"%s\">%s: %s</a>\n\t<br/>\n", user.SocialNetwork, user.SocialNetwork, user.Link)
		}
	}
	ansPage += `</div>
</body>
</html>`
	log.Println(ansPage)
	fmt.Fprintf(rw, ansPage)
}

func getUsedLinks(r *http.Request) []UserInfo {
	r.ParseForm()
	nick := r.FormValue("nickname")
	log.Println(r.Form)

	container := NewRequesterContainer(nick)
	for key, _ := range r.Form {
		if _, ok := container.Requesters[key]; ok {
			container.Requesters[key] = RequesterAvailability{
				container.Requesters[key].requester,
				true,
			}
		}
	}

	return container.GetLinks()
}
