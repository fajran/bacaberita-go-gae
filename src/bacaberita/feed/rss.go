package feed

import (
	"errors"
	"io"
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
			if err != nil {
				return item, err
			}
			item.Date = &date

		} else if child.Tag == "description" {
			str := child.TextContent()
			item.Description = &str
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
			if err != nil {
				return nil, err
			}
			feed.Date = &date

		} else if child.Tag == "image" {
			if err := parseImage(feed, child); err != nil {
				return nil, err
			}

		} else if child.Tag == "item" {
			item, err := parseItem(child)
			if err != nil {
				return nil, err
			}
			feed.Items = append(feed.Items, item)
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
