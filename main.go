package main

import (
	"fmt"

	"github.com/preciousnyasulu/twitter-job-scrapper/twitter"
)

func main() {
	tweets := twitter.GetTweets()

	for _,value := range tweets{
		fmt.Println(value)
	}
}
