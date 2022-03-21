package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// HTMLInfo contains all page info.
type HTMLInfo struct {
	CheckBoxInfo       template.HTML
	CheckBoxInfoString string
	LinkInfo           template.HTML
	LnkInfoString      string
}

// Builds page.
func page(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		addAnswers(rw, r)
		return
	}

	var formPage *template.Template = template.Must(template.ParseFiles("templates/page.html"))
	pageInfo := addCheckBoxes()
	formPage.Execute(rw, pageInfo)
}

// Adds check boxes on page.
func addCheckBoxes() *HTMLInfo {
	var pageInfo *HTMLInfo = new(HTMLInfo)
	pageInfo.CheckBoxInfoString = "<ul>\n"

	for _, page := range Pages {
		pageInfo.CheckBoxInfoString += fmt.Sprintf(`
            <li>
                <label for="%s">%s</label>
                <input type="checkbox" id="%s" name="%s"/>
            </li>
		`, page.ID, page.Name, page.ID, page.ID)
	}

	pageInfo.CheckBoxInfoString += "</ul>"
	log.Println(pageInfo)
	pageInfo.CheckBoxInfo = template.HTML(pageInfo.CheckBoxInfoString)
	return pageInfo
}

// Checking which textboxes are set on and creating container of user info then getting answer.
func getUsedLinks(r *http.Request) []*UserInfo {
	r.ParseForm()
	nick := r.FormValue("nickname")
	log.Println(r.Form)

	container := NewRequesterContainer(nick)
	for key, _ := range r.Form {
		if _, ok := container.Requesters[key]; ok {
			container.Requesters[key] = &RequesterAvailability{
				container.Requesters[key].requester,
				true,
			}
		}
	}

	return container.GetLinks()
}

// Adds answers to page.
func addAnswers(rw http.ResponseWriter, r *http.Request) {
	var answerPage *template.Template = template.Must(template.ParseFiles("templates/page.html"))
	pageInfo := addCheckBoxes()
	pageInfo.LnkInfoString = "<h3>Results:</h3>\n\t\t<ul>\n"
	log.Println(getUsedLinks(r))

	for _, user := range getUsedLinks(r) {
		if user.IsAvailable {
			pageInfo.LnkInfoString += fmt.Sprintf("\t\t\t<li>\n\t\t\t\t<a name=\"%s\" href=\"%s\">%s: %s</a>\n\t\t\t</li>\t\n", user.SocialNetwork, user.Link, user.SocialNetwork, user.Name)
		} else {
			pageInfo.LnkInfoString += fmt.Sprintf("\t\t\t<li>\n\t\t\t\t<a name=\"%s\">%s: %s</a>\n\t\t\t</li>\t\n", user.SocialNetwork, user.SocialNetwork, user.Link)
		}
	}

	pageInfo.LnkInfoString += "\t\t</ul>"
	log.Println(pageInfo)
	pageInfo.LinkInfo = template.HTML(pageInfo.LnkInfoString)
	answerPage.Execute(rw, pageInfo)
	log.Println(answerPage)
}
