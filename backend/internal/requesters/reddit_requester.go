package requesters

type RedditRequester struct {
	// Home page url without "https://".
	mainURL string
	// User's nickname.
	// For example, "olegsama".
	nickname string
	// Requester availability.
	// If selected, it can be used to parse site.
	selected bool
}

// NewRedditRequester is a constructor.
func NewRedditRequester(nickname string) (rr *RedditRequester) {
	rr = &RedditRequester{
		mainURL:  "",
		nickname: nickname,
		selected: false,
	}
	return rr
}

// GetName gets name of a telegram.
func (rr *RedditRequester) GetName() (name string) {
	return "Reddit"
}

// GetNickname gets nickname of a user.
func (rr *RedditRequester) GetNickname() (nickname string) {
	return rr.nickname
}

// IsSelected shows if requester is available.
func (rr *RedditRequester) IsSelected() (selected bool) {
	return rr.selected
}

// SetAvailability sets availability condition.
func (rr *RedditRequester) SetAvailability(cond bool) {
	rr.selected = cond
}

// GetInfo gets url and name of user by their nickname.
func (rr *RedditRequester) GetInfo() (url string, name string, err error) {
	return "", "", nil
}
