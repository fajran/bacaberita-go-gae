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

func TestRSSZeroNineOne(t *testing.T) {
	content, err := ioutil.ReadFile("sampleRss091.xml")
	if err != nil {
		t.Errorf("Fail to read RSS 0.9.1 Sample file")
		return
	}

	r := bytes.NewBuffer(content)
	rss, err := ParseRSS(r)
	if err != nil {
		t.Errorf("Fail to parse RSS: %w", err)
		return
	}

	// Channel
	if rss.Title != "WriteTheWeb" {
		t.Errorf("Invalid channel title")
	}
	if rss.Link != "http://writetheweb.com" {
		t.Errorf("Invalid channel link")
	}
	if rss.Description != "News for web users that write back" {
		t.Errorf("Invalid channel description")
	}
	if rss.Date != nil {
		t.Errorf("Invalid channel time: undefined -> %s", rss.Date)
	}

	// Channel's Image
	if rss.ImageUrl != "http://writetheweb.com/images/mynetscape88.gif" {
		t.Errorf("Invalid image url")
	}
	if rss.ImageTitle != "WriteTheWeb" {
		t.Errorf("Image image title")
	}
	if rss.ImageLink != "http://writetheweb.com" {
		t.Errorf("Image image link")
	}

	// Items
	if len(rss.Items) != 6 {
		t.Errorf("Invalid number of items: 6 -> %s", len(rss.Items))
	}

	// Item #1
	if rss.Items[0].Title != "Giving the world a pluggable Gnutella" {
		t.Errorf("Invalid item #1 title")
	}
	if rss.Items[0].Link != "http://writetheweb.com/read.php?item=24" {
		t.Errorf("Invalid item #1 link")
	}
	if rss.Items[0].Guid != "" {
		t.Errorf("Invalid item #1 GUID")
	}
	if rss.Items[0].Date != nil {
		t.Errorf("Invalid item #1 date: undefined -> %s", rss.Items[0].Date)
	}
	if rss.Items[0].Description != `WorldOS is a framework on which to build programs that work like Freenet or Gnutella -allowing distributed applications using peer-to-peer routing.` {
		t.Errorf("Invalid item #1 description: %s")
	}

	// Item #2
	if rss.Items[1].Title != "Syndication discussions hot up" {
		t.Errorf("Invalid item #2 title")
	}
	if rss.Items[1].Link != "http://writetheweb.com/read.php?item=23" {
		t.Errorf("Invalid item #2 link")
	}
	if rss.Items[1].Guid != "" {
		t.Errorf("Invalid item #2 GUID")
	}
	if rss.Items[1].Date != nil {
		t.Errorf("Invalid item #2 date: undefined -> %s", rss.Items[1].Date)
	}
	if rss.Items[1].Description != `After a period of dormancy, the Syndication mailing list has become active again, with contributions from leaders in traditional media and Web syndication.` {
		t.Errorf("Invalid item #2 description: %s")
	}
}

func TestRSSZeroNineTwo(t *testing.T) {
	content, err := ioutil.ReadFile("sampleRss092.xml")
	if err != nil {
		t.Errorf("Fail to read RSS 0.9.2 Sample file")
		return
	}

	r := bytes.NewBuffer(content)
	rss, err := ParseRSS(r)
	if err != nil {
		t.Errorf("Fail to parse RSS: %w", err)
		return
	}

	// Channel
	if rss.Title != "Dave Winer: Grateful Dead" {
		t.Errorf("Invalid channel title")
	}
	if rss.Link != "http://www.scripting.com/blog/categories/gratefulDead.html" {
		t.Errorf("Invalid channel link")
	}
	if rss.Description != "A high-fidelity Grateful Dead song every day. This is where we're experimenting with enclosures on RSS news items that download when you're not using your computer. If it works (it will) it will be the end of the Click-And-Wait multimedia experience on the Internet. " {
		t.Errorf("Invalid channel description")
	}
	if rss.Date != nil {
		t.Errorf("Invalid channel time: undefined -> %s", rss.Date)
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
	if len(rss.Items) != 22 {
		t.Errorf("Invalid number of items: 4 -> %s", len(rss.Items))
	}

	// Item #1
	if rss.Items[0].Title != "" {
		t.Errorf("Invalid item #1 title")
	}
	if rss.Items[0].Link != "" {
		t.Errorf("Invalid item #1 link")
	}
	if rss.Items[0].Guid != "" {
		t.Errorf("Invalid item #1 GUID")
	}
	if rss.Items[0].Date != nil {
		t.Errorf("Invalid item #1 date: undefined -> %s", rss.Items[0].Date)
	}
	if rss.Items[0].Description != `It's been a few days since I added a song to the Grateful Dead channel. Now that there are all these new Radio users, many of whom are tuned into this channel (it's #16 on the hotlist of upstreaming Radio users, there's no way of knowing how many non-upstreaming users are subscribing, have to do something about this..). Anyway, tonight's song is a live version of Weather Report Suite from Dick's Picks Volume 7. It's wistful music. Of course a beautiful song, oft-quoted here on Scripting News. <i>A little change, the wind and rain.</i>
` {
		t.Errorf("Invalid item #1 description: %s")
	}

	// Item #2
	if rss.Items[1].Title != "" {
		t.Errorf("Invalid item #2 title")
	}
	if rss.Items[1].Link != "" {
		t.Errorf("Invalid item #2 link")
	}
	if rss.Items[1].Guid != "" {
		t.Errorf("Invalid item #2 GUID")
	}
	if rss.Items[1].Date != nil {
		t.Errorf("Invalid item #2 date: undefined -> %s", rss.Items[1].Date)
	}
	if rss.Items[1].Description != `Kevin Drennan started a <a href="http://deadend.editthispage.com/">Grateful Dead Weblog</a>. Hey it's cool, he even has a <a href="http://deadend.editthispage.com/directory/61">directory</a>. <i>A Frontier 7 feature.</i>` {
		t.Errorf("Invalid item #2 description: %s")
	}
}
