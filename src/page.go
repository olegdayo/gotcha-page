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
	if r.Method == http.MethodPost {
		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalln("Reading error")
			return
		}

		ans, err := getAns(buf)
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
}

// Checking which checkboxes are set.
func setUsedLinks(info *Info, container *RequesterContainer) {
	for _, parser := range info.Parsers {
		if _, ok := container.Requesters[parser]; ok {
			container.Requesters[parser].SetAvailability(true)
		}
	}
}

// Adds answers to page.
func getAns(buf []byte) (ans []byte, err error) {
	var info *Info = new(Info)
	log.Println(string(buf))
	err = json.Unmarshal(buf, info)
	if err != nil {
		log.Println("Unmarshal error")
		return nil, err
	}

	// Container initialization and execution.
	container := NewRequesterContainer(info.Nickname)
	setUsedLinks(info, container)

	users := container.GetLinks()
	ans, err = json.Marshal(users)
	if err != nil {
		log.Fatalln("Marshal error")
		return nil, err
	}

	return ans, nil
}
