package twitgo

import (
	"testing"
)

func TestBearerToken(t *testing.T) {
	ck := "CUSTOMER_KEY"
	cs := "CUSTOMER_SECRET"
	at := "ACCESS_TOKEN"
	ats := "ACCESS_TOKEN_SECRET"

	c := &Config{
		ConsumerKey:    ck,
		ConsumerSecret: cs,
		AccessToken:    at,
		AccessSecret:   ats}

	api, err := NewTwitterAPI(c)

	if len(api.Token.AccessToken) == 0 {
		t.Errorf("Expected a token, but get %v", err)
	}
}

func TestSearch(t *testing.T) {
	ck := "CUSTOMER_KEY"
	cs := "CUSTOMER_SECRET"
	at := "ACCESS_TOKEN"
	ats := "ACCESS_TOKEN_SECRET"

	c := &Config{
		ConsumerKey:    ck,
		ConsumerSecret: cs,
		AccessToken:    at,
		AccessSecret:   ats}

	api, err := NewTwitterAPI(c)
	if err != nil {
		t.Error(err)
	}
	result, err := api.Search("@golang")

	if err != nil {
		t.Error(err)
	}

	if len(result.Statuses) == 0 {
		t.Errorf("not results")
	}
	//tweets := result.Statuses
	//
	//for _, tweet := range (tweets) {
	//	fmt.Println(tweet.Text)
	//}
	//json, _ := prettyjson.Marshal(result)
	//fmt.Println(string(json))
}
