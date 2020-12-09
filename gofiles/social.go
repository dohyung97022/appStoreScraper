package main

//-----------------------------social-----------------------------
type social struct {
	email        string
	facebookURL  string
	instagramURL string
	twitterURL   string
	twitchURL    string
	youtubeURL   string
}

func newSocial() *social { return new(social) }
