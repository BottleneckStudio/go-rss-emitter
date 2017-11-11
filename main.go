package main

import (
	"fmt"

	"github.com/chuckpreslar/emission"
	"github.com/mmcdole/gofeed"
)

func main() {
	// We can test everything here, but in the final dist, main.go will not be included
	emitter := emission.NewEmitter()
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("http://feeds.twit.tv/twit.xml")
	fmt.Println(feed.Title)
	
	hello := func(to string) {
		fmt.Printf("Hello %s!\n", to)
	}

	count := func(count int) {
		for i := 0; i < count; i++ {
			fmt.Println(i)
		}
	}

	emitter.On("hello", hello).
		On("count", count).
		Emit("hello", "world").
		Emit("count", 5)
}