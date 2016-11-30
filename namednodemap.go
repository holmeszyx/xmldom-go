package dom

/*
 * NamedNodeMap implementation
 *
 * Copyright (c) 2011,2012 Robert Johnstone
 * Copyright (c) 2010, Jeff Schiller
 */

// used to return the live attributes of a node
type _attrnamednodemap struct {
	e *Element
}

func (m *_attrnamednodemap) Length() uint {
	return uint(len(m.e.attribs))
}
func (m *_attrnamednodemap) Item(index uint) Node {
	if index >= 0 && index < m.Length() {
		item := m.e.attribs[int(index)]
		return newAttr(item.name, item.value)
	}
	return Node(nil)
}

func newAttrNamedNodeMap(e *Element) *_attrnamednodemap {
	nm := new(_attrnamednodemap)
	nm.e = e
	return nm
}
