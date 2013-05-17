package feed

import (
	"time"
)

type Date string

func (self Date) Parse() (time.Time, error) {
	return time.Parse(time.RFC822, string(self))
}

func (self Date) Format(format string) (string, error) {
	t, err := self.Parse()
	if err != nil {
		return "", err
	}
	return t.Format(format), nil
}

func (self Date) MustFormat(format string) string {
	s, err := self.Format(format)
	if err != nil {
		return err.Error()
	}
	return s
}

type Item struct {
	Title       string `xml:"title" datastore:"title"`
	Link        string `xml:"link" datastore:"link"`
	Guid        string `xml:"guid" datastore:"guid"`
	Date        Date   `xml:"pubDate" datastore:"date"`
	Description string `xml:"description" datastore:"description,noindex"`
}

type Feed struct {
	Title       string `xml:"title" datastore:"title,noindex"`
	Link        string `xml:"link" datastore:"link"`
	Description string `xml:"description" datastore:"description,noindex"`
	ImageUrl    string `xml:"image>url" datastore:"image_url,noindex"`
	ImageTitle  string `xml:"image>title" datastore:"image_title,noindex"`
	ImageLink   string `xml:"image>link" datastore:"image_link,noindex"`
	Items       []Item `xml:"item" datastore:"-"`
}
