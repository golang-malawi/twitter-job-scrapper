package twitter

import (
	"context"
	"fmt"
	"strings"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

type Tweets struct {
	Tweet string
}

func GetTweets(search_tags string) (rTweet []Tweets) {
	search_tags =  fmt.Sprint(strings.Split(search_tags,","))
	tags :=  strings.Trim(search_tags, "[]")

	scraper := twitterscraper.New()
	for tweet := range scraper.SearchTweets(context.Background(),
	tags + " hiring job -filter:retweets since:2023-01-01 -?", 50) {
		if tweet.Error != nil {
			panic(tweet.Error)
		}
		getT := Tweets{Tweet: tweet.Text}
		rTweet = append(rTweet, getT)
	}
	return
}
