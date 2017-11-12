package rssemitter_test

import (
	"testing"

	"github.com/Bottleneck/rssemitter"
)

func TestAddRSSLink(t *testing.T) {

	sampleRSSURL := "http://www.sunstar.com.ph/rss/cebu"

	rssInstance := rssemitter.NewFeedEmitter()

	rssInstance.Add(sampleRSSURL)

}

func TestListOfRss(t *testing.T) {

	rssInstace := rssemitter.NewFeedEmitter()

	rssData := rssInstace.List()

	if rssData == nil {
		t.Errorf("RSS Feed is Empty")
	}

}
