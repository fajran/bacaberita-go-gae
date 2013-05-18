package feed

import (
	"encoding/xml"
	"fmt"
	"io"
)

const (
	ELEMENT = iota
	TEXT
)

type Node struct {
	Type   byte
	Parent *Node
	Next   *Node
	Prev   *Node

	// For Element Node
	NS         string
	Tag        string
	Children   []*Node
	Attributes map[string]string

	// For Text Node
	Value string
}

func (this *Node) String() string {
	var text string
	if this.Type == ELEMENT {
		text = fmt.Sprintf("[Node Type=Element NS=%s Tag=%s]", this.NS, this.Tag)
	} else {
		text = fmt.Sprintf("[Node Type=Text]")
	}
	return text
}

func newElementNode(ns string, tag string) *Node {
	node := new(Node)
	node.Type = ELEMENT
	node.NS = ns
	node.Tag = tag
	node.Children = make([]*Node, 0, 0)
	node.Parent = nil
	return node
}

func newTextNode(value string) *Node {
	node := new(Node)
	node.Type = TEXT
	node.Value = value
	node.Parent = nil
	return node
}

func ParseXML(r io.Reader) (*Node, error) {
	var tree *Node = nil
	var node *Node = nil

	decoder := xml.NewDecoder(r)
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		switch t := token.(type) {
		case xml.StartElement:
			child := newElementNode(t.Name.Space, t.Name.Local)

			if tree == nil {
				tree = child
			}

			child.Parent = node
			if node != nil {
				node.Children = append(node.Children, child)
			}

			node = child

		case xml.EndElement:
			// set siblings
			var prev *Node = nil
			for _, n := range node.Children {
				if prev != nil {
					prev.Next = n
				}
				n.Prev = prev
				prev = n
			}

			node = node.Parent

		case xml.CharData:
			text := newTextNode(string([]byte(t)))
			text.Parent = node

			if node != nil {
				node.Children = append(node.Children, text)
			}
		}
	}

	return tree, nil
}
