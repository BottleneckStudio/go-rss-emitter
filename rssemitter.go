package rssemitter

import (
	"fmt"

	"github.com/chuckpreslar/emission"
	"github.com/mmcdole/gofeed"
)

type FeedEmitter interface {
	Add(gofeed.Feed) void
	List() []gofeed.Feed
	Remove(gofeed.Feed) void
	Destroy() void
}
