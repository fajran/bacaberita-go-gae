package feed

import (
	"bytes"
	"io/ioutil"
	"testing"
	"time"
)

func equalDate(t1 time.Time, t2 time.Time) bool {
	// TODO compare date's time zone?
	return t1.Year() == t2.Year() &&
		t1.Month() == t2.Month() &&
		t1.Day() == t2.Day() &&
		t1.Hour() == t2.Hour() &&
		t1.Minute() == t2.Minute() &&
		t1.Second() == t2.Second()
}

func TestRSSTwoZero(t *testing.T) {
	content, err := ioutil.ReadFile("rss2sample.xml")
	if err != nil {
		t.Errorf("Fail to read RSS 2.0 Sample file")
		return
	}

	r := bytes.NewBuffer(content)
	rss, err := ParseRSS(r)
	if err != nil {
		t.Errorf("Fail to parse RSS: %w", err)
		return
	}

	// Channel
	tt := time.Date(2003, time.June, 10, 4, 0, 0, 0, time.UTC)
	if rss.Title != "Liftoff News" {
		t.Errorf("Invalid channel title")
	}
	if rss.Link != "http://liftoff.msfc.nasa.gov/" {
		t.Errorf("Invalid channel link")
	}
	if rss.Description != "Liftoff to Space Exploration." {
		t.Errorf("Invalid channel description")
	}
	if !equalDate(*rss.Date, tt) {
		t.Errorf("Invalid channel time: %s -> %s", tt, rss.Date)
	}

	// Channel's Image
	if rss.ImageUrl != "" {
		t.Errorf("Image URL is not empty")
	}
	if rss.ImageTitle != "" {
		t.Errorf("Image title is not empty")
	}
	if rss.ImageLink != "" {
		t.Errorf("Image link is not empty")
	}

	// Items
	if len(rss.Items) != 4 {
		t.Errorf("Invalid number of items: 4 -> %s", len(rss.Items))
	}

	// Item #1
	tt = time.Date(2003, time.June, 3, 9, 39, 21, 0, time.UTC)
	if rss.Items[0].Title != "Star City" {
		t.Errorf("Invalid item #1 title")
	}
	if rss.Items[0].Link != "http://liftoff.msfc.nasa.gov/news/2003/news-starcity.asp" {
		t.Errorf("Invalid item #1 link")
	}
	if rss.Items[0].Guid != "http://liftoff.msfc.nasa.gov/2003/06/03.html#item573" {
		t.Errorf("Invalid item #1 GUID")
	}
	if !equalDate(*rss.Items[0].Date, tt) {
		t.Errorf("Invalid item #1 date: %s -> %s", tt, rss.Items[0].Date)
	}
	if rss.Items[0].Description != `How do Americans get ready to work with Russians aboard the International Space Station? They take a crash course in culture, language and protocol at Russia's <a href="http://howe.iki.rssi.ru/GCTC/gctc_e.htm">Star City</a>.` {
		t.Errorf("Invalid item #1 description: %s")
	}

	// Item #2
	tt = time.Date(2003, time.May, 30, 11, 6, 42, 0, time.UTC)
	if rss.Items[1].Title != "" {
		t.Errorf("Invalid item #2 title")
	}
	if rss.Items[1].Link != "" {
		t.Errorf("Invalid item #2 link")
	}
	if rss.Items[1].Guid != "http://liftoff.msfc.nasa.gov/2003/05/30.html#item572" {
		t.Errorf("Invalid item #2 GUID")
	}
	if !equalDate(*rss.Items[1].Date, tt) {
		t.Errorf("Invalid item #2 date: %s -> %s", tt, rss.Items[1].Date)
	}
	if rss.Items[1].Description != `Sky watchers in Europe, Asia, and parts of Alaska and Canada will experience a <a href="http://science.nasa.gov/headlines/y2003/30may_solareclipse.htm">partial eclipse of the Sun</a> on Saturday, May 31st.` {
		t.Errorf("Invalid item #2 description: %s")
	}
}
