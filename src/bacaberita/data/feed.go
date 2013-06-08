package data

import (
	"time"
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
