package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// Building form if it is not POST request else answer.
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

// Checking which textboxes are set on and creating container of user info then getting answer.
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

// AnswerPage is answer page info struct.
type AnswerPage struct {
	LinkInfo template.HTML
	String   string
}

// Writing answer to HTML-style string and sending it on server.
func buildAnswerPage(rw http.ResponseWriter, r *http.Request) {
	var ansPage *template.Template = template.Must(template.ParseFiles("templates/answerTemplate.html"))
	var pageInfo *AnswerPage = new(AnswerPage)

	log.Println(getUsedLinks(r))

	for _, user := range getUsedLinks(r) {
		if user.IsAvailable {
			pageInfo.String += fmt.Sprintf("\t<a name=\"%s\" href=\"%s\">%s: %s</a>\n\t<br/>\n", user.SocialNetwork, user.Link, user.SocialNetwork, user.Name)
		} else {
			pageInfo.String += fmt.Sprintf("\t<a name=\"%s\">%s: %s</a>\n\t<br/>\n", user.SocialNetwork, user.SocialNetwork, user.Link)
		}
	}

	log.Println(pageInfo)
	pageInfo.LinkInfo = template.HTML(pageInfo.String)
	ansPage.Execute(rw, pageInfo)
	log.Println(ansPage)
}
