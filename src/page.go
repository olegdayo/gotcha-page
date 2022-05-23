package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Info struct {
	Nickname string   `json:"nickname"`
	Parsers  []string `json:"parsers"`
}

func relation(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		{
			links, err := getPages()
			if err != nil {
				log.Fatalln("Get pages info")
				return
			}

			_, err = rw.Write(links)
			if err != nil {
				log.Fatalln("Write error")
			}
		}

	case http.MethodPost:
		{
			buf, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatalln("Reading error")
				return
			}

			ans, err := getLinks(buf)
			fmt.Println(ans)

			if err != nil {
				log.Fatalln("Getting answers error")
				return
			}

			_, err = rw.Write(ans)
			if err != nil {
				log.Fatalln("Write error")
			}
		}
	default:
	}
}

// Adds answers to page.
func getPages() (pages []byte, err error) {
	pages, err = json.Marshal(Pages)
	if err != nil {
		return nil, err
	}
	return pages, nil
}

// Checking which checkboxes are set.
func setUsedLinks(info *Info, container *RequesterContainer) {
	for _, parser := range info.Parsers {
		if _, ok := container.Requesters[parser]; ok {
			fmt.Println(parser)
			container.Requesters[parser].SetAvailability(true)
		}
	}
}

// Adds answers to page.
func getLinks(selected []byte) (links []byte, err error) {
	var info *Info = new(Info)
	log.Println(string(selected))
	err = json.Unmarshal(selected, info)
	if err != nil {
		log.Println("Unmarshal error")
		return nil, err
	}

	// Container initialization and execution.
	container := NewRequesterContainer(info.Nickname)
	setUsedLinks(info, container)

	users := container.GetLinks()
	links, err = json.Marshal(users)
	if err != nil {
		log.Fatalln("Marshal error")
		return nil, err
	}

	return links, nil
}
