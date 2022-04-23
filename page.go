package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

// HTMLInfo contains all page info.
type HTMLInfo struct {
	NicknameInfo       template.HTML
	NicknameInfoString strings.Builder
	CheckBoxInfo       template.HTML
	CheckBoxInfoString strings.Builder
	LinkInfo           template.HTML
	LinkInfoString     strings.Builder
}

// Builds page.
func page(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		addAnswers(rw, r)
		return
	}

	var formPage *template.Template = template.Must(template.ParseFiles("templates/page.html"))
	pageInfo := addCheckBoxesAndNickname(NewRequesterContainer(""), "")
	formPage.Execute(rw, pageInfo)
}

// Adds check boxes and nickname value in text input field on page.
func addCheckBoxesAndNickname(container *RequesterContainer, nickname string) (pageInfo *HTMLInfo) {
	pageInfo = new(HTMLInfo)
	var isChecked string
	for _, page := range Pages {
		isChecked = ""
		if container.Requesters[page.ID].IsSelected() {
			isChecked = "checked"
		}
		pageInfo.CheckBoxInfoString.WriteString(fmt.Sprintf(`
            <li>
                <label for="%s">%s</label>
                <input type="checkbox" name="%s" id="%s"%s/>
            </li>
		`, page.ID, page.Name, page.ID, page.ID, isChecked))
	}
	log.Println(pageInfo)
	pageInfo.CheckBoxInfo = template.HTML(pageInfo.CheckBoxInfoString.String())
	pageInfo.NicknameInfo = template.HTML(nickname)
	return pageInfo
}

// Checking which textboxes are set.
func setUsedLinks(r *http.Request, container *RequesterContainer) {
	for key, _ := range r.Form {
		if _, ok := container.Requesters[key]; ok {
			container.Requesters[key].SetAvailability(true)
			fmt.Println("OK")
		}
	}
}

// Adds answers to page.
func addAnswers(rw http.ResponseWriter, r *http.Request) {
	var answerPage *template.Template = template.Must(template.ParseFiles("templates/page.html"))

	r.ParseForm()
	nick := r.FormValue("nickname")
	log.Println(r.Form)

	// Container initialization and execution.
	container := NewRequesterContainer(nick)
	setUsedLinks(r, container)
	users := container.GetLinks()
	log.Println(users)

	// Setting page and getting all the info about it.
	pageInfo := addCheckBoxesAndNickname(container, nick)

	// Checking if nickname is ok.
	if nick == "" {
		pageInfo.LinkInfoString.WriteString("<h3>Looks like the nickname is invalid...</h3>\n\t\t<ul>\n")
		log.Println(pageInfo)
		pageInfo.LinkInfo = template.HTML(pageInfo.LinkInfoString.String())
		answerPage.Execute(rw, pageInfo)
		log.Println(answerPage)
		return
	}

	// Checking if any sites are set.
	if len(users) == 0 {
		pageInfo.LinkInfoString.WriteString("<h3>Looks like you didn't select any pages...</h3>\n\t\t<ul>\n")
		log.Println(pageInfo)
		pageInfo.LinkInfo = template.HTML(pageInfo.LinkInfoString.String())
		answerPage.Execute(rw, pageInfo)
		log.Println(answerPage)
		return
	}

	// Filling page info.
	pageInfo.LinkInfoString.WriteString(fmt.Sprintf("<h3>Results for nickname \"%s\":</h3>\n\t\t<ul>\n", nick))
	for _, user := range users {
		if user.IsAvailable {
			pageInfo.LinkInfoString.WriteString(
				fmt.Sprintf("\t\t\t<li>\n\t\t\t\t<a name=\"%s\" href=\"%s\">%s: %s</a>\n\t\t\t</li>\t\n",
					user.SocialNetwork, user.Link, user.SocialNetwork, user.Name,
				))
		} else {
			pageInfo.LinkInfoString.WriteString(
				fmt.Sprintf("\t\t\t<li>\n\t\t\t\t<a name=\"%s\">%s: %s</a>\n\t\t\t</li>\t\n",
					user.SocialNetwork, user.SocialNetwork, user.Link))
		}
	}
	pageInfo.LinkInfoString.WriteString("\t\t</ul>")
	log.Println(pageInfo)
	pageInfo.LinkInfo = template.HTML(pageInfo.LinkInfoString.String())

	// Sending data to html.
	answerPage.Execute(rw, pageInfo)
	log.Println(answerPage)
}
