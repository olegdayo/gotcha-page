package main

import (
	"fmt"
	"gotchaPage/requesters"
	"log"
	"sort"
)

// UserInfo is a struct with all user info.
type UserInfo struct {
	// User's nickname.
	Nickname string
	// Social network name.
	SocialNetwork string
	// User's profile link.
	Link string
	// User's name from <title> tag.
	Name string
	// User availability.
	// True if everything is ok.
	// False if during parsing an error occurred.
	IsAvailable bool
}

// RequesterContainer is a container of requesters.
type RequesterContainer struct {
	// Nickname of a user we are looking for.
	nickname string
	// Requesters.
	Requesters map[string]requesters.Requester
}

type Page struct {
	ID   string
	Name string
	URL  string
}

var Pages []*Page = []*Page{
	{
		ID:   "facebook",
		Name: "Facebook",
		URL:  "facebook.com",
	},
	{
		ID:   "github",
		Name: "GitHub",
		URL:  "github.com",
	},
	{
		ID:   "gitlab",
		Name: "GitLab",
		URL:  "gitlab.com",
	},
	{
		ID:   "instagram",
		Name: "Instagram",
		URL:  "instagram.com",
	},
	{
		ID:   "vk",
		Name: "VK",
		URL:  "vk.com",
	},
	{
		ID:   "youtube",
		Name: "Youtube",
		URL:  "youtube.com/c",
	},
}

// NewRequesterContainer initializes all requesters we have.
// NewRequesterContainer sets requesters availability to false statement.
func NewRequesterContainer(nickname string) (rc *RequesterContainer) {
	rc = new(RequesterContainer)
	rc.Requesters = make(map[string]requesters.Requester)
	for _, page := range Pages {
		rc.Requesters[page.ID] = requesters.NewSocialNetworkRequester(page.Name, page.URL, nickname)
	}
	return rc
}

// Function getNumberOfAvailableRequesters gets number of selected requesters.
func (rc *RequesterContainer) getNumberOfAvailableRequesters() (number int) {
	number = 0
	for _, requester := range rc.Requesters {
		if requester.IsSelected() {
			number++
		}
	}
	return number
}

// GetLink gets all users' with given nickname info from given site.
func GetLink(requester requesters.Requester, linksChannel chan<- *UserInfo) {
	// Getting info.
	link, name, err := requester.GetInfo()
	var user *UserInfo

	if err == nil {
		// Everything is ok, adding.
		log.Println(requester.GetName() + ": " + link)
		user = &UserInfo{
			Nickname:      requester.GetNickname(),
			SocialNetwork: requester.GetName(),
			Link:          link,
			Name:          name,
			IsAvailable:   true,
		}
	} else {
		// Error occurred.
		log.Println(requester.GetName() + ": " + err.Error())
		user = &UserInfo{
			Nickname:      requester.GetNickname(),
			SocialNetwork: requester.GetName(),
			Link:          err.Error(),
			Name:          name,
			IsAvailable:   false,
		}
	}

	linksChannel <- user
}

// GetLinks gets all users' with given nickname info from given slice of sites.
func (rc *RequesterContainer) GetLinks() (links []*UserInfo) {
	var selectedRequestersNumber int = rc.getNumberOfAvailableRequesters()
	linksChannel := make(chan *UserInfo, selectedRequestersNumber)

	for _, requester := range rc.Requesters {
		// If requester is not available -> skip.
		if !requester.IsSelected() {
			continue
		}

		log.Println(requester.GetName())
		go GetLink(requester, linksChannel)
	}

	links = make([]*UserInfo, selectedRequestersNumber)
	for i := 0; i < selectedRequestersNumber; i++ {
		links[i] = <-linksChannel
		fmt.Println(links[i])
	}
	close(linksChannel)

	sort.Slice(
		links,
		func(i int, j int) bool {
			return links[i].SocialNetwork < links[j].SocialNetwork
		},
	)
	return links
}
