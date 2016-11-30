include $(GOROOT)/src/Make.inc

TARG=xml/dom
GOFILES=\
	const.go \
	core_interfaces.go \
	node.go \
	attr.go \
	document.go \
	element.go \
	characterdata.go \
	comment.go \
	text.go \
	nodelists.go \
	namednodemap.go \
	dom.go

include $(GOROOT)/src/Make.pkg
