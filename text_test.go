package dom

import (
	"testing"
)

func TestTextNodeType(t *testing.T) {
	d, _ := ParseStringXml(`<parent>mom</parent>`)
	r := d.DocumentElement()
	txt := r.ChildNodes().Item(0)
	if txt.NodeType() != TEXT_NODE {
		t.Errorf("Did not get the correct node type for a text node")
	}
	if _, ok := txt.(*Text); !ok {
		t.Errorf("Could not convert text node to type *Text")
	}
}

func TestTextNodeName(t *testing.T) {
	d, _ := ParseStringXml(`<parent>mom</parent>`)
	r := d.DocumentElement()
	txt := r.ChildNodes().Item(0)
	if txt.NodeName() != "#text" {
		t.Errorf("Did not get #text for nodeName of a text node")
	}
}

func TestTextNodeValue(t *testing.T) {
	d, _ := ParseStringXml(`<parent>mom</parent>`)
	r := d.DocumentElement()
	txt := r.ChildNodes().Item(0)
	nval := txt.NodeValue()
	if nval != "mom" {
		t.Errorf("Did not get the correct node value for a text node (got %#v)", nval)
	}
}

func TestTextNodeSubstring(t *testing.T) {
	d, _ := ParseStringXml(`<parent>momnonmom</parent>`)
	r := d.DocumentElement()
	txt := r.ChildNodes().Item(0).(*Text)
	nval := txt.SubstringData(3, 3)
	if nval != "non" {
		t.Errorf("Did not get the correct node value for a text node (got %#v)", nval)
	}
}
