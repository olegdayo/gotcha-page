package main

import "gotchaPage/parsers"

type ParserContainer struct {
	nickname string
	Parsers  []parsers.Parser
}

func NewParsersContainer(nickname string) *ParserContainer {
	pc := new(ParserContainer)
	pc.Parsers = []parsers.Parser{
		parsers.NewTelegramParser(nickname),
		parsers.NewVKParser(nickname),
	}
	return pc
}

func (pc *ParserContainer) GetLinks() [][2]string {
	var links [][2]string
	var link string
	var err error
	for _, parser := range pc.Parsers {
		link, err = parser.GetLink()
		if err == nil {
			links = append(links, [2]string{parser.GetName(), link})
		}
	}
	return links
}
