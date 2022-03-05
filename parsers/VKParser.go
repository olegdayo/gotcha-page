package parsers

import (
	"errors"
	"fmt"
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
	defer page.Body.Close()
	if page.StatusCode == 404 {
		return "", errors.New("page not found")
	} else if page.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("Status code is %d", page.StatusCode))
	}

	return link, nil
}
