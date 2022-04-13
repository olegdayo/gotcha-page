package tests

import (
	"fmt"
	"gotchaPage/requesters"
	"testing"
)

func TestNewSocialNetworkRequester(t *testing.T) {
	testCases := []struct {
		name     string
		mainURL  string
		nickname string
		expected *requesters.SocialNetworkRequester
	}{
		{
			name:     "VK",
			mainURL:  "vk.com",
			nickname: "olegsama",
			expected: &requesters.SocialNetworkRequester{
				Name:      "VK",
				MainURL:   "vk.com",
				Nickname:  "olegsama",
				Available: false,
			},
		},
		{
			name:     "Github",
			mainURL:  "github.com",
			nickname: "OFFLUCK",
			expected: &requesters.SocialNetworkRequester{
				Name:      "Github",
				MainURL:   "github.com",
				Nickname:  "OFFLUCK",
				Available: false,
			},
		},
		{
			name:     "Gitlab",
			mainURL:  "gitlab.com",
			nickname: "OFFLUCK",
			expected: &requesters.SocialNetworkRequester{
				Name:      "Gitlab",
				MainURL:   "gitlab.com",
				Nickname:  "OFFLUCK",
				Available: false,
			},
		},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprintf("Test№%d", index), func(t *testing.T) {
			got := requesters.NewSocialNetworkRequester(
				testCase.name,
				testCase.mainURL,
				testCase.nickname,
			)

			if *testCase.expected != *got {
				t.Errorf("Error while testing %s", testCase.expected.GetName())
			}
		})
	}
}

func TestGetName(t *testing.T) {
	testCases := []struct {
		requester *requesters.SocialNetworkRequester
		expected  string
	}{
		{
			requesters.NewSocialNetworkRequester(
				"VK",
				"vk.com",
				"olegsama",
			),
			"VK",
		},
		{
			requesters.NewSocialNetworkRequester(
				"Github",
				"github.com",
				"OFFLUCK",
			),
			"Github",
		},
		{
			requesters.NewSocialNetworkRequester(
				"Gitlab",
				"gitlab.com",
				"OFFLUCK",
			),
			"Gitlab",
		},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprintf("Test№%d", index), func(t *testing.T) {
			got := testCase.requester.GetName()
			if testCase.expected != got {
				t.Errorf("Expected: %s; got: %s\n", testCase.expected, got)
			}
		})
	}
}

func TestGetNickname(t *testing.T) {
	testCases := []struct {
		requester *requesters.SocialNetworkRequester
		expected  string
	}{
		{
			requesters.NewSocialNetworkRequester(
				"VK",
				"vk.com",
				"olegsama",
			),
			"olegsama",
		},
		{
			requesters.NewSocialNetworkRequester(
				"Github",
				"github.com",
				"OFFLUCK",
			),
			"OFFLUCK",
		},
		{
			requesters.NewSocialNetworkRequester(
				"Gitlab",
				"gitlab.com",
				"OFFLUCK",
			),
			"OFFLUCK",
		},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprintf("Test№%d", index), func(t *testing.T) {
			got := testCase.requester.GetNickname()
			if testCase.expected != got {
				t.Errorf("Expected: %s; got: %s\n", testCase.expected, got)
			}
		})
	}
}

func TestIsAvailable(t *testing.T) {
	testCases := []struct {
		requester *requesters.SocialNetworkRequester
		expected  bool
	}{
		{
			requester: &requesters.SocialNetworkRequester{
				Name:      "VK",
				MainURL:   "vk.com",
				Nickname:  "olegsama",
				Available: false,
			},
			expected: false,
		},
		{
			requester: &requesters.SocialNetworkRequester{
				Name:      "Github",
				MainURL:   "github.com",
				Nickname:  "OFFLUCK",
				Available: true,
			},
			expected: true,
		},
		{
			requester: &requesters.SocialNetworkRequester{
				Name:      "Gitlab",
				MainURL:   "gitlab.com",
				Nickname:  "OFFLUCK",
				Available: true,
			},
			expected: true,
		},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprintf("Test№%d", index), func(t *testing.T) {
			got := testCase.requester.IsAvailable()
			if testCase.expected != got {
				t.Errorf("Expected: %v; got: %v\n", testCase.expected, got)
			}
		})
	}
}

func TestSetAvailability(t *testing.T) {
	testCases := []struct {
		requester *requesters.SocialNetworkRequester
		expected  bool
	}{
		{
			requester: &requesters.SocialNetworkRequester{
				Name:      "VK",
				MainURL:   "vk.com",
				Nickname:  "olegsama",
				Available: false,
			},
			expected: false,
		},
		{
			requester: &requesters.SocialNetworkRequester{
				Name:      "Github",
				MainURL:   "github.com",
				Nickname:  "OFFLUCK",
				Available: true,
			},
			expected: true,
		},
		{
			requester: &requesters.SocialNetworkRequester{
				Name:      "Gitlab",
				MainURL:   "gitlab.com",
				Nickname:  "OFFLUCK",
				Available: true,
			},
			expected: true,
		},
	}

	var setter bool = true
	for index, testCase := range testCases {
		t.Run(fmt.Sprintf("Test№%d", index), func(t *testing.T) {
			testCase.requester.SetAvailability(setter)
			testCase.expected = setter
			setter = !setter
			got := testCase.requester.IsAvailable()
			if testCase.expected != got {
				t.Errorf("Expected: %v; got: %v\n", testCase.expected, got)
			}
		})
	}
}

func TestGetInfo(t *testing.T) {
	testCases := []struct {
		requester    *requesters.SocialNetworkRequester
		expectedLink string
		expectedName string
	}{
		{
			requesters.NewSocialNetworkRequester(
				"VK",
				"vk.com",
				"olegsama",
			),
			"https://vk.com/olegsama",
			"Олег Сидоренков | ВКонтакте",
		},
		{
			requesters.NewSocialNetworkRequester(
				"Github",
				"github.com",
				"OFFLUCK",
			),
			"https://github.com/OFFLUCK",
			"OFFLUCK (Oleg) · GitHub",
		},
		{
			requesters.NewSocialNetworkRequester(
				"Gitlab",
				"gitlab.com",
				"OFFLUCK",
			),
			"https://gitlab.com/OFFLUCK",
			"Oleg · GitLab",
		},
		{
			requesters.NewSocialNetworkRequester(
				"VK",
				"vk.com",
				"dmvdfcjdjk123211hj23123bhwhb1hb3j",
			),
			"page not found",
			"",
		},
		{
			requesters.NewSocialNetworkRequester(
				"Github",
				"github.com",
				"dm5vdfcj31djk2321151e34123214211hj2323e123sd211342bhwhb1hb3j",
			),
			"page not found",
			"",
		},
		{
			requesters.NewSocialNetworkRequester(
				"Youtube",
				"youtube.com/c",
				"jcsxiuaiunxiu378wbedxs78w33bd2eqw9emimads0wq9oiqwd",
			),
			"page not found",
			"",
		},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprintf("Test№%d", index), func(t *testing.T) {
			gotLink, gotName, err := testCase.requester.GetInfo()
			if err == nil {
				if gotLink != testCase.expectedLink || gotName != testCase.expectedName {
					t.Errorf("Expected: %s, %s; got: %s, %s\n", testCase.expectedLink, testCase.expectedName, gotLink, gotName)
				}
				return
			}

			if err.Error() == "page not found" {
				if gotLink != testCase.expectedLink {
					t.Errorf("Expected: %s, %s; got: %s, %s\n", testCase.expectedLink, testCase.expectedName, gotLink, gotName)
				}
				return
			}

			t.Fatalf("Unexpected error: %s", err)
		})
	}
}
