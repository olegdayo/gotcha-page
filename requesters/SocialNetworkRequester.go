package requesters

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

type SocialNetworkRequester struct {
	Name     string
	MainURL  string
	Nickname string
}

func NewSocialNetworkRequester(name string, mainURl string, nickname string) *SocialNetworkRequester {
	snr := new(SocialNetworkRequester)
	snr.Name = name
	snr.MainURL = mainURl
	snr.Nickname = nickname
	return snr
}

func (snr *SocialNetworkRequester) GetName() string {
	return snr.Name
}

func (snr *SocialNetworkRequester) GetInfo() (string, string, error) {
	var link string = fmt.Sprintf("https://%s/", snr.MainURL) + snr.Nickname

	page, err := http.Get(link)
	if err != nil {
		return "", "", err
	}

	defer page.Body.Close()

	if page.StatusCode == 404 {
		return "page not found", "", errors.New("page not found")
	} else if page.StatusCode != 200 {
		return "", "", errors.New(fmt.Sprintf("Status code is %d", page.StatusCode))
	}

	info, err := goquery.NewDocumentFromReader(page.Body)

	if err != nil {
		return "", "", err
	}

	return strings.TrimSpace(link), strings.TrimSpace(info.Find("title").Text()), nil
}
