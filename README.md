# Twitgo


This is an Twitter REST API wrapper written in Go (Under construction)

## Authentication

 First of all you must to create an application in apps.twitter.com after that, you should generate the
  - Consumer Key
  - Consumer Secret
  - Access Token
  - Access Token Secret

With the credentials, it must create the configuration structure with this parameters, to get the right instance of Twitter API.


```go
ck  := "CUSTOMER_KEY"
cs  := "CUSTOMER_SECRET"
at  := "ACCESS_TOKEN"
ats := "ACCESS_TOKEN_SECRET"

c := &twitgo.Config{
		ConsumerKey:    ck,
		ConsumerSecret: cs,
		AccessToken:    at,
		AccessSecret:   ats}

api, err := twitgo.NewTwitterAPI(c)
```

### Usage

#### Search

```go
result, err := api.Search("GoLang")

if err != nil {
  fmt.Println(err)
}

tweets := result.Statuses
for _, tweet := range (tweets) {
  fmt.Println(tweet.Text)
}
```
