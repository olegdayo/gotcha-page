package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type FormPage struct {
	CheckBoxInfo template.HTML
	String       string
}

// Building form if it is not POST request else answer.
func page(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		buildAnswerPage(rw, r)
		return
	}

	var formPage *template.Template = template.Must(template.ParseFiles("templates/form.html"))
	var pageInfo *FormPage = new(FormPage)
	pageInfo.String = "<ul>\n"

	for _, page := range Pages {
		pageInfo.String += fmt.Sprintf(`
            <li>
                <label for="%s">%s</label>
                <input type="checkbox" id="%s" name="%s"/>
            </li>
		`, page.ID, page.Name, page.ID, page.ID)
	}

	pageInfo.String += "</ul>\n"
	log.Println(pageInfo)
	pageInfo.CheckBoxInfo = template.HTML(pageInfo.String)
	formPage.Execute(rw, pageInfo)
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
