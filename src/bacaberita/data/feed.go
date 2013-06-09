package data

import (
	"bytes"
	"time"

	"appengine"
	"appengine/datastore"

	"bacaberita/parser"
	"bacaberita/utils"
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

func RegisterFeed(c appengine.Context, url string) (*datastore.Key, *Feed, error) {
	feed := new(Feed)
	feed.Url = url
	feed.Created = time.Now()

	key := feed.NewKey(c)
	err := datastore.Get(c, key, feed)

	if err == nil {
		return key, feed, nil
	} else if err != datastore.ErrNoSuchEntity {
		return key, feed, err
	}

	key, err = datastore.Put(c, key, feed)

	return key, feed, err
}

func GetFeed(c appengine.Context, url string) (*Feed, error) {
	feed := new(Feed)
	feed.Url = url

	key := feed.NewKey(c)
	err := datastore.Get(c, key, feed)

	return feed, err
}

func UpdateFeed(c appengine.Context, feed *Feed) (*datastore.Key, error) {
	content, err := utils.Download(c, feed.Url)
	if err != nil {
		return nil, err
	}

	r := bytes.NewBuffer(content)
	data, err := parser.ParseRSS(r)
	if err != nil {
		return nil, err
	}

	feed.UpdateFromParser(data)
	feed.Updated = time.Now()

	key := feed.NewKey(c)
	datastore.Put(c, key, feed)

	err = StoreFeedItems(c, data, key)

	return key, err
}

func (feed *Feed) UpdateFromParser(data *parser.Feed) {
	if data.Title != nil {
		feed.Title = *data.Title
	}
	if data.Link != nil {
		feed.Link = *data.Link
	}
	if data.Description != nil {
		feed.Description = *data.Description
	}
	if data.Date != nil {
		feed.Date = *data.Date
	}
	if data.ImageUrl != nil {
		feed.ImageUrl = *data.ImageUrl
	}
	if data.ImageTitle != nil {
		feed.ImageTitle = *data.ImageTitle
	}
	if data.ImageLink != nil {
		feed.ImageLink = *data.ImageLink
	}
}

func (feed *Feed) NewKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Feed", feed.Url, 0, nil)
}
