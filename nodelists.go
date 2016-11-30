package dom

/*
 * NodeList implementations
 *
 * Copyright (c) 2011,2012 Robert Johnstone
 * Copyright (c) 2009, Rob Russell
 * Copyright (c) 2010, Jeff Schiller
 */

// A _childNodelist only stores a reference to its parent node.
// This way the list can be live, each time Length() or Item is
// called, fresh results are returned.
type _childNodelist struct {
	list *[]Node
}

func (nl *_childNodelist) Length() uint {
	return uint(len(*nl.list))
}

func (nl *_childNodelist) Item(index uint) Node {
	if index < uint(len(*nl.list)) {
		// TODO: what if index == nl.p.c.Len() -1 and a node is deleted right now?
		return (*nl.list)[int(index)]
	}
	return nil
}

func newChildNodelist(p *_node) *_childNodelist {
	return &_childNodelist{&p.c}
}

// A _tagNodeList only stores a reference to the node and the tagname
// on which getElementsByTagName() was called so that the list can be
// live.  TODO: Do we really query every time or can we cache the results
// somehow?
type _tagNodeList struct {
	e    *Element
	tag  string
	list []Node
}

func (nl *_tagNodeList) Length() uint {
	return uint(len(nl.list))
}

func (nl *_tagNodeList) Item(index uint) Node {
	if index < uint(len(nl.list)) {
		return (nl.list)[int(index)]
	}
	return nil
}

func addTagNodeList(list *[]Node, e *Element, tag string) {
	for i := 0; i < len(e.c); i++ {
		test := e.c[i]
		if test.NodeType() == ELEMENT_NODE {
			if test.NodeName() == tag {
				*list = append(*list, test)
			}
			addTagNodeList(list, test.(*Element), tag)
		}
	}
}

func newTagNodeList(p *Element, tag string) *_tagNodeList {
	nl := new(_tagNodeList)
	nl.e = p
	nl.tag = tag
	addTagNodeList(&nl.list, p, tag)
	return nl
}
