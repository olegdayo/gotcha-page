package parsers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type VKParser struct {
	Name     string
	nickname string
}

func NewVKParser(nickname string) *VKParser {
	vkp := new(VKParser)
	vkp.Name = "VK"
	vkp.nickname = nickname
	return vkp
}

func (vkp *VKParser) GetName() string {
	return vkp.Name
}

func (vkp *VKParser) GetLink() (string, error) {
	var link string = "https://vk.com/" + vkp.nickname

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
