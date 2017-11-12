package rssemitter

import (
	// "fmt"

	// "github.com/chuckpreslar/emission"
	"github.com/mmcdole/gofeed"
)

// FeedEmitter is our public interface
type FeedEmitter interface {
	Add(gofeed.Feed)
	List() []gofeed.Feed
	Remove(gofeed.Feed)
	Destroy()
}

// RssFeedEmitter is the base model for our plugin
type RssFeedEmitter struct {
	feeds  []gofeed.Feed
	parser *gofeed.Parser
}

// NewFeedEmitter creates a new instance of RssFeedEmitter
func NewFeedEmitter() *RssFeedEmitter {
	feedEmitter := RssFeedEmitter{
		parser: gofeed.NewParser(),
	}
	return &feedEmitter
}

// Add method expects a url string and adds it to the
// list of feeds in the RssFeedEmitter
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

// List returns the current list of feeds
func (r RssFeedEmitter) List() *[]gofeed.Feed {
	return &r.feeds
}

// Remove expects a string and removes the feed from the list
func (r RssFeedEmitter) Remove(url string) {
	for i := 0; i < len(r.feeds); i++ {
		if r.feeds[i].FeedLink == url {
			r.feeds = append(r.feeds[:i], r.feeds[i+1:]...)
		}
	}
}

// Destroy clears the feed emitter
func (r RssFeedEmitter) Destroy() *RssFeedEmitter {
	r = *NewFeedEmitter()
	return &r
}
