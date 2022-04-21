package requesters

import (
	"context"
	"errors"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type RedditRequester struct {
	// Social network name.
	// For github it will be "Github".
	Name string
	// Home page url without "https://".
	// For github it will be "github.com".
	MainURL string
	// User's nickname.
	// For example, "OFFLUCK".
	Nickname string
	// Requester availability.
	// If available, it can be used to parse site.
	Available bool
}

// NewRedditRequester is a constructor.
func NewRedditRequester(name string, mainURl string, nickname string) *RedditRequester {
	rr := new(RedditRequester)
	rr.Name = name
	rr.MainURL = mainURl
	rr.Nickname = nickname
	rr.Available = false
	return rr
}

// GetName gets name of a reddit.
func (rr *RedditRequester) GetName() string {
	return rr.Name
}

// GetNickname gets nickname of a user.
func (rr *RedditRequester) GetNickname() string {
	return rr.Nickname
}

// IsAvailable shows if requester is available.
func (rr *RedditRequester) IsAvailable() bool {
	return rr.Available
}

// SetAvailability sets availability condition.
func (rr *RedditRequester) SetAvailability(cond bool) {
	rr.Available = cond
}

//GetInfo gets url and name of user by their nickname.
func (rr *RedditRequester) GetInfo() (string, string, error) {
	client := reddit.DefaultClient()

	user, _, err := client.User.Get(context.Background(), rr.Nickname)
	if err != nil {
		return "", "", errors.New("cannot reach user for now")
	}

	return rr.MainURL + "/" + rr.Nickname, user.Name, nil
}
