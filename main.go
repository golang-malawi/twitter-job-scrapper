package main

import (
	"fmt"
	"github.com/preciousnyasulu/twitter-job-scrapper/twitter"
	// "github.com/preciousnyasulu/twitter-job-scrapper/telegram"
)

func main() {
	tweets := twitter.GetTweets("software developer")
	for _,value := range tweets{
		fmt.Println(value)
	}

	// telegram.GetUpdate()
}
