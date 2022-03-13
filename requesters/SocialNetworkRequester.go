package requesters

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
)

type SocialNetworkRequester struct {
	Name     string
	mainURL  string
	Nickname string
}

func NewSocialNetworkRequester(name string, mainURl string, nickname string) *SocialNetworkRequester {
	snr := new(SocialNetworkRequester)
	snr.Name = name
	snr.mainURL = mainURl
	snr.Nickname = nickname
	return snr
}

func (snr *SocialNetworkRequester) GetName() string {
	return snr.Name
}

func (snr *SocialNetworkRequester) GetInfo() (string, string, error) {
	var link string = fmt.Sprintf("https://%s/", snr.mainURL) + snr.Nickname

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
