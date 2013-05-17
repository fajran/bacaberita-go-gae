package feed

import (
	"encoding/xml"
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
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Guid        string `xml:"guid"`
	Date        Date   `xml:"pubDate"`
	Description string `xml:"description"`
}

type Feed struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	ImageUrl    string `xml:"image>url"`
	ImageTitle  string `xml:"image>title"`
	ImageLink   string `xml:"image>link"`
	Items       []Item `xml:"item"`
}

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Feed    Feed     `xml:"channel"`
}
