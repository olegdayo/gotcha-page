package main

import (
	"fmt"
	"gotchaPage/requesters"
)

type RequesterAvailability struct {
	requester requesters.Requester
	Available bool
}

type UserInfo struct {
	SocialNetwork string
	Link          string
	Name          string
	IsAvailable   bool
}

type RequesterContainer struct {
	nickname   string
	Requesters map[string]RequesterAvailability
}

func NewRequesterContainer(nickname string) *RequesterContainer {
	pc := new(RequesterContainer)
	pc.Requesters = map[string]RequesterAvailability{
		"telegram": {requesters.NewTelegramRequester(nickname), false},
		"vk":       {requesters.NewVKRequester(nickname), false},
		"github":   {requesters.NewGithubRequester(nickname), false},
	}
	return pc
}

func (rc *RequesterContainer) GetLinks() []UserInfo {
	var links []UserInfo
	var link string
	var name string
	var err error
	for _, requesterAvailability := range rc.Requesters {
		if !requesterAvailability.Available {
			continue
		}

		fmt.Print(requesterAvailability.requester.GetName() + ": ")
		link, name, err = requesterAvailability.requester.GetInfo()
		if err == nil {
			fmt.Println(link)
			links = append(links, UserInfo{
				SocialNetwork: requesterAvailability.requester.GetName(),
				Link:          link,
				Name:          name,
				IsAvailable:   true,
			})
		} else {
			fmt.Println(err)
			links = append(links, UserInfo{
				SocialNetwork: requesterAvailability.requester.GetName(),
				Link:          "page not found",
				Name:          name,
				IsAvailable:   false,
			})
		}
	}
	return links
}
