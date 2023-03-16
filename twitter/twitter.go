package twitter

import (
	"context"
	"fmt"
	"strings"
	twitterscraper "github.com/n0madic/twitter-scraper"
)

type Tweets struct {
	Tweet string
	Url string
}

func GetTweets(search_tags string) (rTweet []Tweets) {
	search_tags =  fmt.Sprint(strings.Split(search_tags,","))
	tags :=  strings.Trim(search_tags, "[]")
	max_results :=10

	scraper := twitterscraper.New()
	for tweet := range scraper.SearchTweets(context.Background(),
	fmt.Sprintf("%s hiring job -filter:retweets since:2023-01-01 -?",tags), max_results) {
		if tweet.Error != nil {
			panic(tweet.Error)
		}
		getT := Tweets{Tweet: tweet.Text,Url: tweet.PermanentURL}
		rTweet = append(rTweet, getT)
	}
	return
}
