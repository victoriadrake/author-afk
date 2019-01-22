package main

import (
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

var (
	consumerKey       = getenv("TWITTER_CONSUMER_KEY")
	consumerSecret    = getenv("TWITTER_CONSUMER_SECRET")
	accessToken       = getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = getenv("TWITTER_ACCESS_TOKEN_SECRET")
	rssFeed           = getenv("RSS_FEEDS")
	prefix			  = getenv("PREFIX")
	suffix			  = getenv("SUFFIX")
)

// MyResponse for AWS SAM
type MyResponse struct {
	StatusCode string `json:"StatusCode"`
	Message    string `json:"Body"`
}

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		log.Print("no environment variable " + name)
	}
	return v
}

func getRSS(rssFeed string) []string {
	feeds := os.Getenv("RSS_FEEDS")
	feedlist := strings.Split(feeds, ";")
	return feedlist
}

func tweetFeed() (MyResponse, error) {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	api.SetLogger(anaconda.BasicLogger)

	rand.Seed(time.Now().UnixNano())

	feeds := getRSS(rssFeed)
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
