package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var ansPage string = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8"/>
    <title>My test page</title>
    <link rel="stylesheet" href="assets/page_style.css">
</head>

<body>
`

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
		return
	}

	buildAnswerPage(rw, r)
}

func buildAnswerPage(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(getUsedLinks(r))
	for _, user := range getUsedLinks(r) {
		if user.IsAvailable {
			ansPage += fmt.Sprintf("<a name=\"%s\" href=\"%s\">%s: %s</a>\n<br/>\n", user.SocialNetwork, user.Link, user.SocialNetwork, user.Name)
		} else {
			ansPage += fmt.Sprintf("<a name=\"%s\">%s: %s</a>\n<br/>\n", user.SocialNetwork, user.SocialNetwork, user.Link)
		}
	}
	ansPage += `</body>
</html>`
	rw.Write([]byte(ansPage))
}

func getUsedLinks(r *http.Request) []UserInfo {
	r.ParseForm()
	nick := r.FormValue("nickname")
	fmt.Println(r.Form)

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
