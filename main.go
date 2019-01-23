package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mmcdole/gofeed"
)

// MyResponse for AWS SAM
type MyResponse struct {
	StatusCode string `json:"StatusCode"`
	Message    string `json:"Body"`
}

func getenv(name string) (string, error) {
	v := os.Getenv(name)
	if v == "" {
		return v, errors.New("no environment variable: " + name)
	}
	return v, nil
}

func getRSS(rssFeed string) ([]string, error) {
	if rssFeed == "" {
		return []string{""}, errors.New("no feeds present")
	}
	return strings.Split(rssFeed, ";"), nil
}

func tweetFeed() (MyResponse, error) {

	consumerKey, err := getenv("TWITTER_CONSUMER_KEY")
	consumerSecret, err := getenv("TWITTER_CONSUMER_SECRET")
	accessToken, err := getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret, err := getenv("TWITTER_ACCESS_TOKEN_SECRET")
	rssFeed, err := getenv("RSS_FEEDS")
	prefix, err := getenv("PREFIX")
	suffix, err := getenv("SUFFIX")

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	api.SetLogger(anaconda.BasicLogger)

	rand.Seed(time.Now().UnixNano())

	feeds, err := getRSS(rssFeed)
	if err != nil {
		log.Fatalf("error getting feed: %v", err.Error())
	}
	rss := feeds[rand.Intn(len(feeds))]
	fp := gofeed.NewParser()

	feed, err := fp.ParseURL(rss)

	if err != nil {
		log.Fatalf("error getting feed: %v", err)
	}

	limit := len(feed.Items)
	pick := rand.Intn(limit)
	rssItem := feed.Items[pick]
	tweet := prefix + " " + rssItem.Title + " " + rssItem.Link + " " + suffix

	resp, err := api.PostTweet(tweet, url.Values{})
	if err != nil {
		log.Fatalf("error posting tweet: %v", err)
	}
	return MyResponse{
		Message:    fmt.Sprintf("tweeted: " + resp.Text),
		StatusCode: "200",
	}, nil
}

func main() {

	lambda.Start(tweetFeed)

}
