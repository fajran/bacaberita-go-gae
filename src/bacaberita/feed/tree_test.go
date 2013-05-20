package feed

import (
	"bytes"
	"testing"
)

func TestError(t *testing.T) {
	str := `<?xml version="1.0" encoding="UTF-8"?><parent></no-parent>`
	r := bytes.NewBuffer([]byte(str))

	_, err := ParseXML(r)
	if err == nil {
		t.Errorf("Parser does not report error")
	}
}

func verifyElementNode(t *testing.T, node *Node, ns string, tag string, size int) {
	if node.Type != ELEMENT {
		t.Errorf("Invalid node type")
		return
	}

	if ns != node.NS {
		t.Errorf("Invalid namespace: %s -> %s", ns, node.NS)
	}
	if tag != node.Tag {
		t.Errorf("Invalid tag name: %s -> %s", tag, node.Tag)
	}
	if size != len(node.Children) {
		t.Errorf("Invalid number of children: %d -> %d",
			size, len(node.Children))
	}
}

func verifyTextNode(t *testing.T, node *Node, value string) {
	if node.Type != TEXT {
		t.Errorf("Invalid node type")
		return
	}

	if value != node.Value {
		t.Errorf("Value mismatch: (len=%d) %w -> (len=%d) %w",
			len(value), []byte(value), len(node.Value), []byte(node.Value))
	}
}

func verifySiblings(t *testing.T, nodes ...*Node) {
	for i := 0; i < len(nodes); i += 1 {
		if i == 0 && nodes[i].Prev != nil {
			t.Errorf("The first node's previous is not nil")
		}
		if i > 0 {
			if nodes[i-1].Next != nodes[i] {
				t.Errorf("Node #%d's next is not Node #%d", i-1, i)
			}
			if nodes[i].Prev != nodes[i-1] {
				t.Errorf("Node #%d's previous is not Node #%d", i, i-1)
			}
		}
		if i == len(nodes)-1 && nodes[i].Next != nil {
			t.Errorf("The last node's next is not nil")
		}
	}
}

func verifyAttribute(t *testing.T, node *Node, ns string, name string, value string) {
	if node.Type != ELEMENT {
		t.Errorf("Invalid node type")
		return
	}

	for _, attr := range node.Attributes {
		if ns == attr.NS && name == attr.Name {
			if value != attr.Value {
				t.Errorf("Attribute's value mismatch, ns=%s name=%s: %s -> %s",
					ns, name, value, attr.Value)
			}
			return
		}
	}

	t.Errorf("Attribute is not found, ns=%s name=%s", ns, name)
}

func TestBasic(t *testing.T) {
	str := `
<?xml version="1.0" encoding="UTF-8"?><parent><child-1>first child</child-1><child-2>second child</child-2></parent>`
	r := bytes.NewBuffer([]byte(str))

	tree, err := ParseXML(r)
	if err != nil {
		t.Errorf("Parser should not report error: %v", err)
		return
	}

	verifyElementNode(t, tree, "", "parent", 2)
	verifyElementNode(t, tree.Children[0], "", "child-1", 1)
	verifyElementNode(t, tree.Children[1], "", "child-2", 1)
	verifySiblings(t, tree.Children[0], tree.Children[1])
}

func TestTextNode(t *testing.T) {
	str := `<?xml version="1.0" encoding="UTF-8"?>
<parent>
  <child-1>first child</child-1>
  <child-2>second child</child-2>
</parent>`
	r := bytes.NewBuffer([]byte(str))

	tree, err := ParseXML(r)
	if err != nil {
		t.Errorf("Parser should not report error: %v", err)
		return
	}

	verifyElementNode(t, tree, "", "parent", 5)
	verifyTextNode(t, tree.Children[0], "\n  ")
	verifyElementNode(t, tree.Children[1], "", "child-1", 1)
	verifyTextNode(t, tree.Children[2], "\n  ")
	verifyElementNode(t, tree.Children[3], "", "child-2", 1)
	verifyTextNode(t, tree.Children[4], "\n")
	verifySiblings(t, tree.Children[0], tree.Children[1], tree.Children[2],
		tree.Children[3], tree.Children[4])
}

func TestNamespace(t *testing.T) {
	str := `<?xml version="1.0" encoding="UTF-8"?>
<parent xmlns="http://example.com/ns"
  xmlns:extra="http://extra">
  <child-1>uses default namespace</child-1>
  <extra:child-2>in the extra namespace</extra:child-2>
  <child-3 xmlns="http://change-default">
    <first>one</first>
    <second>two</second>
  </child-3>
</parent>`
	r := bytes.NewBuffer([]byte(str))

	tree, err := ParseXML(r)
	if err != nil {
		t.Errorf("Parser should not report error: %v", err)
		return
	}

	verifyElementNode(t, tree, "http://example.com/ns", "parent", 7)
	verifyElementNode(t, tree.Children[1], "http://example.com/ns", "child-1", 1)
	verifyElementNode(t, tree.Children[3], "http://extra", "child-2", 1)
	verifyElementNode(t, tree.Children[5], "http://change-default", "child-3", 5)
	verifyElementNode(t, tree.Children[5].Children[1],
		"http://change-default", "first", 1)
	verifyElementNode(t, tree.Children[5].Children[3],
		"http://change-default", "second", 1)
}

func TestNoDeclaration(t *testing.T) {
	str := `<parent><child-1>one</child-1><child-2>two</child-2></parent>`
	r := bytes.NewBuffer([]byte(str))

	tree, err := ParseXML(r)
	if err != nil {
		t.Errorf("Parser should not report error: %v", err)
		return
	}

	verifyElementNode(t, tree, "", "parent", 2)
	verifyElementNode(t, tree.Children[0], "", "child-1", 1)
	verifyElementNode(t, tree.Children[1], "", "child-2", 1)
}

func TestEmptyElement(t *testing.T) {
	str := `<parent><child-1></child-1><child-2/></parent>`
	r := bytes.NewBuffer([]byte(str))

	tree, err := ParseXML(r)
	if err != nil {
		t.Errorf("Parser should not report error: %v", err)
		return
	}

	verifyElementNode(t, tree, "", "parent", 2)
	verifyElementNode(t, tree.Children[0], "", "child-1", 0)
	verifyElementNode(t, tree.Children[1], "", "child-2", 0)
}

func TestAttribute(t *testing.T) {
	str := `<parent version="1.0"></parent>`
	r := bytes.NewBuffer([]byte(str))

	tree, err := ParseXML(r)
	if err != nil {
		t.Errorf("Parser should not report error: %v", err)
		return
	}

	verifyAttribute(t, tree, "", "version", "1.0")
}

func TestAttributeNamespace(t *testing.T) {
	str := `<parent version="2.0" xmlns:extra="http://extra"><child-1 name="one" extra:name="satu"/></parent>`
	r := bytes.NewBuffer([]byte(str))

	tree, err := ParseXML(r)
	if err != nil {
		t.Errorf("Parser should not report error: %v", err)
		return
	}

	verifyAttribute(t, tree, "", "version", "2.0")

	child := tree.Children[0]
	verifyAttribute(t, child, "", "name", "one")
	verifyAttribute(t, child, "http://extra", "name", "satu")
}
