package data

import (
	"time"

	"appengine"
	"appengine/datastore"
)

type Feed struct {
	Title       string
	Link        string
	Description string
	Date        time.Time

	ImageUrl   string
	ImageTitle string
	ImageLink  string

	Url     string
	Updated time.Time
	Created time.Time
}

type FeedItem struct {
	Title       string
	Link        string
	Guid        string
	Date        time.Time
	Description string

	MediaUrl    string
	MediaLength int
	MediaType   string

	Feed    Feed
	Created time.Time
	Updated time.Time
	Content string
}

func FeedKey(c appengine.Context, url string) *datastore.Key {
	return datastore.NewKey(c, "Feed", url, 0, nil)
}

func RegisterFeed(c appengine.Context, url string) (*datastore.Key, *Feed, error) {
	feed := new(Feed)

	key := FeedKey(c, url)
	err := datastore.Get(c, key, feed)

	if err == nil {
		return key, feed, nil
	} else if err != datastore.ErrNoSuchEntity {
		return key, feed, err
	}

	feed.Url = url
	feed.Created = time.Now()

	key, err = datastore.Put(c, key, feed)

	return key, feed, err
}
