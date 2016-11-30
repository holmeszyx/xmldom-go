package dom

import (
	"strconv"
	"testing"
)

// Document.nodeName should be #document
// see http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-1841493061
func TestDocumentNodeName(t *testing.T) {
	const str = "<foo></foo>"

	d, err := ParseStringXml(str)
	if err != nil {
		t.Errorf("Error parsing simple XML document (%v).", err)
		if d != nil {
			t.Errorf("Document not nil on return.")
		}
		return
	}
	if d == nil {
		t.Errorf("Document is nil")
	}
	if d.NodeName() != "#document" {
		t.Errorf("Document.nodeName != #document")
	}
	if string(d.ToXml()) != str {
		t.Logf("Received %v instead of %v.", d.ToXml(), str)
		t.Errorf("Error rebuilding XML.")
	}

}

// Document.nodeType should be 9
func TestDocumentNodeType(t *testing.T) {
	d, _ := ParseStringXml("<foo></foo>")
	if d.NodeType() != 9 {
		t.Errorf("Document.nodeType not equal to 9")
	}
	if d.NodeType() != DOCUMENT_NODE {
		t.Errorf("Document.nodeType not equal to DOCUMENT_NODE")
	}
}

func TestDocumentNodeValue(t *testing.T) {
	d, _ := ParseStringXml("<foo></foo>")
	if d.NodeValue() != "" {
		t.Errorf("Document.nodeValue not empty")
	}
}

// Document.documentElement should return an object implementing Element
//func TestDocumentElementIsAnElement(t *testing.T) {
//  d, _ := ParseStringXml("<foo></foo>");
//  n,ok := (d.DocumentElement()).(*dom.Element);
//  if (!ok || n.NodeType() != 1) {
//  	t.Errorf("Document.documentElement did not return an Element");
//  }
//}

func TestDocumentElementNodeName(t *testing.T) {
	d, _ := ParseStringXml("<foo></foo>")
	root := d.DocumentElement()
	if root.NodeName() != "foo" {
		t.Errorf("Element.nodeName not set correctly")
	}
}

func TestDocumentElementTagName(t *testing.T) {
	d, _ := ParseStringXml("<foo></foo>")
	root := d.DocumentElement()
	if root.TagName() != "foo" {
		t.Errorf("Element.tagName not set correctly")
	}
}

// Element.nodeType should be 1
func TestElementNodeType(t *testing.T) {
	test_cases := []string{"<foo></foo>", "<parent>mom</parent>", "<parent><foo></foo></parent>"}

	for _, v := range test_cases {
		d, _ := ParseStringXml(v)
		root := d.DocumentElement()
		if root.NodeType() != 1 {
			t.Errorf("Element.nodeType not equal to 1")
		}
		if root.NodeType() != ELEMENT_NODE {
			t.Errorf("Element.nodeType not equal to 1")
		}
	}
}

func TestElementNodeName(t *testing.T) {
	test_cases := []struct{ text, expected string }{
		{"<foo></foo>", "foo"},
		{"<parent>mom</parent>", "parent"},
		{"<parent><foo></foo></parent>", "parent"},
	}

	for _, v := range test_cases {
		d, _ := ParseStringXml(v.text)
		r := d.DocumentElement()
		if r.NodeName() != v.expected {
			t.Errorf("Did not get '%s' for nodeName of a root node for %s", v.expected, v.text)
		}
	}
}

func TestElementNodeValue(t *testing.T) {
	test_cases := []string{"<foo></foo>", "<parent>mom</parent>", "<parent><foo></foo></parent>"}

	for _, v := range test_cases {
		d, _ := ParseStringXml(v)
		root := d.DocumentElement()
		if root.NodeValue() != "" {
			t.Errorf("Element.nodeValue not empty")
		}
	}
}

func TestElementGetAttribute(t *testing.T) {
	d, _ := ParseStringXml("<foo bar='baz'></foo>")
	root := d.DocumentElement()
	if root.GetAttribute("bar") != "baz" {
		t.Errorf("Element.getAttribute() did not return the attribute value")
	}
	if root.GetAttribute("baz") != "" {
		t.Errorf("Element.getAttribute() returned the attribute value for a non-existant attribute")
	}
}

func TestElementSetAttribute(t *testing.T) {
	d, _ := ParseStringXml("<foo></foo>")
	root := d.DocumentElement()
	root.SetAttribute("bar", "baz")
	if root.GetAttribute("bar") != "baz" {
		t.Errorf("Element.getAttribute() did not return the attribute value")
	}
	if root.GetAttribute("baz") != "" {
		t.Errorf("Element.getAttribute() returned the attribute value for a non-existant attribute")
	}
}

func TestNodeListLength(t *testing.T) {
	d, _ := ParseStringXml(`<foo><bar></bar><baz></baz></foo>`)
	root := d.DocumentElement()
	children := root.ChildNodes()
	l := int(children.Length())
	if l != 2 {
		t.Errorf("NodeList.length did not return the correct number of children (" + strconv.Itoa(l) + " instead of 2)")
	}
}

func TestNodeListItem(t *testing.T) {
	d, _ := ParseStringXml(`<foo><bar></bar><baz></baz></foo>`)
	root := d.DocumentElement()
	children := root.ChildNodes()
	if children.Item(1).NodeName() != "baz" ||
		children.Item(0).NodeName() != "bar" {
		t.Errorf("NodeList.item(i) did not return the correct child")
	}
}

func TestNodeListItemForNull(t *testing.T) {
	d, _ := ParseStringXml(`<foo><bar></bar><baz></baz></foo>`)
	root := d.DocumentElement()
	children := root.ChildNodes()
	if children.Item(2) != nil ||
		children.Item(100000) != nil {
		t.Errorf("NodeList.item(i) did not return nil")
	}
}

func TestNodeParentNode(t *testing.T) {
	d, _ := ParseStringXml(`<foo><bar><baz></baz></bar></foo>`)

	root := d.DocumentElement()
	child := root.ChildNodes().Item(0)
	grandchild := child.ChildNodes().Item(0)

	if d != root.ParentNode().(*Document) ||
		child.ParentNode() != root ||
		grandchild.ParentNode() != child ||
		grandchild.ParentNode().ParentNode() != root {
		t.Errorf("Node.ParentNode() did not return the correct parent")
	}
}

func TestNodeParentNodeOnRoot(t *testing.T) {
	d, _ := ParseStringXml(`<foo></foo>`)

	root := d.DocumentElement()

	if root.ParentNode().(*Document) != d {
		t.Errorf("documentElement.ParentNode() did not return the document")
	}
}

func TestNodeParentNodeOnDocument(t *testing.T) {
	d, _ := ParseStringXml(`<foo></foo>`)
	if d.ParentNode() != nil {
		t.Errorf("document.ParentNode() did not return nil")
	}
}

// the root node of the document is a child node
func TestNodeDocumentChildNodesLength(t *testing.T) {
	d, _ := ParseStringXml(`<foo></foo>`)
	if d.ChildNodes().Length() != 1 {
		t.Errorf("document.ChildNodes().Length() did not return the number of children")
	}
}

func TestNodeDocumentChildNodeIsRoot(t *testing.T) {
	d, _ := ParseStringXml(`<foo></foo>`)
	root := d.DocumentElement()
	if d.ChildNodes().Item(0) != root {
		t.Errorf("document.ChildNodes().Item(0) is not the documentElement")
	}
}

func TestDocumentCreateElement(t *testing.T) {
	d, _ := ParseStringXml(`<foo></foo>`)
	ne := d.CreateElement("child")
	if ne.NodeType() != ELEMENT_NODE {
		t.Errorf("document.CreateNode('child') did not create an element node")
	}
	if ne.NodeName() != "child" {
		t.Errorf("document.CreateNode('child') did not create a <child> Element")
	}
	if ne.OwnerDocument() != Node(d) {
		t.Errorf("document.CreateNode('child') did not create a <child> Element with a proper owner")
	}
}

func TestDocumentCreateTextNode(t *testing.T) {
	d, _ := ParseStringXml(`<foo></foo>`)
	ne := d.CreateTextNode("some text")
	if ne.NodeType() != TEXT_NODE {
		t.Errorf("document.CreateTextNode('some text') did not create a text node")
	}
	if ne.NodeValue() != "some text" {
		t.Errorf("document.CreateTextNode('some text') did not element with the correct text")
	}
	if ne.OwnerDocument() != Node(d) {
		t.Errorf("document.CreateTextNode('some text') did not create a node with a proper owner")
	}
}

func TestAppendChild(t *testing.T) {
	d, _ := ParseStringXml(`<parent></parent>`)
	root := d.DocumentElement()
	ne := d.CreateElement("child")
	appended := root.AppendChild(ne)
	if appended != ne ||
		root.ChildNodes().Length() != 1 ||
		root.ChildNodes().Item(0) != ne {
		t.Errorf("Node.appendChild() did not add the new element")
	}
}

func TestAppendChildParent(t *testing.T) {
	d, _ := ParseStringXml(`<parent></parent>`)
	root := d.DocumentElement()
	ne := d.CreateElement("child")
	root.AppendChild(ne)
	if ne.ParentNode() != Node(root) {
		t.Errorf("Node.appendChild() did not set the parent node")
	}
}

func TestRemoveChild(t *testing.T) {
	d, _ := ParseStringXml(`<parent><child1><grandchild></grandchild></child1><child2></child2></parent>`)

	root := d.DocumentElement()
	child1 := root.ChildNodes().Item(0)
	grandchild := child1.ChildNodes().Item(0)

	child1.RemoveChild(grandchild)

	if child1.ChildNodes().Length() != 0 {
		t.Errorf("Node.removeChild() did not remove child")
	}
}

func TestRemoveChildReturned(t *testing.T) {
	d, _ := ParseStringXml(`<parent><child1><grandchild></grandchild></child1><child2></child2></parent>`)

	root := d.DocumentElement()
	child1 := root.ChildNodes().Item(0)
	grandchild := child1.ChildNodes().Item(0)

	re := child1.RemoveChild(grandchild)

	if grandchild != re {
		t.Errorf("Node.removeChild() did not return the removed node")
	}
}

func TestRemoveChildParentNull(t *testing.T) {
	d, _ := ParseStringXml(`<parent><child></child></parent>`)

	root := d.DocumentElement()
	child := root.ChildNodes().Item(0)

	root.RemoveChild(child)

	if child.ParentNode() != nil {
		t.Errorf("Node.removeChild() did not null out the parentNode")
	}
}

// See http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-184E7107
// "If the newChild is already in the tree, it is first removed."
func TestAppendChildExisting(t *testing.T) {
	d, _ := ParseStringXml(`<parent><child1><grandchild></grandchild></child1><child2></child2></parent>`)

	root := d.DocumentElement()
	child1 := root.ChildNodes().Item(0)
	child2 := root.ChildNodes().Item(1)
	grandchild := child1.ChildNodes().Item(0)

	child2.AppendChild(grandchild)

	if child1.ChildNodes().Length() != 0 ||
		child2.ChildNodes().Length() != 1 {
		t.Errorf("Node.appendChild() did not remove existing child from old parent")
	}
}

func TestAttributesOnDocument(t *testing.T) {
	d, _ := ParseStringXml(`<parent></parent>`)
	if d.Attributes() != nil {
		t.Errorf("Document.attributes() does not return null")
	}
}

func TestAttributesOnElement(t *testing.T) {
	d, _ := ParseStringXml(`<parent attr1="val" attr2="val"><child></child></parent>`)
	r := d.DocumentElement()
	c := r.ChildNodes().Item(0)

	if r.Attributes() == nil || r.Attributes().Length() != 2 ||
		c.Attributes() == nil || c.Attributes().Length() != 0 {
		t.Errorf("Element.attributes().length did not return the proper value")
	}
}

func TestAttrNodeName(t *testing.T) {
	d, _ := ParseStringXml(`<parent attr1="val" attr2="val"/>`)
	r := d.DocumentElement()

	if r.Attributes().Length() != 2 {
		t.Errorf("Element.Attributes().Length() did not return the proper value")
	}

	a := r.Attributes()
	if a.Item(0).NodeName() == "attr1" {
		if a.Item(1).NodeName() != "attr2" {
			t.Errorf("Element.attributes().item(i).nodeName did not return the proper value")
		}
	} else if a.Item(0).NodeName() == "attr2" {
		if a.Item(1).NodeName() != "attr1" {
			t.Errorf("Element.attributes().item(i).nodeName did not return the proper value")
		}
	} else {
		t.Errorf("Element.attributes().item(i).nodeName did not return the proper value")
	}
}

func TestAttrNodeValue(t *testing.T) {
	d, _ := ParseStringXml(`<parent attr1="val1" attr2="val2"/>`)
	r := d.DocumentElement()

	a := r.Attributes()

	if a.Item(0).NodeName() == "attr1" {
		if a.Item(0).NodeValue() != "val1" || a.Item(1).NodeValue() != "val2" {
			t.Errorf("Element.attributes().item(i).nodeValue did not return the proper value")
		}
	} else if a.Item(0).NodeName() == "attr2" {
		if a.Item(0).NodeValue() != "val2" || a.Item(1).NodeValue() != "val1" {
			t.Errorf("Element.attributes().item(i).nodeValue did not return the proper value")
		}
	} else {
		t.Errorf("Element.attributes().item(i).nodeName did not return the proper value")
	}
}

func TestAttributesSetting(t *testing.T) {
	d, _ := ParseStringXml(`<parent attr1="val" attr2="val"><child></child></parent>`)
	r := d.DocumentElement()

	prelen := r.Attributes().Length()

	r.SetAttribute("foo", "bar")

	if prelen != 2 || r.Attributes().Length() != 3 {
		t.Errorf("Element.attributes() not updated when setting a new attribute")
	}
}

func TestToXml(t *testing.T) {
	d1, _ := ParseStringXml(`<parent attr="val">mom<foo/></parent>`)
	s := d1.ToXml()
	d2, _ := ParseStringXml(string(s))
	r2 := d2.DocumentElement()

	if r2.NodeName() != "parent" ||
		r2.GetAttribute("attr") != "val" ||
		r2.ChildNodes().Length() != 2 ||
		r2.ChildNodes().Item(0).NodeValue() != "mom" {
		t.Errorf("ToXml() did not serialize the DOM to text")
	}
}

func TestNodeHasChildNodes(t *testing.T) {
	d, _ := ParseStringXml(`<parent><child/><child>kid</child></parent>`)
	r := d.DocumentElement()
	child1 := r.ChildNodes().Item(0)
	child2 := r.ChildNodes().Item(1)
	text2 := child2.ChildNodes().Item(0)
	if r.HasChildNodes() != true ||
		child1.HasChildNodes() != false ||
		child2.HasChildNodes() != true ||
		text2.HasChildNodes() != false {
		t.Errorf("Node.HasChildNodes() not implemented correctly")
	}
}

func TestChildNodesNodeListLive(t *testing.T) {
	d, _ := ParseStringXml(`<parent></parent>`)
	r := d.DocumentElement()
	children := r.ChildNodes()
	n0 := children.Length()
	c1 := d.CreateElement("child")
	r.AppendChild(c1)
	r.AppendChild(d.CreateElement("child"))
	n2 := children.Length()
	r.RemoveChild(c1)
	n1 := children.Length()
	if n0 != 0 || n1 != 1 || n2 != 2 {
		t.Errorf("NodeList via Node.ChildNodes() was not live")
	}
}

func TestAttributesNamedNodeMapLive(t *testing.T) {
	d, _ := ParseStringXml(`<parent attr1="val1" attr2="val2"></parent>`)
	r := d.DocumentElement()
	attrs := r.Attributes()
	n2 := attrs.Length()
	r.SetAttribute("attr3", "val3")
	n3 := attrs.Length()
	if n2 != 2 || n3 != 3 {
		t.Errorf("NamedoNodeMap via Node.Attributes() was not live")
	}
}

func TestNodeOwnerDocument(t *testing.T) {
	d, _ := ParseStringXml(`<parent><child/><child>kid</child></parent>`)
	r := d.DocumentElement()
	child1 := r.ChildNodes().Item(0).(*Element)
	child2 := r.ChildNodes().Item(1).(*Element)
	text2 := child2.ChildNodes().Item(0).(*Text)
	if r.OwnerDocument() != d ||
		child1.OwnerDocument() != d ||
		child2.OwnerDocument() != d ||
		text2.OwnerDocument() != d {
		t.Errorf("Node.OwnerDocument() did not return the Document object")
	}
}

func TestDocumentGetElementById(t *testing.T) {
	d, _ := ParseStringXml(`<parent id="p"><child/><child id="c"/></parent>`)
	r := d.DocumentElement()
	child2 := r.ChildNodes().Item(1).(*Element)
	p := d.GetElementById("p")
	c := d.GetElementById("c")
	n := d.GetElementById("nothing")
	if p != r ||
		c != child2 ||
		n != nil {
		t.Errorf("Document.GetElementById() not implemented properly")
	}
}

func TestNodeInsertBefore(t *testing.T) {
	d, _ := ParseStringXml(`<parent><child0/><child2/></parent>`)
	r := d.DocumentElement()
	child0 := r.ChildNodes().Item(0)
	child2 := r.ChildNodes().Item(1)
	child1 := d.CreateElement("child1")
	alsoChild1 := r.InsertBefore(child1, child2).(*Element)
	if alsoChild1 != child1 ||
		r.ChildNodes().Length() != 3 ||
		r.ChildNodes().Item(0) != child0 ||
		child0.NodeName() != "child0" ||
		r.ChildNodes().Item(1).(*Element) != child1 ||
		child1.NodeName() != "child1" ||
		r.ChildNodes().Item(2) != child2 ||
		child2.NodeName() != "child2" {
		t.Errorf("Node.InsertBefore() did not insert the new element")
	}
}

func TestNodeReplaceChild(t *testing.T) {
	d, _ := ParseStringXml(`<parent><child0/><child2/></parent>`)
	r := d.DocumentElement()
	children := r.ChildNodes()
	child0 := r.ChildNodes().Item(0)
	child2 := r.ChildNodes().Item(1)
	child1 := d.CreateElement("child1")
	alsoChild2 := r.ReplaceChild(child1, child2)
	if children.Length() != 2 ||
		r.ChildNodes().Item(0) != child0 ||
		alsoChild2 != child2 ||
		r.ChildNodes().Item(1) != Node(child1) {
		t.Errorf("Node.ReplaceChild() not implemented properly")
	}
}

func TestElementGetElementsByTagName(t *testing.T) {
	d, _ := ParseStringXml(
		`<parent id="p"><child>
      <grandchild />
    </child><child>
      <grandchild />
    </child><child/>
  </parent>`)

	r := d.DocumentElement()
	childless := r.ChildNodes().Item(2).(*Element)
	grandchildren := r.GetElementsByTagName("grandchild")
	no_offspring := childless.GetElementsByTagName("grandchild")

	if grandchildren.Length() != 2 {
		t.Errorf("Element.GetElementsByTagName() returned %d children instead of 2", grandchildren.Length())
	} else if no_offspring.Length() != 0 {
		t.Errorf("Element.GetElementsByTagName() returned %d children instead of 0", no_offspring.Length())
	}
}

func TestNodeFirstChild(t *testing.T) {
	d, _ := ParseStringXml(`<parent><child0/><child1/><child2/></parent>`)
	r := d.DocumentElement()
	children := r.ChildNodes()
	child0 := children.Item(0)
	if child0.FirstChild() != nil {
		t.Errorf("Node.firstChild did not return null on an empty node")
	} else if r.FirstChild() != child0 {
		t.Errorf("Node.firstChild did not return the first child")
	}
}

func TestNodeFirstChildAfterInsert(t *testing.T) {
	d, _ := ParseStringXml(`<parent><child1/><child2/></parent>`)
	r := d.DocumentElement()
	children := r.ChildNodes()
	child1 := children.Item(0)
	if r.FirstChild() != child1 {
		t.Errorf("Node.firstChild did not return the first child")
	}

	child0 := d.CreateElement("child0")
	r.InsertBefore(child0, child1)

	if r.FirstChild() != child0 {
		t.Errorf("Node.firstChild did not return the first child after inserting a new element")
	}
}

func TestNodeLastChildAfterAppend(t *testing.T) {
	d, _ := ParseStringXml(`<parent><child0/><child1/></parent>`)
	r := d.DocumentElement()
	children := r.ChildNodes()
	child1 := children.Item(1)
	if r.LastChild() != child1 {
		t.Errorf("Node.lasstChild did not return the last child")
	}

	child2 := d.CreateElement("child2")
	r.AppendChild(child2)

	if r.LastChild() != child2 {
		t.Errorf("Node.lastChild did not return the last child after appending a new element")
	}
}

func TestNodeFirstChildAfterRemove(t *testing.T) {
	d, _ := ParseStringXml(`<parent><child1/><child2/></parent>`)
	r := d.DocumentElement()
	children := r.ChildNodes()
	child1 := children.Item(0)
	child2 := children.Item(1)

	if r.FirstChild() != child1 {
		t.Errorf("Node.firstChild did not return the first child")
	}

	r.RemoveChild(r.FirstChild())

	if r.FirstChild() != child2 {
		t.Errorf("Node.firstChild did not return the first child after removing an element")
	}
}

func TestNodeLastChild(t *testing.T) {
	d, _ := ParseStringXml(`<parent><child0/><child1/><child2/></parent>`)
	r := d.DocumentElement()
	children := r.ChildNodes()
	child2 := children.Item(2)
	if child2.LastChild() != nil {
		t.Errorf("Node.lastChild did not return null on an empty node")
	} else if r.LastChild() != child2 {
		t.Errorf("Node.lastChild did not return the last child")
	}
}
func TestNodePreviousSibling(t *testing.T) {
	d, _ := ParseStringXml(`<parent><child0/><child1/><child2/></parent>`)
	r := d.DocumentElement()
	children := r.ChildNodes()
	child0 := children.Item(0)
	child1 := children.Item(1)
	child2 := children.Item(2)
	if child0.PreviousSibling() != nil {
		t.Errorf("Node.previousSibling did not return null on the first child")
	} else if child1.PreviousSibling() != child0 {
		t.Errorf("Node.previousSibling did not return the previous sibling")
	} else if child2.PreviousSibling().PreviousSibling() != child0 {
		t.Errorf("child2.previousSibling.previousSibling did not return child0")
	}
}
func TestNodeNextSibling(t *testing.T) {
	d, _ := ParseStringXml(`<parent><child0/><child1/><child2/></parent>`)
	r := d.DocumentElement()
	children := r.ChildNodes()
	child0 := children.Item(0)
	child1 := children.Item(1)
	child2 := children.Item(2)
	if child2.NextSibling() != nil {
		t.Errorf("Node.nextSibling did not return null on the last child")
	} else if child1.NextSibling() != child2 {
		t.Errorf("Node.nextSibling did not return the next sibling")
	} else if child0.NextSibling().NextSibling() != child2 {
		t.Errorf("child0.nextSibling.nextSibling did not return child2")
	}
}

func TestNodeNextPrevSibling(t *testing.T) {
	d, _ := ParseStringXml(`<parent><child0/><child1/></parent>`)
	r := d.DocumentElement()
	children := r.ChildNodes()
	child0 := children.Item(0)
	child1 := children.Item(1)
	if child0.NextSibling().PreviousSibling() != child0 {
		t.Errorf("Node.nextSibling.previousSibling did not return itself")
	} else if child1.PreviousSibling().NextSibling() != child1 {
		t.Errorf("Node.previousSibling.nextSibling did not return itself")
	}
}

func TestElementRemoveAttribute(t *testing.T) {
	d, _ := ParseStringXml(`<parent attr="val"/>`)
	r := d.DocumentElement()
	r.RemoveAttribute("attr")
	if r.GetAttribute("attr") != "" {
		t.Errorf("Element.RemoveAttribute() did not remove the attribute, GetAttribute() returns '%s'", r.GetAttribute("attr"))
	}
}

func TestElementHasAttribute(t *testing.T) {
	d, _ := ParseStringXml(`<parent attr="val"/>`)
	r := d.DocumentElement()
	yes := r.HasAttribute("attr")
	r.RemoveAttribute("attr")
	no := r.HasAttribute("attr")
	if yes != true {
		t.Errorf("Element.HasAttribute() returned false when an attribute was present")
	} else if no != false {
		t.Errorf("Element.HasAttribute() returned true after removing an attribute")
	}
}

func TestCommentElementIsParsed(t *testing.T) {
	const str = `<root><!-- comment--></root>`
	d, err := ParseStringXml(str)
	if err != nil || d == nil {
		t.Errorf("Parsing XML containing a comment unsuccessful.")
	}
	if string(d.ToXml()) != str {
		t.Logf("Received %v instead of %v.", d.ToXml(), str)
		t.Errorf("Error rebuilding XML with comment.")
	}
}

func TestCommentElementHasText(t *testing.T) {
	d, _ := ParseStringXml(`<root><!-- comment --></root>`)
	r := d.DocumentElement()
	c := r.ChildNodes()

	if c.Length() != 1 {
		t.Errorf("Error parsing XML comment.")
	}
	if c.Item(0).NodeType() != COMMENT_NODE || c.Item(0).NodeValue() != " comment " {
		t.Errorf("Error parsing XML comment.")
	}
}

func TestToText(t *testing.T) {
	d, _ := ParseStringXml(`<root><child>Some text <span>that</span> is marked-up.</child></root>`)
	r := d.DocumentElement()

	if string(d.ToText(false)) != "Some text that is marked-up." {
		t.Errorf("Error reconstructing text of a marked-up document.")
	}
	if string(r.ChildNodes().Item(0).(*Element).ToText(false)) != "Some text that is marked-up." {
		t.Errorf("Error reconstructing text of a marked-up element.")
	}
}

func TestToTextUnescaped(t *testing.T) {
	d, _ := ParseStringXml(`<root><child>Some text <span>(&amp;)</span> is marked-up.</child></root>`)
	r := d.DocumentElement()

	if string(d.ToText(false)) != "Some text (&) is marked-up." {
		t.Errorf("Error reconstructing text of a marked-up document.")
	}
	if string(r.ChildNodes().Item(0).(*Element).ToText(false)) != "Some text (&) is marked-up." {
		t.Errorf("Error reconstructing text of a marked-up element.")
	}
}

func TestToTextEscaped(t *testing.T) {
	d, _ := ParseStringXml(`<root><child>Some text <span>(&amp;)</span> is marked-up.</child></root>`)
	r := d.DocumentElement()

	if string(d.ToText(true)) != "Some text (&amp;) is marked-up." {
		t.Errorf("Error reconstructing text of a marked-up document.")
	}
	if string(r.ChildNodes().Item(0).(*Element).ToText(true)) != "Some text (&amp;) is marked-up." {
		t.Errorf("Error reconstructing text of a marked-up element.")
	}
}
