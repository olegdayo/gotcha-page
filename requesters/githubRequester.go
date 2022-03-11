package requesters

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
)

type GithubRequester struct {
	Name     string
	Nickname string
}

func NewGithubRequester(nickname string) *GithubRequester {
	gr := new(GithubRequester)
	gr.Name = "Github"
	gr.Nickname = nickname
	return gr
}

func (gr *GithubRequester) GetName() string {
	return gr.Name
}

func (gr *GithubRequester) GetInfo() (string, string, error) {
	var link string = "https://github.com/" + gr.Nickname

	page, err := http.Get(link)
	if err != nil {
		return "", "", err
	}

	// Trying to close the response.
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(page.Body)

	if page.StatusCode == 404 {
		return "page not found", "", errors.New("page not found")
	} else if page.StatusCode != 200 {
		return "", "", errors.New(fmt.Sprintf("Status code is %d", page.StatusCode))
	}

	info, err := goquery.NewDocumentFromReader(page.Body)

	if err != nil {
		panic(err)
	}

	return link, info.Find("title").Text(), nil
}
