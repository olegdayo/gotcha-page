package tests

import (
	"fmt"
	"gotchaPage/src/requesters"
	"testing"
)

func TestNewSocialNetworkRequester(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		mainURL  string
		nickname string
		selected bool
		expected *requesters.SocialNetworkRequester
	}{
		{
			name:     "VK",
			mainURL:  "vk.com",
			nickname: "olegsama",
			selected: false,
			expected: requesters.NewSocialNetworkRequester(
				"VK",
				"vk.com",
				"olegsama",
			),
		},
		{
			name:     "GitHub",
			mainURL:  "github.com",
			nickname: "OFFLUCK",
			selected: true,
			expected: requesters.NewSocialNetworkRequester(
				"GitHub",
				"github.com",
				"OFFLUCK",
			),
		},
		{
			name:     "GitLab",
			mainURL:  "gitlab.com",
			nickname: "OFFLUCK",
			selected: true,
			expected: requesters.NewSocialNetworkRequester(
				"GitLab",
				"gitlab.com",
				"OFFLUCK",
			),
		},
	}
	testCases[1].expected.SetAvailability(true)
	testCases[2].expected.SetAvailability(true)

	for index, testCase := range testCases {
		t.Run(fmt.Sprintf("Test№%d", index), func(t *testing.T) {
			got := requesters.NewSocialNetworkRequester(
				testCase.name,
				testCase.mainURL,
				testCase.nickname,
			)
			got.SetAvailability(index > 0)

			if *testCase.expected != *got {
				t.Errorf("Error while testing %s", testCase.expected.GetName())
			}
		})
	}
}

func TestGetName(t *testing.T) {
	t.Parallel()
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
				"GitHub",
				"github.com",
				"OFFLUCK",
			),
			"GitHub",
		},
		{
			requesters.NewSocialNetworkRequester(
				"GitLab",
				"gitlab.com",
				"OFFLUCK",
			),
			"GitLab",
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
	t.Parallel()
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
				"GitHub",
				"github.com",
				"OFFLUCK",
			),
			"OFFLUCK",
		},
		{
			requesters.NewSocialNetworkRequester(
				"GitLab",
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
	t.Parallel()
	testCases := []struct {
		requester *requesters.SocialNetworkRequester
		expected  bool
	}{
		{
			requester: requesters.NewSocialNetworkRequester(
				"VK",
				"vk.com",
				"olegsama",
			),
			expected: false,
		},
		{
			requester: requesters.NewSocialNetworkRequester(
				"GitHub",
				"github.com",
				"OFFLUCK",
			),
			expected: true,
		},
		{
			requester: requesters.NewSocialNetworkRequester(
				"GitLab",
				"gitlab.com",
				"OFFLUCK",
			),
			expected: true,
		},
	}
	testCases[1].requester.SetAvailability(true)
	testCases[2].requester.SetAvailability(true)

	for index, testCase := range testCases {
		t.Run(fmt.Sprintf("Test№%d", index), func(t *testing.T) {
			got := testCase.requester.IsSelected()
			if testCase.expected != got {
				t.Errorf("Expected: %v; got: %v\n", testCase.expected, got)
			}
		})
	}
}

func TestSetAvailability(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		requester *requesters.SocialNetworkRequester
		expected  bool
	}{
		{
			requester: requesters.NewSocialNetworkRequester(
				"VK",
				"vk.com",
				"olegsama",
			),
			expected: false,
		},
		{
			requester: requesters.NewSocialNetworkRequester(
				"GitHub",
				"github.com",
				"OFFLUCK",
			),
			expected: true,
		},
		{
			requester: requesters.NewSocialNetworkRequester(
				"GitLab",
				"gitlab.com",
				"OFFLUCK",
			),
			expected: true,
		},
	}
	testCases[1].requester.SetAvailability(true)
	testCases[2].requester.SetAvailability(true)

	var setter bool = true
	for index, testCase := range testCases {
		t.Run(fmt.Sprintf("Test№%d", index), func(t *testing.T) {
			testCase.requester.SetAvailability(setter)
			testCase.expected = setter
			setter = !setter
			got := testCase.requester.IsSelected()
			if testCase.expected != got {
				t.Errorf("Expected: %v; got: %v\n", testCase.expected, got)
			}
		})
	}
}

func TestGetInfo(t *testing.T) {
	t.Parallel()
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
				"GitHub",
				"github.com",
				"OFFLUCK",
			),
			"https://github.com/OFFLUCK",
			"OFFLUCK (Oleg) · GitHub",
		},
		{
			requesters.NewSocialNetworkRequester(
				"GitLab",
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
				"GitHub",
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
