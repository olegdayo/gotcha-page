package main

import (
	"gotchaPage/requesters"
	"log"
	"sort"
)

// RequesterAvailability struct which shows if we will use the requester or not.
type RequesterAvailability struct {
	// Requester itself.
	requester requesters.Requester
	// Is it available or not.
	Available bool
}

// UserInfo is a struct with all user info.
type UserInfo struct {
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
	Requesters map[string]RequesterAvailability
}

// NewRequesterContainer initializes all requesters we have.
// NewRequesterContainer sets requesters availability to false statement.
func NewRequesterContainer(nickname string) *RequesterContainer {
	pc := new(RequesterContainer)
	pc.Requesters = map[string]RequesterAvailability{
		"github":    {requesters.NewSocialNetworkRequester("Github", "github.com", nickname), false},
		"gitlab":    {requesters.NewSocialNetworkRequester("Gitlab", "gitlab.com", nickname), false},
		"instagram": {requesters.NewSocialNetworkRequester("Instagram", "instagram.com", nickname), false},
		"telegram":  {requesters.NewTelegramRequester(nickname), false},
		"vk":        {requesters.NewSocialNetworkRequester("VK", "vk.com", nickname), false},
		"youtube":   {requesters.NewSocialNetworkRequester("YouTube", "youtube.com/c", nickname), false},
	}
	return pc
}

// GetLinks gets all users' with given nickname info from given slice of sites.
func (rc *RequesterContainer) GetLinks() []UserInfo {
	var links []UserInfo
	var link string
	var name string
	var err error

	for _, requesterAvailability := range rc.Requesters {
		log.Println(requesterAvailability.requester.GetName())
		// If requester is not available -> skip.
		if !requesterAvailability.Available {
			continue
		}

		// Getting info
		link, name, err = requesterAvailability.requester.GetInfo()

		if err == nil {
			// Everything is ok, adding.
			log.Println(requesterAvailability.requester.GetName() + ": " + link)
			links = append(links, UserInfo{
				SocialNetwork: requesterAvailability.requester.GetName(),
				Link:          link,
				Name:          name,
				IsAvailable:   true,
			})
		} else {
			// Error occurred.
			log.Println(requesterAvailability.requester.GetName() + ": " + err.Error())
			links = append(links, UserInfo{
				SocialNetwork: requesterAvailability.requester.GetName(),
				Link:          "page not found",
				Name:          name,
				IsAvailable:   false,
			})
		}
	}

	sort.Slice(links, func(i int, j int) bool {
		return links[i].SocialNetwork < links[j].SocialNetwork
	})
	return links
}
