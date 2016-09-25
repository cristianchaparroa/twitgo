package twitgo

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const (
	OAUTH_TOKEN_TWITTER = "https://api.twitter.com/oauth2/token"
	SEARCH_TWITTER      = "https://api.twitter.com/1.1/search/tweets.json?q="
)

type Config struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
	Debug          bool
}
type BearerToken struct {
	AccessToken string `json:"access_token"`
}

type TwitGo struct {
	Config *Config
	Token  *BearerToken
}

func NewTwitterAPI(c *Config) (*TwitGo, error) {
	t := &TwitGo{Config: c}
	tb, err := t.token()
	fmt.Println(tb.AccessToken)
	if err != nil {
		return t, err
	}
	t.Token = tb
	return t, nil
}

func (t *TwitGo) token() (*BearerToken, error) {
	client := &http.Client{}
	encodedKeySecret := b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s",
		url.QueryEscape(t.Config.ConsumerKey),
		url.QueryEscape(t.Config.ConsumerSecret))))

	reqBody := bytes.NewBuffer([]byte(`grant_type=client_credentials`))
	req, err := http.NewRequest("POST", OAUTH_TOKEN_TWITTER, reqBody)
	if err != nil {
		log.Fatal(err, client, req)
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encodedKeySecret))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Add("Content-Length", strconv.Itoa(reqBody.Len()))

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err, res)
		return nil, err
	}

	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err, respBody)
		return nil, err
	}

	var b BearerToken
	json.Unmarshal(respBody, &b)

	return &b, nil
}

//
func (t *TwitGo) Search(q string) (*SearchResult, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", SEARCH_TWITTER, q)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.Token.AccessToken))

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result SearchResult
	json.Unmarshal(respBody, &result)
	return &result, nil
}
