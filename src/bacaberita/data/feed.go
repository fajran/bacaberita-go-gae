package data

import (
	"bytes"
	"fmt"
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

func StoreFeedItems(c appengine.Context, data *parser.Feed, parent *datastore.Key) error {
	for _, item := range data.Items {
		feedItem := new(FeedItem)
		feedItem.UpdateFromParser(item)

		key := feedItem.NewKey(c, parent)

		key, err := datastore.Put(c, key, feedItem)
		if err != nil {
			c.Errorf("Error inserting item: url=%s error=%w", feedItem.Link, err)
			return err
		}
	}

	return nil
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

func (item *FeedItem) UpdateFromParser(data *parser.Item) {
	if data.Title != nil {
		item.Title = *data.Title
	}
	if data.Link != nil {
		item.Link = *data.Link
	}
	if data.Guid != nil {
		item.Guid = *data.Guid
	}
	if data.Date != nil {
		item.Date = *data.Date
	}
	if data.Description != nil {
		item.Description = *data.Description
	}

	if data.Media != nil {
		item.MediaUrl = data.Media.Url
		item.MediaLength = data.Media.Length
		item.MediaType = data.Media.Type
	}
}

func (item *FeedItem) NewKey(c appengine.Context, parent *datastore.Key) *datastore.Key {
	id := fmt.Sprintf("%s-%s", utils.Sha1(parent.StringID()), utils.Sha1(item.Link))

	return datastore.NewKey(c, "FeedItem", id, 0, parent)
}
