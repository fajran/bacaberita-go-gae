package data

import (
	"fmt"
	"time"

	"appengine"
	"appengine/datastore"

	"bacaberita/parser"
	"bacaberita/utils"
)

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
