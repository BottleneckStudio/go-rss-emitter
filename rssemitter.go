package rssemitter

import (
	// "fmt"

	// "github.com/chuckpreslar/emission"
	"github.com/mmcdole/gofeed"
)

type FeedEmitter interface {
	Add(gofeed.Feed)
	List() []gofeed.Feed
	Remove(gofeed.Feed)
	Destroy()
}

type RssFeedEmitter struct {
	feeds  []gofeed.Feed
	parser *gofeed.Parser
}

func NewFeedEmitter() *RssFeedEmitter {
	feedEmitter := RssFeedEmitter{
		parser: gofeed.NewParser(),
	}
	return &feedEmitter
}

func (r RssFeedEmitter) Add(url string) {
	for _, feed := range r.feeds {
		if feed.FeedLink == url {
			return
		}
	}
	newFeed, err := r.parser.ParseURL(url)
	if err != nil {
		return
	}
	r.feeds = append(r.feeds, *newFeed)
}

func (r RssFeedEmitter) List() *[]gofeed.Feed {
	return &r.feeds
}

func (r RssFeedEmitter) Remove(url string) {
	for i := 0; i < len(r.feeds); i++ {
		if r.feeds[i].FeedLink == url {
			r.feeds = append(r.feeds[:i], r.feeds[i+1:]...)
		}
	}
}

func (r RssFeedEmitter) Destroy() *RssFeedEmitter {
	r = *NewFeedEmitter()
	return &r
}
