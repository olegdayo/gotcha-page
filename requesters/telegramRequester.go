package requesters

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

type TelegramRequester struct {
	// Home page url without "https://".
	mainURL string
	// User's nickname.
	// For example, "olegsama".
	nickname string
	// Requester availability.
	// If selected, it can be used to parse site.
	selected bool
}

// NewTelegramRequester is a constructor.
func NewTelegramRequester(mainURL string, nickname string) *TelegramRequester {
	tr := &TelegramRequester{
		mainURL:  mainURL,
		nickname: nickname,
		selected: false,
	}
	return tr
}

// GetName gets name of a telegram.
func (tr *TelegramRequester) GetName() string {
	return "Telegram"
}

// GetNickname gets nickname of a user.
func (tr *TelegramRequester) GetNickname() string {
	return tr.nickname
}

// IsSelected shows if requester is available.
func (tr *TelegramRequester) IsSelected() bool {
	return tr.selected
}

// SetAvailability sets availability condition.
func (tr *TelegramRequester) SetAvailability(cond bool) {
	tr.selected = cond
}

// GetInfo gets url and name of user by their nickname.
func (tr *TelegramRequester) GetInfo() (string, string, error) {
	var link string = fmt.Sprintf("https://%s/", tr.mainURL) + tr.nickname

	// Getting response.
	page, err := http.Get(link)
	if err != nil {
		return "", "", err
	}

	// Closing response before leaving the function.
	defer page.Body.Close()

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
