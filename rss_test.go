package rssemitter_test

import (
	"testing"

	"github.com/Bottleneck/rssemitter"
)

func TestAddExecutesCorrectRss(t *testing.T) {

	sampleRSSURL := "http://www.sunstar.com.ph/rss/cebu"

	rssInstance := rssemitter.NewFeedEmitter()

	rssInstance.Add(sampleRSSURL)

}
