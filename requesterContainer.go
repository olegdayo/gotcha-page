package main

import (
	"gotchaPage/requesters"
	"log"
	"sort"
	"sync"
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
	Requesters map[string]*RequesterAvailability
}

type Page struct {
	ID   string
	Name string
	URL  string
}

var Pages []*Page = []*Page{
	{"github", "Github", "github.com"},
	{"gitlab", "Gitlab", "gitlab.com"},
	{"instagram", "Instagram", "instagram.com"},
	{"vk", "VK", "vk.com"},
	{"youtube", "Youtube", "youtube.com/c"},
}

// NewRequesterContainer initializes all requesters we have.
// NewRequesterContainer sets requesters availability to false statement.
func NewRequesterContainer(nickname string) *RequesterContainer {
	pc := new(RequesterContainer)
	pc.Requesters = make(map[string]*RequesterAvailability)
	for _, page := range Pages {
		pc.Requesters[page.ID] = &RequesterAvailability{
			requesters.NewSocialNetworkRequester(page.Name, page.URL, nickname),
			false,
		}
	}
	return pc
}

// GetLink gets all users' with given nickname info from given site.
func GetLink(requesterAvailability *RequesterAvailability, links *[]*UserInfo, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	// Getting info
	link, name, err := requesterAvailability.requester.GetInfo()

	mutex.Lock()
	if err == nil {
		// Everything is ok, adding.
		log.Println(requesterAvailability.requester.GetName() + ": " + link)
		*links = append(*links, &UserInfo{
			SocialNetwork: requesterAvailability.requester.GetName(),
			Link:          link,
			Name:          name,
			IsAvailable:   true,
		})
	} else {
		// Error occurred.
		log.Println(requesterAvailability.requester.GetName() + ": " + err.Error())
		*links = append(*links, &UserInfo{
			SocialNetwork: requesterAvailability.requester.GetName(),
			Link:          "page not found",
			Name:          name,
			IsAvailable:   false,
		})
	}
	mutex.Unlock()
}

// GetLinks gets all users' with given nickname info from given slice of sites.
func (rc *RequesterContainer) GetLinks() []*UserInfo {
	var links []*UserInfo
	wg := sync.WaitGroup{}

	for _, requesterAvailability := range rc.Requesters {
		log.Println(requesterAvailability.requester.GetName())
		// If requester is not available -> skip.
		if !requesterAvailability.Available {
			continue
		}

		wg.Add(1)
		mutex := sync.Mutex{}
		go GetLink(requesterAvailability, &links, &wg, &mutex)
	}

	wg.Wait()
	sort.Slice(links, func(i int, j int) bool {
		return links[i].SocialNetwork < links[j].SocialNetwork
	})
	return links
}
