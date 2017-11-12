package rssemitter_test

import (
	"testing"

	"github.com/BottleneckStudio/rssemitter"
)

func TestRSSInstance(t *testing.T) {
	rssInstance := rssemitter.NewFeedEmitter()
	t.Log(rssInstance)
	if rssInstance == nil {
		t.Errorf("RSS Emitter was not initialized")
	}
}

func TestAddRSSLink(t *testing.T) {

	sampleRSSURL := "http://www.sunstar.com.ph/rss/cebu"

	rssInstance := rssemitter.NewFeedEmitter()

	rssInstance.Add(sampleRSSURL)

	t.Logf("Length of Feed Slice after adding: %d", len(*rssInstance.List()))
	if len(*rssInstance.List()) != 1 {
		t.Errorf("RSS Feed still empty")
	}
}

func TestListOfRss(t *testing.T) {

	rssInstance := rssemitter.NewFeedEmitter()

	rssInstance.Add("http://www.sunstar.com.ph/rss/cebu")

	rssData := rssInstance.List()
	t.Logf("Length of Feed Slice after adding: %d", len(*rssInstance.List()))
	if len(*rssData) == 0 {
		t.Errorf("RSS Feed length does not match actual number of feeds")
	}
}

func TestRemoveOfRss(t *testing.T) {

	rssInstance := rssemitter.NewFeedEmitter()
	testFeedUrl := "http://www.sunstar.com.ph/rss/cebu"
	rssInstance.Add(testFeedUrl)

	t.Log(len(*rssInstance.List()))
	t.Logf("Length of Feed Slice after adding: %d", len(*rssInstance.List()))
	if len(*rssInstance.List()) == 0 {
		t.Errorf("RSS Feed length does not match actual number of feeds")
	}

	rssInstance.Remove(testFeedUrl)
	t.Logf("Length of Feed Slice after removing: %d", len(*rssInstance.List()))
	if len(*rssInstance.List()) > 0 {
		t.Errorf("RSS Feed should be empty")
	}
}
