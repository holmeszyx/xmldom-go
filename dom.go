package dom

/*
 * Implements a very small, very non-compliant subset of the DOM Core Level 3
 * http://www.w3.org/TR/DOM-Level-3-Core/
 *
 * Copyright (c) 2011,2012 Robert Johnstone
 * Copyright (c) 2009, Rob Russell
 * Copyright (c) 2010, Jeff Schiller
 */

// FIXME: we use the empty string "" to denote a 'null' value when the data type
// according to the DOM API is expected to be a string. Perhaps return a pointer to a string?

import (
	"encoding/xml"
	"io"
	"strings"
)

const (
	DEBUG = true
)

type SyntaxError struct {
	Msg string
}

func (se *SyntaxError) Error() string {
	return se.Msg
}

// ====================================

// these are the package-level functions that are the real workhorses
// they only use interface types

func appendChild(p Node, c Node) Node {
	// if the child is already in the tree somewhere,
	// remove it before reparenting
	if c.ParentNode() != nil {
		removeChild(c.ParentNode(), c)
	}
	i := p.ChildNodes().Length()
	p.insertChildAt(c, i)
	c.setParent(p)
	return c
}

func removeChild(p Node, c Node) Node {
	p.removeChild(c)
	c.setParent(nil)
	return c
}

/*
func prevSibling(n Node) Node {
  children := n.ParentNode().ChildNodes()
  //fmt.Println(n)
  for i := children.Length()-1; i > 0; i-- {
    //fmt.Println("  ", i, "  ", children.Item(i))
    if children.Item(i) == n {
      return children.Item(i-1)
    }
  }
  return Node(nil)
}
*/

func ParseString(s string, strict bool, autoClose []string, entity map[string]string) (doc *Document, err error) {
	doc, err = Parse(strings.NewReader(s), strict, autoClose, entity)
	return
}

func ParseStringHtml(s string) (doc *Document, err error) {
	doc, err = Parse(strings.NewReader(s), false, xml.HTMLAutoClose, xml.HTMLEntity)
	return
}

func ParseStringXml(s string) (doc *Document, err error) {
	doc, err = Parse(strings.NewReader(s), true, nil, nil)
	return
}

func ParseHtml(r io.Reader) (doc *Document, err error) {
	doc, err = Parse(r, false, xml.HTMLAutoClose, xml.HTMLEntity)
	return
}

func ParseXml(r io.Reader) (doc *Document, err error) {
	doc, err = Parse(r, true, nil, nil)
	return
}

func Parse(r io.Reader, strict bool, autoClose []string, entity map[string]string) (doc *Document, err error) {
	// Create parser and get first token
	p := xml.NewDecoder(r)
	t, err := p.Token()
	if err != nil {
		return nil, err
	}
	p.Strict = strict
	p.AutoClose = autoClose
	p.Entity = entity

	d := newDoc()
	e := (Node)(nil) // e is the current parent
	for t != nil {
		switch token := t.(type) {
		case xml.StartElement:
			el := newElem(token)
			for ar := range token.Attr {
				el.SetAttribute(token.Attr[ar].Name.Local, token.Attr[ar].Value)
			}
			if e == nil {
				// set doc root
				// this element is a child of e, the last element we found
				e = d.setRoot(el)
			} else {
				// this element is a child of e, the last element we found
				e = e.AppendChild(el)
			}
		case xml.CharData:
			if e == nil {
				// Have not yet seen root element
				// Ignore white space, otherwise throw error
				if strings.TrimSpace(string([]byte(t.(xml.CharData)))) != "" {
					return nil, &SyntaxError{"Text not allowed outside of root element."}
				}
			} else {
				e.AppendChild(newText(token))
			}
		case xml.EndElement:
			e = e.ParentNode()
		case xml.Comment:
			e.AppendChild(newComment(token))

		default:
			// TODO: add handling for other types (text nodes, etc)
		}
		// get the next token
		t, err = p.Token()
	}

	// Make sure that reading stopped on EOF
	if err != io.EOF {
		return nil, err
	}

	// All is good, return the document
	return d, nil
}

// called recursively
func toXml(n Node) []byte {
	s := ""

	switch n.NodeType() {
	case ELEMENT_NODE:
		s += "<" + n.NodeName()

		// iterate over attributes
		for i := uint(0); i < n.Attributes().Length(); i++ {
			a := n.Attributes().Item(i)
			s += " " + a.NodeName() + "=\"" + a.NodeValue() + "\""
		}

		s += ">"

		// iterate over children
		for ch := uint(0); ch < n.ChildNodes().Length(); ch++ {
			s += string(toXml(n.ChildNodes().Item(ch)))
		}

		s += "</" + n.NodeName() + ">"

	case TEXT_NODE:
		s += string(n.(*Text).EscapedBytes())
		break

	case COMMENT_NODE:
		s += "<!--" + string(n.(*Comment).EscapedBytes()) + "-->"
		break

	}
	return []byte(s)
}

// called recursively
func toText(n Node, escape bool) []byte {
	switch n.NodeType() {
	case ELEMENT_NODE:
		// iterate over children
		s := []byte(nil)
		for ch := uint(0); ch < n.ChildNodes().Length(); ch++ {
			s = append(s, toText(n.ChildNodes().Item(ch), escape)...)
		}
		return s

	case TEXT_NODE:
		if escape {
			return n.(*Text).EscapedBytes()
		}
		return n.(*Text).content

	}
	return []byte(nil)
}
