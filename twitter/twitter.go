package twitter

import (
	"context"
	twitterscraper "github.com/n0madic/twitter-scraper"
)

type Tweets struct {
	Tweet string
}

func GetTweets() (rTweet []Tweets) {
	scraper := twitterscraper.New()
	// var getT []Tweets
	for tweet := range scraper.SearchTweets(context.Background(),
		"hiring #hiring junior developer remote -filter:retweets since:2023-01-01 -?", 50) {
		if tweet.Error != nil {
			panic(tweet.Error)
		}
		getT := Tweets{Tweet: tweet.Text}
		rTweet = append(rTweet, getT)
	}
	return
}
