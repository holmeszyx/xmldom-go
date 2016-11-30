package dom

/*
* Attr implementation
*
* Copyright (c) 2011,2012 Robert Johnstone
* Copyright (c) 2010, Jeff Schiller
 */

import (
	"encoding/xml"
)

type _attr struct {
	_node
	v string // value (for attr)
}

func (a *_attr) NodeType() uint           { return ATTRIBUTE_NODE }
func (a *_attr) NodeName() string         { return a.n.Local }
func (a *_attr) NodeValue() string        { return a.v }
func (a *_attr) PreviousSibling() Node    { return previousSibling(a, a.p.ChildNodes()) }
func (a *_attr) NextSibling() Node        { return nextSibling(a, a.p.ChildNodes()) }
func (a *_attr) AppendChild(n Node) Node  { return n }
func (a *_attr) RemoveChild(n Node) Node  { return n }
func (a *_attr) ParentNode() Node         { return Node(nil) }
func (a *_attr) OwnerDocument() *Document { return ownerDocument(a) }
func (a *_attr) ChildNodes() NodeList     { return NodeList(nil) }
func (a *_attr) Attributes() NamedNodeMap { return NamedNodeMap(nil) }

func newAttr(name string, val string) *_attr {
	a := _attr{_node{nil, nil, xml.Name{"", name}}, val}
	return &a
}
