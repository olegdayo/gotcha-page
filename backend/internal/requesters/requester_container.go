package requesters

import (
	"fmt"
	"github.com/olegdayo/gotcha-page/backend/internal/sharedData"
	"log"
	"sort"
)

// UserInfo is a struct with all user info.
type UserInfo struct {
	// User's nickname.
	Nickname string `json:"nickname"`
	// Social network name.
	SocialNetwork string `json:"url"`
	// User's profile link.
	Link string `json:"link"`
	// User's name from <title> tag.
	Name string `json:"name"`
	// User availability.
	// True if everything is ok.
	// False if during parsing an error occurred.
	IsAvailable bool `json:"available"`
}

// RequesterContainer is a container of requesters.
type RequesterContainer struct {
	// Nickname of a user we are looking for.
	nickname string
	// Requesters.
	Requesters map[string]Requester
}

// NewRequesterContainer initializes all requesters we have.
// NewRequesterContainer sets requesters availability to false statement.
func NewRequesterContainer(nickname string) (rc *RequesterContainer) {
	rc = new(RequesterContainer)
	rc.Requesters = make(map[string]Requester)
	for _, network := range sharedData.GetConfig().Networks {
		rc.Requesters[network.ID] = NewSocialNetworkRequester(network.Name, network.URL, nickname)
	}
	return rc
}

// SetUsedLinks sets ticked checkboxes.
func (rc *RequesterContainer) SetUsedLinks(clients ...string) {
	for _, parser := range clients {
		if _, ok := rc.Requesters[parser]; ok {
			log.Println(parser)
			rc.Requesters[parser].SetAvailability(true)
		}
	}
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
func GetLink(requester Requester, linksChannel chan<- *UserInfo) {
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
			Link:          link,
			Name:          fmt.Sprintf("%s: %s not found", requester.GetName(), requester.GetNickname()),
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
		log.Println(links[i])
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
