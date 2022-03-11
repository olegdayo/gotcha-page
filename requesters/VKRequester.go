package requesters

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
)

type VKRequester struct {
	Name     string
	Nickname string
}

func NewVKRequester(nickname string) *VKRequester {
	vkr := new(VKRequester)
	vkr.Name = "VK"
	vkr.Nickname = nickname
	return vkr
}

func (vkr *VKRequester) GetName() string {
	return vkr.Name
}

func (vkr *VKRequester) GetInfo() (string, string, error) {
	var link string = "https://vk.com/" + vkr.Nickname

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
		return "", "", errors.New("page not found")
	} else if page.StatusCode != 200 {
		return "", "", errors.New(fmt.Sprintf("Status code is %d", page.StatusCode))
	}

	info, err := goquery.NewDocumentFromReader(page.Body)

	if err != nil {
		panic(err)
	}

	return link, info.Find("title").Text(), nil
}
