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

type Comment struct {
	CharacterData
}

func (n *Comment) NodeType() uint           { return COMMENT_NODE }
func (n *Comment) NodeName() (s string)     { return "#comment" }
func (n *Comment) NodeValue() (s string)    { return string(n.content) }
func (n *Comment) PreviousSibling() Node    { return previousSibling(Node(n), n.p.ChildNodes()) }
func (n *Comment) NextSibling() Node        { return nextSibling(Node(n), n.p.ChildNodes()) }
func (n *Comment) OwnerDocument() *Document { return ownerDocument(n) }

func newComment(token xml.Comment) *Comment {
	n := new(Comment)
	n.content = token.Copy()
	return n
}
