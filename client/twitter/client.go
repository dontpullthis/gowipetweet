package twitter

import (
	"fmt"
	"log"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	oauth_twitter "github.com/dghubble/oauth1/twitter"
	"github.com/dontpullthis/gowipetweet/system/config"
)

var ClientInstance Client

type Client struct {
	Client *twitter.Client
}

func newClient(client *twitter.Client) Client {
	return Client{
		client,
	}
}

func (c Client) MustDeleteTweet(tweetIdString string) {
	tweetId, err := strconv.ParseInt(tweetIdString, 10, 64)
	if err != nil {
		log.Fatal("Failed to convert tweet ID "+tweetIdString+" to int. ", err)
	}

	_, _, err = c.Client.Statuses.Destroy(tweetId, nil)
	serr, ok := err.(twitter.APIError)
	if ok {
		switch serr.Errors[0].Code {
		case 144:
			fmt.Printf("WARN: tweet with ID %s does not exist\n", tweetIdString)
		case 179:
			fmt.Printf("WARN: tweet with ID %s is restricted\n", tweetIdString)
		}
	} else if err != nil {
		log.Fatal("Failed to delete a tweet with ID "+tweetIdString+". ", err)
	}
}

func MustInitialize(cfg *config.Config) bool {
	isConfigUpdated := false
	if cfg.Auth.Token == "" || cfg.Auth.TokenSecret == "" {
		token, tokenSecret := mustGetTokens(cfg)
		cfg.Auth.Token = token
		cfg.Auth.TokenSecret = tokenSecret

		isConfigUpdated = true
	}

	tokenObject := oauth1.NewToken(cfg.Auth.Token, cfg.Auth.TokenSecret)
	config := oauth1.Config{
		ConsumerKey:    cfg.Auth.ConsumerKey,
		ConsumerSecret: cfg.Auth.ConsumerKeySecret,
		CallbackURL:    "oob",
		Endpoint:       oauth_twitter.AuthorizeEndpoint,
	}
	httpClient := config.Client(oauth1.NoContext, tokenObject)

	ClientInstance = newClient(twitter.NewClient(httpClient))

	return isConfigUpdated
}
