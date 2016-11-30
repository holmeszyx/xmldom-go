package dom

import (
	"strconv"
)

/*
 * Text node implementation
 *
 * Copyright (c) 2011,2012 Robert Johnstone
 * Copyright (c) 2010, Jeff Schiller
 * Copyright (c) 2009, Rob Russell
 */

type CharacterData struct {
	_node
	content []byte
}

func (n *CharacterData) NodeType() uint           { return CDATA_SECTION_NODE }
func (n *CharacterData) NodeName() (s string)     { return "#cdata-section" }
func (n *CharacterData) NodeValue() (s string)    { return string(n.content) }
func (n *CharacterData) PreviousSibling() Node    { return previousSibling(Node(n), n.p.ChildNodes()) }
func (n *CharacterData) NextSibling() Node        { return nextSibling(Node(n), n.p.ChildNodes()) }
func (n *CharacterData) OwnerDocument() *Document { return ownerDocument(n) }

func (n *CharacterData) Data() string {
	return string(n.content)
}

func (n *CharacterData) SetData(s string) {
	n.content = []byte(s)
}

func (n *CharacterData) Length() uint32 {
	return uint32(len(n.content))
}

func (n *CharacterData) SubstringData(offset uint32, count uint32) string {
	// Code does not follow DOM specification
	// Offset and count should be in code points (?)

	if offset+count >= uint32(len(n.content)) {
		// May still throw error if offset is too large
		return string(n.content[offset:])
	}

	// return slice
	return string(n.content[offset : offset+count])
}

func (n *CharacterData) AppendData(data string) {
	n.content = append(n.content, []byte(data)...)
}

func (n *CharacterData) InsertData(offset uint32, data string) {
	if offset == 0 {
		n.content = append([]byte(data), n.content...)
	}

	tmp := append(n.content[0:offset], []byte(data)...)
	n.content = append(tmp, n.content[offset:]...)
}

func (n *CharacterData) DeleteData(offset, count uint32) {
	if offset == 0 {
		if count > uint32(len(n.content)) {
			n.content = nil
			return
		}

		n.content = n.content[count:]
		return
	}

	if offset+count > uint32(len(n.content)) {
		n.content = n.content[0:offset]
		return
	}

	n.content = append(n.content[0:offset], n.content[offset+count:]...)
}

func (n *CharacterData) ReplaceData(offset, count uint32, data string) {
	if offset == 0 {
		n.content = append([]byte(data), n.content[count:]...)
		return
	}

	tmp := append(n.content[0:offset], []byte(data)...)
	n.content = append(tmp, n.content[offset+count:]...)
}

func (n *CharacterData) String() string {
	return string(n.content)
}

func (n *CharacterData) EscapedBytes() []byte {
	runes := []rune(string(n.content))

	output := make([]byte, 0)

	for _, r := range runes {
		switch {
		case r == '<':
			output = append(output, []byte("&lt;")...)
		case r == '>':
			output = append(output, []byte("&gt;")...)
		case r == '&':
			output = append(output, []byte("&amp;")...)
		case r == '\'':
			output = append(output, []byte("&apos;")...)
		case r == '"':
			output = append(output, []byte("&quot;")...)
		case r < 128:
			output = append(output, byte(r))
		default:
			s := "&#" + strconv.Itoa(int(r)) + ";"
			output = append(output, []byte(s)...)
		}
	}

	return output
}
