package requesters

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type VKRequester struct {
	Name     string
	nickname string
}

func NewVKRequester(nickname string) *VKRequester {
	vkr := new(VKRequester)
	vkr.Name = "VK"
	vkr.nickname = nickname
	return vkr
}

func (vkr *VKRequester) GetName() string {
	return vkr.Name
}

func (vkr *VKRequester) GetLink() (string, error) {
	var link string = "https://vk.com/" + vkr.nickname

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
