package html

import (
	"strings"

	attribute "github.com/multiverse-os/webframe/html/attribute"
	entity "github.com/multiverse-os/webframe/html/entity"
	text "github.com/multiverse-os/webframe/html/text"
)

const DOCTYPE = `DOCTYPE`

type Tag int
type Tags map[Tag]bool

type Attributes map[attribute.Key][]string

//TODO: Consider all caps naming to avoid any confusion with which tags have multple words and therefore camel-case naming

// REF: https://developer.mozilla.org/en-US/docs/Web/HTML/Element
const (
	A Tag = iota // A
	Abbr
	// Deprecated: Acronym
	Address
	// Deprecated: Applet
	Area
	Article
	Aside
	Audio
	B // B - should be moving to strong
	Base
	// Deprecated: BaseFont
	BDI
	BDO
	// Deprecated: BGSound
	// Deprecated: Big
	// Deprecated: Blink
	BlockQuote
	Body
	Br
	Button
	Canvas // C
	Caption
	// Deprecated: Center
	Cite
	Code
	Col
	ColGroup
	Content
	Data // D
	DataList
	DD
	Del
	Details
	DFN
	Dialog
	// Deprecated: Dir
	Div
	DL
	DT
	Em // E
	Embed
	FieldSet // F
	FigCaption
	Figure
	// Deprecated: Font
	Footer
	Form
	// Deprecated: Frame
	// Deprecated: FrameSet
	H1 // G H
	H2
	H3
	H4
	H5
	H6
	Head
	Header
	// Deprecated: HGroup
	HR
	HTML
	I // I
	IFrame
	Img
	Input
	Ins
	// Deprecated: IsIndex
	KBD // J K
	KeyGen
	Label // L
	Legend
	LI
	Link
	// Deprecated: Listing
	Main // M
	Map
	Mark
	// Deprecated: Marquee
	Menu
	MenuItem
	Meta
	Meter
	Nav // N
	// Deprecated: NoBR
	// Deprecated: NoFrames
	NoScript
	Object // O
	OL
	OptGroup
	Option
	Output
	P // P
	Param
	Picture
	// Deprecated: Plaintext
	Pre
	Progress
	Q  // Q
	RP // R
	RT
	RTC
	Ruby
	S // S
	Samp
	Script
	Section
	Select
	Shadow
	Slot
	Small
	Source
	// Deprecated: Spacer
	Span
	// Deprecated: Strike
	Strong
	Style
	Sub
	Summary
	Sup
	Table // T
	TBody
	TD
	Template
	TextArea
	TFoot
	TH
	THead
	Time
	Title
	TR
	Track
	// Depcreated: TT
	U // U
	UL
	Var // V
	Video
	WBR // W
	// X Y Z
	// Deprecated: Xmp
	InvalidTag
)

func (t Tag) String() string {
	switch t {
	case A:
		return `a`
	case Abbr:
		return `abbr`
	case Area:
		return `area`
	case Article:
		return `article`
	case Aside:
		return `aside`
	case Audio:
		return `audio`
	case B:
		return `b`
	case Base:
		return `base`
	case BDI:
		return `bdi`
	case BDO:
		return `bdo`
	case BlockQuote:
		return `blockquote`
	case Body:
		return `body`
	case Br:
		return `br`
	case Button:
		return `button`
	case Canvas:
		return `canvas`
	case Caption:
		return `caption`
	case Cite:
		return `cite`
	case Code:
		return `code`
	case Col:
		return `col`
	case ColGroup:
		return `colgroup`
	case Content:
		return `content`
	case Data:
		return `data`
	case DataList:
		return `datalist`
	case DD:
		return `dd`
	case Del:
		return `del`
	case Details:
		return `details`
	case DFN:
		return `dfn`
	case Dialog:
		return `dialog`
	case Div:
		return `div`
	case DL:
		return `dl`
	case DT:
		return `dt`
	case Em:
		return `em`
	case Embed:
		return `embed`
	case FieldSet:
		return `fieldset`
	case FigCaption:
		return `figcaption`
	case Figure:
		return `figure`
	case Footer:
		return `footer`
	case Form:
		return `form`
	case H1:
		return `h1`
	case H2:
		return `h2`
	case H3:
		return `h3`
	case H4:
		return `h4`
	case H5:
		return `h5`
	case H6:
		return `h6`
	case Head:
		return `head`
	case Header:
		return `header`
	case HR:
		return `hr`
	case HTML:
		return `html`
	case I:
		return `i`
	case IFrame:
		return `iframe`
	case Img:
		return `img`
	case Input:
		return `input`
	case Ins:
		return `ins`
	case KBD:
		return `kbd`
	case KeyGen:
		return `keygen`
	case Label:
		return `label`
	case LI:
		return `li`
	case Link:
		return `link`
	case Main:
		return `main`
	case Map:
		return `map`
	case Mark:
		return `mark`
	case Menu:
		return `menu`
	case MenuItem:
		return `menuitem`
	case Meta:
		return `meta`
	case Meter:
		return `meter`
	case Nav:
		return `nav`
	case NoScript:
		return `noscript`
	case Object:
		return `object`
	case OL:
		return `ol`
	case OptGroup:
		return `optgroup`
	case Option:
		return `option`
	case Output:
		return `output`
	case P:
		return `p`
	case Param:
		return `param`
	case Picture:
		return `picture`
	case Pre:
		return `pre`
	case Progress:
		return `progress`
	case Q:
		return `q`
	case RP:
		return `rp`
	case RT:
		return `rt`
	case RTC:
		return `rtc`
	case Ruby:
		return `ruby`
	case S:
		return `s`
	case Samp:
		return `samp`
	case Script:
		return `script`
	case Section:
		return `section`
	case Select:
		return `select`
	case Shadow:
		return `shadow`
	case Slot:
		return `slot`
	case Small:
		return `small`
	case Source:
		return `source`
	case Span:
		return `span`
	case Strong:
		return `strong`
	case Style:
		return `style`
	case Sub:
		return `sub`
	case Summary:
		return `summary`
	case Sup:
		return `sup`
	case Table:
		return `table`
	case TBody:
		return `tbody`
	case TD:
		return `td`
	case Template:
		return `template`
	case TextArea:
		return `textarea`
	case TFoot:
		return `tfoot`
	case TH:
		return `th`
	case THead:
		return `thead`
	case Time:
		return `time`
	case Title:
		return `title`
	case TR:
		return `tr`
	case Track:
		return `track`
	case U:
		return `u`
	case UL:
		return `ul`
	case Var:
		return `var`
	case Video:
		return `video`
	case WBR:
		return `wbr`
	default:
		return `undefined`
	}
}

func (t Tag) Element() Element {
	return NewElement(t)
}

func (t Tag) Tags(children ...Element) Element {
	element := NewElement(t)
	element.content = children
	return element
}

func (t Tag) Elements(children ...Element) Element {
	return t.Tags(children...)
}

func (t Tag) Attribute(attr attribute.Attribute) Element {
	var attributes = map[attribute.Key][]string{}
	attributes[attr.Key] = attr.Values
	return Element{
		Tag:        t,
		Attributes: attributes,
	}
}

func (t Tag) Attributes(attrs ...attribute.Attribute) Element {
	var attributes = map[attribute.Key][]string{}
	for _, attr := range attrs {
		attributes[attr.Key] = attr.Values
	}
	return Element{
		Tag:        t,
		Attributes: attributes,
	}
}

// Content Helpers
// /////////////////////////////////////////////////////////////////////////////
func (t Tag) Text(content string) Element {
	element := NewElement(t)
	element.content = string(content)
	return element
}

func (t Tag) ChildElement(element Element) Element {
	newElement := NewElement(t)
	newElement.content = element
	return newElement
}

func (t Tag) ChildElements(elements ...Element) Element {
	element := NewElement(t)
	element.content = elements
	return element
}

func (t Tag) Containing(elements ...Element) Element {
	element := NewElement(t)
	element.content = elements
	return element
}

// String / HTML output
// /////////////////////////////////////////////////////////////////////////////
func (as Attributes) String() (output string) {
	if len(as) > 0 {
		for key, values := range as {
			output += " " + key.String() + `="` + strings.Join(values, " ") + `"`
		}
	}
	return output
}

func (t Tag) EmptyElement(attributes Attributes) string {
	return text.AngleBracketType.Enclose(t.String() + attributes.String() + entity.Space.String() + entity.ForwardSlash.String()).String()
}

func (t Tag) Open(attributes Attributes) string {
	switch t {
	case HTML:
		return text.AngleBracketType.Enclose(entity.Bang.String()+DOCTYPE+entity.Space.String()+strings.ToUpper(t.String())).String() +
			text.AngleBracketType.Enclose(t.String()).String()
	case Meta:
		return t.EmptyElement(attributes)
	default:
		return text.AngleBracketType.Enclose(t.String() + attributes.String()).String()
	}
}

func (t Tag) Close() string {
	return text.AngleBracketType.Enclose(entity.ForwardSlash.String() + t.String()).String()
}

func MarshalTag(tagName string) Tag {
	for tag, allowed := range map[Tag]bool{} {
		if allowed && tag.String() == tagName {
			return tag
		}
	}
	return InvalidTag
}
