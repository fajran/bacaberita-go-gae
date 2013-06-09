package parser

import (
	"time"
)

type Media struct {
	Url    string
	Length int
	Type   string
}

type Item struct {
	Title       *string
	Link        *string
	Guid        *string
	Date        *time.Time
	Description *string
	Media       *Media
}

type Feed struct {
	Title       *string
	Link        *string
	Description *string
	Date        *time.Time
	ImageUrl    *string
	ImageTitle  *string
	ImageLink   *string
	Items       []*Item
}
