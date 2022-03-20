package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// AnswerPage is answer page info struct.
type AnswerPage struct {
	LinkInfo template.HTML
	String   string
}

// Writing answer to HTML-style string and sending it on server.
func buildAnswerPage(rw http.ResponseWriter, r *http.Request) {
	var answerPage *template.Template = template.Must(template.ParseFiles("templates/answer.html"))
	var pageInfo *AnswerPage = new(AnswerPage)
	pageInfo.String = "<ul>\n"
	log.Println(getUsedLinks(r))

	for _, user := range getUsedLinks(r) {
		if user.IsAvailable {
			pageInfo.String += fmt.Sprintf("\t<li>\n\t<a name=\"%s\" href=\"%s\">%s: %s</a>\n\t</li>\t\n", user.SocialNetwork, user.Link, user.SocialNetwork, user.Name)
		} else {
			pageInfo.String += fmt.Sprintf("\t<li>\n\t<a name=\"%s\">%s: %s</a>\n\t</li>\t\n", user.SocialNetwork, user.SocialNetwork, user.Link)
		}
	}

	pageInfo.String += "</ul>\n"
	log.Println(pageInfo)
	pageInfo.LinkInfo = template.HTML(pageInfo.String)
	answerPage.Execute(rw, pageInfo)
	log.Println(answerPage)
}
