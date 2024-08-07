package requesters

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"strings"
)

type SocialNetworkRequester struct {
	// Social network name.
	// For github it will be "GitHub".
	name string
	// Home page url without "https://".
	// For github it will be "github.com".
	mainURL string
	// User's nickname.
	// For example, "olegdayo".
	nickname string
	// Requester availability.
	// If selected, it can be used to parse site.
	selected bool
}

// NewSocialNetworkRequester is a constructor.
func NewSocialNetworkRequester(name string, mainURL string, nickname string) (snr *SocialNetworkRequester) {
	snr = &SocialNetworkRequester{
		name:     name,
		mainURL:  mainURL,
		nickname: nickname,
		selected: false,
	}
	return snr
}

// GetName gets name of a social network.
func (snr *SocialNetworkRequester) GetName() (name string) {
	return snr.name
}

// GetNickname gets nickname of a user.
func (snr *SocialNetworkRequester) GetNickname() (nickname string) {
	return snr.nickname
}

// IsSelected shows if requester is available.
func (snr *SocialNetworkRequester) IsSelected() (selected bool) {
	return snr.selected
}

// SetAvailability sets availability condition.
func (snr *SocialNetworkRequester) SetAvailability(cond bool) {
	snr.selected = cond
}

// GetInfo gets url and name of user by their nickname.
func (snr *SocialNetworkRequester) GetInfo() (url string, name string, err error) {
	var link string = fmt.Sprintf("https://%s/", snr.mainURL) + snr.nickname

	// Getting response.
	page, err := http.Get(link)
	if err != nil {
		return "", "", err
	}

	// Closing response before leaving the function.
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(page.Body)

	if page.StatusCode == 404 {
		// Page not found.
		return "page not found", "", errors.New("page not found")
	} else if page.StatusCode != 200 {
		// Some other error.
		// For example, 403 forbidden.
		return "", "", errors.New(fmt.Sprintf("status code is %d", page.StatusCode))
	}

	// Getting goquery document.
	info, err := goquery.NewDocumentFromReader(page.Body)
	if err != nil {
		return "", "", err
	}

	// The link is ok -> sending it and getting user's name from <title> tag.
	return strings.TrimSpace(link), strings.TrimSpace(info.Find("title").Text()), nil
}
