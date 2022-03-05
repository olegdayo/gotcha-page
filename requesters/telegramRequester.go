package requesters

type TelegramRequester struct {
	Name     string
	nickname string
}

func NewTelegramRequester(nickname string) *TelegramRequester {
	tr := new(TelegramRequester)
	tr.Name = "Telegram"
	tr.nickname = nickname
	return tr
}

func (tr *TelegramRequester) GetName() string {
	return tr.Name
}

func (tr *TelegramRequester) GetLink() (string, error) {
	//var link string = "t.me/" + tr.nickname
	return "", nil
}
