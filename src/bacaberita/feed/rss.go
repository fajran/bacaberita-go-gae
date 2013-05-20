package feed

import (
	"errors"
	"io"
)

func parseImage(feed *Feed, image *Node) error {
	for _, child := range image.ChildElements() {
		if child.Tag == "url" {
			feed.ImageUrl = child.TextContent()

		} else if child.Tag == "title" {
			feed.ImageTitle = child.TextContent()

		} else if child.Tag == "link" {
			feed.ImageLink = child.TextContent()
		}
	}

	return nil
}

func parseItem(node *Node) (Item, error) {
	item := Item{}
	item.Date = nil

	for _, child := range node.ChildElements() {
		if child.Tag == "title" {
			item.Title = child.TextContent()

		} else if child.Tag == "link" {
			item.Link = child.TextContent()

		} else if child.Tag == "guid" {
			item.Guid = child.TextContent()

		} else if child.Tag == "pubDate" {
			date, err := ParseDate(child.TextContent())
			if err != nil {
				return item, err
			}
			item.Date = &date

		} else if child.Tag == "description" {
			item.Description = child.TextContent()
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
			feed.Title = child.TextContent()

		} else if child.Tag == "link" {
			feed.Link = child.TextContent()

		} else if child.Tag == "description" {
			feed.Description = child.TextContent()

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
