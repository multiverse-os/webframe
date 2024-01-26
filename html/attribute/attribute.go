package attribute

import (
	"strconv"
	"strings"

	entity "github.com/multiverse-os/webkit/html/entity"
	text "github.com/multiverse-os/webkit/html/text"
)

const defaultRowValue = "3"

type Attributes []Attribute
type Key int

type Attribute struct {
	Key    Key
	Values []string
}

func (self Attribute) Value() string {
	return strings.Join(self.Values, " ")
}

func New(key Key, values ...string) Attribute {
	return Attribute{
		Key:    key,
		Values: values,
	}
}

func (k Key) Values(values ...string) Attribute {
	return Attribute{
		Key:    k,
		Values: values,
	}
}

func (as Attributes) String() (output string) {
	for _, attribute := range as {
		output += attribute.String()
	}
	return output
}

func (a Attribute) String() (value string) {
	// TODO: This is where allow list should filter
	switch a.Key {
	case Rows: // These can only be int
		if len(a.Values) >= 1 {
			// NOTE: Failing Atoi turns rowValue to 0
			intValue, err := strconv.Atoi(a.Values[0])
			if err != nil && (0 < intValue && intValue <= 128) {
				value = defaultRowValue
			} else {
				value = string(intValue)
			}
		} else {
			value = defaultRowValue
		}
		return value
	default:
		value = a.Value()
	}
	return entity.Space.String() + a.Key.String() + entity.EqualSign.String() + text.DoubleQuoteType.Enclose(value).String()
}

// REF: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes
const (
	InvalidKey Key = iota
	Name
	Id
	Class
	Style
	Type
	Action
	Lang // Language subtag: i.e. 'en-US' or 'fr'
	Charset
	Rows
	Cols
	Placeholder
	Content
	Method
	Src
	Href
	Value
	Relative
	AccessKey       // Keyboard key or combo
	AutoCapitalize  // off, none, on, sentences, words, characters
	ContentEditable // Bool
	Hidden
	ContextMenu
	Title
	TabIndex
	InputMode // none, text, decimal, numeric, tel, search, email, url
	Is
	Draggable
	Dir
	ItemId
	ItemType
	ItemScope
	ItemRef
	ItemProp
)

// Aliases
const (
	Rel = Relative
)

func MarshalKey(attributeName string) Key {
	for attribute, allowed := range map[Key]bool{} {
		if allowed && attribute.String() == attributeName {
			return attribute
		}
	}
	return InvalidKey
}

func (k Key) String() string {
	switch k {
	case Class:
		return "class"
	case Name:
		return "name"
	case Id:
		return "id"
	case Style:
		return "style"
	case Lang:
		return "lang"
	case Type:
		return "type"
	case Action:
		return "action"
	case Charset:
		return "charset"
	case Rows:
		return "rows"
	case Cols:
		return "cols"
	case Placeholder:
		return "placeholder"
	case Content:
		return "content"
	case Method:
		return "method"
	case Src:
		return "src"
	case Href:
		return "href"
	case Value:
		return "value"
	default:
		return "invalid"
	}
}

// //////////////////////////////////////////////////////////////////////////////
// Type Attribute Options
type TypeOption int
type TypeOptions map[TypeOption]bool

// REF: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input
const (
	InvalidType TypeOption = iota
	Button
	Checkbox
	Color
	Date
	DateTime
	DateTimeLocal
	Email
	File
	HiddenInput
	Image
	Month
	Number
	Password
	Radio
	Range
	Reset
	Search
	Submit
	Tel
	Text
	Time
	URL
	Week
	//Select
	//Bool
	//Checkboxes - Custom/Abstracted
	//TextArea
)

func MarshalTypeOption(optionName string) TypeOption {
	for option, allowed := range map[TypeOption]bool{} {
		if allowed && option.String() == optionName {
			return option
		}
	}
	return InvalidType
}

func (to TypeOption) String() string {
	switch to {
	case Text:
		return "text"
	case Password:
		return "password"
	case Radio:
		return "radio"
	case Range:
		return "range"
	case Reset:
		return "reset"
	case Search:
		return "search"
	case Tel:
		return "tel"
	case Checkbox:
		return "checkbox"
	case Color:
		return "color"
	case Date:
		return "date"
	case Time:
		return "time"
	case URL:
		return "url"
	case Week:
		return "week"
	case Month:
		return "month"
	case DateTime:
		return "datetime"
	case DateTimeLocal:
		return "datetime"
	case Email:
		return "email"
	case File:
		return "file"
	case Number:
		return "number"
	case HiddenInput:
		return "hidden"
	case Button:
		return "button"
	case Submit:
		return "submit"
	default: // Invalid
		return "invalid"
	}
}
