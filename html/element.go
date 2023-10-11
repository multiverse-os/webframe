package html

import (
	"errors"
	"strings"

	attribute "github.com/multiverse-os/maglev/html/attribute"
	entity "github.com/multiverse-os/maglev/html/entity"
)

// [SECURITY] TODO: Every link that goes outside should have target="_no_rel" to prevent certian types of XSS

// Anatomy Of HTML Element
// ///////////////////////////////////////////////////////////////////////////////////////////////
// [Mozilla HTML](https://developer.mozilla.org/en-US/docs/Learn/HTML/Introduction_to_HTML/Getting_started)
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element
//
// [Opening Tag]                        [Closing Tag]
//
//	     \                                   /
//	     <p> Paragraph text between tags </p>
//	         |__________________________|
//	                     |
//	                  [content]
//	|____________________________________________|
//	                     |
//	    [All of it together is an HTML element]
type Elements []Element

type Element struct {
	Tag        Tag
	Block      bool
	Attributes map[attribute.Key][]string
	content    interface{}
}

func NewElement(tag Tag) Element {
	return Element{
		Tag:        tag,
		Attributes: make(map[attribute.Key][]string),
	}
}

// HTML parsing into Go structs (incomplete)
// /////////////////////////////////////////////////////////////////////////////
func ParseElement(html string) (element Element, err error) {
	if len(html) == 0 {
		return element, errors.New("error: empty input string")
	}
	element.Tag = MarshalTag(strings.Split(strings.Split(html[1:], entity.GreaterThanSign.String())[0], entity.Space.String())[0])
	return element, nil
}

// Content Helpers
// /////////////////////////////////////////////////////////////////////////////
func (e Element) Text(content string) Element {
	e.content = string(content)
	return e
}

func (e Element) ChildElement(element Element) Element {
	e.content = element
	return e
}

func (e Element) ChildElements(elements ...Element) Element {
	e.content = elements
	return e
}

func (e Element) Containing(elements ...Element) Element {
	e.content = elements
	return e
}

// String output
// /////////////////////////////////////////////////////////////////////////////
func (e Element) String() (output string) {
	output += e.Tag.Open(e.Attributes)
	switch content := e.content.(type) {
	case string:
		if len(content) > 0 {
			output += content
		}
	case []Element:
		for _, element := range content {
			output += element.String()
		}
	case Element:
		output += content.String()
	}
	if e.Tag != Meta {
		output += e.Tag.Close()
	}
	return output
}

func (e Element) Bytes() []byte {
	return []byte(e.String())
}
