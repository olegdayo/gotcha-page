package tests

import (
	"fmt"
	"gotchaPage/requesters"
	"testing"
)

func TestNewGithubRequester(t *testing.T) {
	expected := &requesters.GithubRequester{
		Name:     "Github",
		Nickname: "CoolNickName",
	}
	got := requesters.NewGithubRequester("CoolNickName")

	if (expected.GetName() != got.GetName()) || (expected.Nickname != got.Nickname) {
		t.Error("Objects are not equal!")
		fmt.Print("Expected: ")
		fmt.Println(expected)
		fmt.Print("Got: ")
		fmt.Println(got)
	}
}
