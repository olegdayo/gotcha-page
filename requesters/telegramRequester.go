package requesters

type TelegramRequester struct {
	Name     string
	Nickname string
}

// NewTelegramRequester is a constructor.
func NewTelegramRequester(nickname string) *TelegramRequester {
	tr := new(TelegramRequester)
	tr.Name = "Telegram"
	tr.Nickname = nickname
	return tr
}

// GetName gets name of a social network.
func (tr *TelegramRequester) GetName() string {
	return tr.Name
}

// GetNickname gets nickname of a user.
func (tr *TelegramRequester) GetNickname() string {
	return tr.Nickname
}

// GetInfo gets url and name of user by their nickname.
func (tr *TelegramRequester) GetInfo() (string, string, error) {
	// var link string = "t.me/" + tr.nickname
	// TODO...
	return "", "", nil
}
