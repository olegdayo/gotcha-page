package parsers

type TelegramParser struct {
	Name     string
	nickname string
}

func NewTelegramParser(nickname string) *TelegramParser {
	tp := new(TelegramParser)
	tp.Name = "Telegram"
	tp.nickname = nickname
	return tp
}

func (tp *TelegramParser) GetName() string {
	return tp.Name
}

func (tp *TelegramParser) GetLink() (string, error) {
	//var link string = "t.me/" + tp.nickname
	return "", nil
}
