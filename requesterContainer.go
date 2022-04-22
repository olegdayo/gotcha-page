package main

import (
	"gotchaPage/requesters"
	"log"
	"sort"
	"sync"
)

// RequesterAvailability is now unused.
// RequesterAvailability struct which shows if we will use the requester or not.
type RequesterAvailability struct {
	// Requester itself.
	requester requesters.Requester
	// Is it available or not.
	Available bool
}

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
	{"facebook", "Facebook", "facebook.com"},
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
	pc.Requesters = make(map[string]requesters.Requester)
	for _, page := range Pages {
		pc.Requesters[page.ID] = requesters.NewSocialNetworkRequester(page.Name, page.URL, nickname)
	}
	return pc
}

// GetLink gets all users' with given nickname info from given site.
func GetLink(requester requesters.Requester, links *[]*UserInfo, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
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

	mutex.Lock()
	*links = append(*links, user)
	mutex.Unlock()
}

// GetLinks gets all users' with given nickname info from given slice of sites.
func (rc *RequesterContainer) GetLinks() []*UserInfo {
	var links []*UserInfo
	wg := sync.WaitGroup{}

	for _, requesterAvailability := range rc.Requesters {
		// If requester is not available -> skip.
		if !requesterAvailability.IsAvailable() {
			continue
		}

		log.Println(requesterAvailability.GetName())

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
