package main

import (
	"fmt"
	"gotchaPage/requesters"
)

type RequesterContainer struct {
	nickname   string
	Requesters []requesters.Requester
}

func NewRequesterContainer(nickname string) *RequesterContainer {
	pc := new(RequesterContainer)
	pc.Requesters = []requesters.Requester{
		requesters.NewTelegramRequester(nickname),
		requesters.NewVKRequester(nickname),
		requesters.NewGithubRequester(nickname),
	}
	return pc
}

func (rc *RequesterContainer) GetLinks() [][2]string {
	var links [][2]string
	var link string
	var err error
	for _, requester := range rc.Requesters {
		fmt.Print(requester.GetName() + ":")
		link, err = requester.GetLink()
		if err == nil {
			fmt.Println(link)
			links = append(links, [2]string{requester.GetName(), link})
		} else {
			fmt.Println(err)
		}
	}
	return links
}
