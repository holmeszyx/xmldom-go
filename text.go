package dom

/*
 * Text node implementation
 *
 * Copyright (c) 2011,2012 Robert Johnstone
 * Copyright (c) 2010, Jeff Schiller
 * Copyright (c) 2009, Rob Russell
 */

import (
	"encoding/xml"
)

type Text struct {
	CharacterData
}

func (n *Text) NodeType() uint           { return TEXT_NODE }
func (n *Text) NodeName() (s string)     { return "#text" }
func (n *Text) NodeValue() (s string)    { return string(n.content) }
func (n *Text) PreviousSibling() Node    { return previousSibling(Node(n), n.p.ChildNodes()) }
func (n *Text) NextSibling() Node        { return nextSibling(Node(n), n.p.ChildNodes()) }
func (n *Text) OwnerDocument() *Document { return ownerDocument(n) }

func newText(token xml.CharData) *Text {
	n := new(Text)
	n.content = token.Copy()
	return n
}
