package feed

import (
	"time"
)

type Item struct {
	Title       string
	Link        string
	Guid        string
	Date        *time.Time
	Description string
}

type Feed struct {
	Title       string
	Link        string
	Description string
	Date        *time.Time
	ImageUrl    string
	ImageTitle  string
	ImageLink   string
	Items       []Item
}
