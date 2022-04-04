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
	LinkInfoString     string
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

	for _, page := range Pages {
		pageInfo.CheckBoxInfoString += fmt.Sprintf(`
            <li>
                <label for="%s">%s</label>
                <input type="checkbox" name="%s" id="%s"/>
            </li>
		`, page.ID, page.Name, page.ID, page.ID)
	}
	log.Println(pageInfo)
	pageInfo.CheckBoxInfo = template.HTML(pageInfo.CheckBoxInfoString)
	return pageInfo
}

// Checking which textboxes are set on and creating container of user info then getting answer.
func getUsedLinks(r *http.Request, nickname string) []*UserInfo {
	container := NewRequesterContainer(nickname)
	for key, _ := range r.Form {
		if _, ok := container.Requesters[key]; ok {
			container.Requesters[key] = &RequesterAvailability{
				container.Requesters[key].requester,
				true,
			}
			fmt.Println("OK")
		}
	}

	return container.GetLinks()
}

// Adds answers to page.
func addAnswers(rw http.ResponseWriter, r *http.Request) {
	var answerPage *template.Template = template.Must(template.ParseFiles("templates/page.html"))
	pageInfo := addCheckBoxes()

	r.ParseForm()
	nick := r.FormValue("nickname")
	log.Println(r.Form)

	if nick == "" {
		pageInfo.LinkInfoString = "<h3>Looks like the nickname is invalid...</h3>\n\t\t<ul>\n"
		log.Println(pageInfo)
		pageInfo.LinkInfo = template.HTML(pageInfo.LinkInfoString)
		answerPage.Execute(rw, pageInfo)
		log.Println(answerPage)
		return
	}

	users := getUsedLinks(r, nick)
	log.Println(users)

	if len(users) == 0 {
		pageInfo.LinkInfoString = "<h3>Looks like you didn't select any pages...</h3>\n\t\t<ul>\n"
		log.Println(pageInfo)
		pageInfo.LinkInfo = template.HTML(pageInfo.LinkInfoString)
		answerPage.Execute(rw, pageInfo)
		log.Println(answerPage)
		return
	}

	pageInfo.LinkInfoString = fmt.Sprintf("<h3>Results for nickname \"%s\":</h3>\n\t\t<ul>\n", nick)

	for _, user := range users {
		if user.IsAvailable {
			pageInfo.LinkInfoString += fmt.Sprintf("\t\t\t<li>\n\t\t\t\t<a name=\"%s\" href=\"%s\">%s: %s</a>\n\t\t\t</li>\t\n", user.SocialNetwork, user.Link, user.SocialNetwork, user.Name)
		} else {
			pageInfo.LinkInfoString += fmt.Sprintf("\t\t\t<li>\n\t\t\t\t<a name=\"%s\">%s: %s</a>\n\t\t\t</li>\t\n", user.SocialNetwork, user.SocialNetwork, user.Link)
		}
	}

	pageInfo.LinkInfoString += "\t\t</ul>"
	log.Println(pageInfo)
	pageInfo.LinkInfo = template.HTML(pageInfo.LinkInfoString)
	answerPage.Execute(rw, pageInfo)
	log.Println(answerPage)
}
