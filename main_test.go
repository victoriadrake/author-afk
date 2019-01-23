package main

import (
	"os"
	"reflect"
	"testing"
)

var SUFFIX = "hashtag"

func Test_getenv(t *testing.T) {
	if os.Getenv("SUFFIX") == "" {
		t.Skip("skipping test; SUFFIX not set")
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "envString", args: args{"SUFFIX"}, want: "hashtag", wantErr: false},
		{name: "noString", args: args{""}, want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getenv(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("getenv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getenv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRSS(t *testing.T) {
	type args struct {
		rssFeed string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "twoFeeds", args: args{"https://blog.com/rss.xml;https://blog2.com/rss.xml"}, want: []string{"https://blog.com/rss.xml", "https://blog2.com/rss.xml"}, wantErr: false},
		{name: "noFeeds", args: args{""}, want: []string{""}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getRSS(tt.args.rssFeed)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRSS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRSS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tweetFeed(t *testing.T) {
	if os.Getenv("RSS_FEEDS") == "" {
		t.Skip("skipping test; RSS_FEEDS not set")
	}
	tests := []struct {
		name    string
		want    MyResponse
		wantErr bool
	}{
		{name: "doesItTweet", want: MyResponse{Message: "tweeted: ", StatusCode: "200"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tweetFeed()
			if (err != nil) != tt.wantErr {
				t.Errorf("tweetFeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tweetFeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
