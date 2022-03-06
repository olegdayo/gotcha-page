package requesters

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GithubRequester struct {
	Name     string
	nickname string
}

func NewGithubRequester(nickname string) *GithubRequester {
	gr := new(GithubRequester)
	gr.Name = "Github"
	gr.nickname = nickname
	return gr
}

func (gr *GithubRequester) GetName() string {
	return gr.Name
}

func (gr *GithubRequester) GetLink() (string, error) {
	var link string = "https://github.com/" + gr.nickname

	page, err := http.Get(link)
	if err != nil {
		return "", err
	}

	// Trying to close the response.
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(page.Body)

	if page.StatusCode == 404 {
		return "", errors.New("page not found")
	} else if page.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("Status code is %d", page.StatusCode))
	}

	return link, nil
}
