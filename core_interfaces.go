package dom

/*
 * Part of the xml/dom Go package
 *
 * Declares the interfaces from DOM Core Level 3
 * http://www.w3.org/TR/DOM-Level-3-Core/
 *
 * Copyright (c) 2011,2012 Robert Johnstone
 * Copyright (c) 2009, Rob Russell
 * Copyright (c) 2010, Jeff Schiller
 */

// TODO: split this out into separate interfaces again eventually

type (
	// DOM3: http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-1950641247
	Node interface {
		AppendChild(Node) Node
		RemoveChild(Node) Node
		InsertBefore(Node, Node) Node
		ReplaceChild(Node, Node) Node
		// attributes
		NodeName() string
		NodeValue() string
		NodeType() uint
		ParentNode() Node
		ChildNodes() NodeList
		Attributes() NamedNodeMap
		HasChildNodes() bool
		FirstChild() Node
		LastChild() Node
		PreviousSibling() Node
		NextSibling() Node

		// internal interface methods needed for implementations (not part of the DOM)
		setParent(Node)
		insertChildAt(Node, uint)
		removeChild(Node)
	}

	// http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-637646024
	Attr interface {
		Node
		OwnerDocument() *Document
	}

	// DOM3: http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-536297177
	NodeList interface {
		Length() uint
		Item(index uint) Node
	}

	// http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-1780488922
	NamedNodeMap interface {
		Length() uint
		Item(index uint) Node
	}
)
