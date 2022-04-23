package requesters

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
func NewTelegramRequester(nickname string) (tr *TelegramRequester) {
	tr = &TelegramRequester{
		mainURL:  "",
		nickname: nickname,
		selected: false,
	}
	return tr
}

// GetName gets name of a telegram.
func (tr *TelegramRequester) GetName() (name string) {
	return "Telegram"
}

// GetNickname gets nickname of a user.
func (tr *TelegramRequester) GetNickname() (nickname string) {
	return tr.nickname
}

// IsSelected shows if requester is available.
func (tr *TelegramRequester) IsSelected() (selected bool) {
	return tr.selected
}

// SetAvailability sets availability condition.
func (tr *TelegramRequester) SetAvailability(cond bool) {
	tr.selected = cond
}

// GetInfo gets url and name of user by their nickname.
func (tr *TelegramRequester) GetInfo() (url string, name string, err error) {
	return "", "", nil
}
