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
