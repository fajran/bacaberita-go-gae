package feed

import (
	"errors"
	"io"
	"strconv"
)

func parseImage(feed *Feed, image *Node) error {
	for _, child := range image.ChildElements() {
		if child.Tag == "url" {
			str := child.TextContent()
			feed.ImageUrl = &str

		} else if child.Tag == "title" {
			str := child.TextContent()
			feed.ImageTitle = &str

		} else if child.Tag == "link" {
			str := child.TextContent()
			feed.ImageLink = &str
		}
	}

	return nil
}

func parseMedia(node *Node) (*Media, error) {
	if node.HasAttribute("url") && node.HasAttribute("length") && node.HasAttribute("type") {
		url := node.GetAttribute("url")
		typ := node.GetAttribute("type")
		slength := node.GetAttribute("length")

		length, err := strconv.ParseInt(slength, 10, 0)

		if err == nil {
			media := new(Media)
			media.Url = url
			media.Length = int(length)
			media.Type = typ
			return media, nil
		}

		return nil, err
	}
	return nil, errors.New("Insufficient attributes")
}

func parseItem(node *Node) (Item, error) {
	item := Item{}
	item.Date = nil

	for _, child := range node.ChildElements() {
		if child.Tag == "title" {
			str := child.TextContent()
			item.Title = &str

		} else if child.Tag == "link" {
			str := child.TextContent()
			item.Link = &str

		} else if child.Tag == "guid" {
			str := child.TextContent()
			item.Guid = &str

		} else if child.Tag == "pubDate" {
			date, err := ParseDate(child.TextContent())
			if err == nil {
				item.Date = &date
			}

		} else if child.Tag == "description" {
			str := child.TextContent()
			item.Description = &str

		} else if child.Tag == "enclosure" {
			media, err := parseMedia(child)
			if err == nil {
				item.Media = media
			}
		}
	}

	return item, nil
}

func parseChannel(channel *Node) (*Feed, error) {
	feed := new(Feed)
	feed.Items = make([]Item, 0, 0)
	feed.Date = nil

	for _, child := range channel.ChildElements() {
		if child.Tag == "title" {
			str := child.TextContent()
			feed.Title = &str

		} else if child.Tag == "link" {
			str := child.TextContent()
			feed.Link = &str

		} else if child.Tag == "description" {
			str := child.TextContent()
			feed.Description = &str

		} else if child.Tag == "pubDate" {
			date, err := ParseDate(child.TextContent())
			if err == nil {
				feed.Date = &date
			}

		} else if child.Tag == "image" {
			parseImage(feed, child)

		} else if child.Tag == "item" {
			item, err := parseItem(child)
			if err == nil {
				feed.Items = append(feed.Items, item)
			}
		}
	}

	return feed, nil
}

func ParseRSS(r io.Reader) (*Feed, error) {
	tree, err := ParseXML(r)
	if err != nil {
		return nil, err
	}

	if tree.Tag != "rss" {
		return nil, errors.New("Not an RSS feed")
	}

	// make sure "channel" element is present
	children := tree.ChildElements()
	if len(children) == 0 || children[0].Tag != "channel" {
		return nil, errors.New("Unable to parse feed: channel element is not found")
	}

	feed, err := parseChannel(children[0])
	if err != nil {
		return nil, err
	}

	return feed, nil
}
