package dom

/*
 * Part of the xml/dom Go package
 *
 * Tests some of the constants used to identify node types.
 *
 * Copyright (c) 2011, Robert Johnstone
 */

import (
	"testing"
)

func TestConst(t *testing.T) {
	if ELEMENT_NODE != 1 {
		t.Errorf("Value of ELEMENT_NODE is incorrect.")
	}
	if ATTRIBUTE_NODE != 2 {
		t.Errorf("Value of ATTRIBUTE_NODE is incorrect.")
	}
	if TEXT_NODE != 3 {
		t.Errorf("Value of TEXT_NODE is incorrect.")
	}
	if DOCUMENT_NODE != 9 {
		t.Errorf("Value of DOCUMENT_NODE is incorrect.")
	}
}
