package html

import (
	attribute "github.com/multiverse-os/webkit/html/attribute"
)

//var NoAttributes = []Attribute{}

// Attribute Helpers
// /////////////////////////////////////////////////////////////////////////////
func (e Element) Attribute(key attribute.Key, values ...string) Element {
	e.Attributes[key] = values
	return e
}

func (e Element) RemoveAttribute(key attribute.Key) Element {
	delete(e.Attributes, key)
	return e
}

// Tag chainable method attribute helpers
///////////////////////////////////////////////////////////////////////////////
// TODO: Should validate input, ensure single item ones do not have extra values
// space separated. Validate they are alphanumeric, no spaces, of a certain
// minimum and maximum length. This will be another layer to avoiding XSS/etc

// If we fail to catch it on data insertion into Golang struct, we will
// hopefully get it on render. Focusing on security at each layer we can provide
// a greater guarantee of reliability
func (t Tag) Name(name string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.Name,
		Values: []string{name},
	})
}

func (t Tag) Id(id string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.Id,
		Values: []string{id},
	})
}

func (t Tag) Class(class ...string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.Class,
		Values: class,
	})
}

func (t Tag) Style(styles ...string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.Style,
		Values: styles,
	})
}

func (t Tag) Type(option attribute.TypeOption) Element {
	// TODO: Ensure that its a valid option
	return t.Attribute(attribute.Attribute{
		Key:    attribute.Type,
		Values: []string{option.String()},
	})
}

func (t Tag) Action(action string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.Action,
		Values: []string{action},
	})
}

func (t Tag) Lang(lang string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.Lang,
		Values: []string{lang},
	})
}

func (t Tag) Charset(charset string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.Charset,
		Values: []string{charset},
	})
}

func (t Tag) Rows(rows string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.Rows,
		Values: []string{rows},
	})
}

func (t Tag) Cols(cols string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.Cols,
		Values: []string{cols},
	})
}

func (t Tag) Placeholder(placeholder string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.Placeholder,
		Values: []string{placeholder},
	})
}

func (t Tag) Content(content string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.Content,
		Values: []string{content},
	})
}

func (t Tag) Method(method string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.Method,
		Values: []string{method},
	})
}

func (t Tag) Src(src string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.Src,
		Values: []string{src},
	})
}

func (t Tag) Href(href string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.Href,
		Values: []string{href},
	})
}

func (t Tag) Value(value string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.Value,
		Values: []string{value},
	})
}

func (t Tag) Relative(relative string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.Relative,
		Values: []string{relative},
	})
}

func (t Tag) AccessKey(accessKey string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.AccessKey,
		Values: []string{accessKey},
	})
}

func (t Tag) AutoCapitalize(autoCapitalize string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.AutoCapitalize,
		Values: []string{autoCapitalize},
	})
}

func (t Tag) ContentEditable(contentEditable string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.ContentEditable,
		Values: []string{contentEditable},
	})
}

func (t Tag) Title(title string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.Title,
		Values: []string{title},
	})
}

func (t Tag) TabIndex(tabIndex string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.TabIndex,
		Values: []string{tabIndex},
	})
}

func (t Tag) InputMode(inputMode string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.InputMode,
		Values: []string{inputMode},
	})
}

func (t Tag) Is(is string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.Is,
		Values: []string{is},
	})
}

func (t Tag) Draggable(draggable string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.Draggable,
		Values: []string{draggable},
	})
}

func (t Tag) Dir(dir string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.Dir,
		Values: []string{dir},
	})
}

func (t Tag) ContextMenu(contextMenu string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.ContextMenu,
		Values: []string{contextMenu},
	})
}

func (t Tag) ItemId(itemId string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.ItemId,
		Values: []string{itemId},
	})
}

func (t Tag) ItemType(itemType string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.ItemType,
		Values: []string{itemType},
	})
}

func (t Tag) ItemScope(itemScope string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.ItemScope,
		Values: []string{itemScope},
	})
}

func (t Tag) ItemRef(itemRef string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.ItemRef,
		Values: []string{itemRef},
	})
}

func (t Tag) ItemProp(itemProp string) Element {
	return t.Attribute(attribute.Attribute{
		Key:    attribute.ItemProp,
		Values: []string{itemProp},
	})
}

// Aliasing
func (t Tag) ID(id string) Element         { return t.Id(id) }
func (t Tag) ItemID(itemId string) Element { return t.ItemId(itemId) }
func (t Tag) Source(src string) Element    { return t.Src(src) }
func (t Tag) Address(src string) Element   { return t.Src(src) }

// Chainable helpers
// /////////////////////////////////////////////////////////////////////////////
func (e Element) Name(name string) Element {
	return e.Attribute(attribute.Name, name)
}

func (e Element) Id(id string) Element {
	return e.Attribute(attribute.Id, id)
}

func (e Element) Class(class ...string) Element {
	return e.Attribute(attribute.Class, class...)
}

func (e Element) Style(style string) Element {
	return e.Attribute(attribute.Style, style)
}

func (e Element) Type(option string) Element {
	return e.Attribute(attribute.Type, option)
}

func (e Element) Action(action string) Element {
	return e.Attribute(attribute.Action, action)
}

func (e Element) Lang(lang string) Element {
	return e.Attribute(attribute.Lang, lang)
}

func (e Element) Charset(charset string) Element {
	return e.Attribute(attribute.Charset, charset)
}

func (e Element) Rows(rows string) Element {
	return e.Attribute(attribute.Rows, rows)
}

func (e Element) Cols(cols string) Element {
	return e.Attribute(attribute.Cols, cols)
}

func (e Element) Placeholder(placeholder string) Element {
	return e.Attribute(attribute.Placeholder, placeholder)
}

func (e Element) Content(content string) Element {
	return e.Attribute(attribute.Content, content)
}

func (e Element) Method(method string) Element {
	return e.Attribute(attribute.Method, method)
}

func (e Element) Src(srcValue string) Element {
	return e.Attribute(attribute.Src, srcValue)
}

func (e Element) Href(hrefValue string) Element {
	return e.Attribute(attribute.Href, hrefValue)
}

func (e Element) Value(value string) Element {
	return e.Attribute(attribute.Value, value)
}

func (e Element) Relative(relative string) Element {
	return e.Attribute(attribute.Relative, relative)
}

func (e Element) Rel(rel string) Element {
	return e.Attribute(attribute.Rel, rel)
}

func (e Element) AccessKey(accessKey string) Element {
	return e.Attribute(attribute.AccessKey, accessKey)
}

func (e Element) AutoCapitalize(autoCapitalize string) Element {
	return e.Attribute(attribute.AutoCapitalize, autoCapitalize)
}

func (e Element) ContentEditable(contentEditable string) Element {
	return e.Attribute(attribute.ContentEditable, contentEditable)
}

func (e Element) Title(title string) Element {
	return e.Attribute(attribute.Title, title)
}

func (e Element) TabIndex(tabIndex string) Element {
	return e.Attribute(attribute.TabIndex, tabIndex)
}

func (e Element) InputMode(inputMode string) Element {
	return e.Attribute(attribute.InputMode, inputMode)
}

func (e Element) Is(is string) Element {
	return e.Attribute(attribute.Is, is)
}

func (e Element) Draggable(draggable string) Element {
	return e.Attribute(attribute.Draggable, draggable)
}

func (e Element) Dir(dir string) Element {
	return e.Attribute(attribute.Dir, dir)
}

func (e Element) ContextMenu(contextMenu string) Element {
	return e.Attribute(attribute.ContextMenu, contextMenu)
}

func (e Element) ItemId(itemId string) Element {
	return e.Attribute(attribute.ItemId, itemId)
}

func (e Element) ItemType(itemType string) Element {
	return e.Attribute(attribute.ItemType, itemType)
}

func (e Element) ItemScope(itemScope string) Element {
	return e.Attribute(attribute.ItemScope, itemScope)
}

func (e Element) ItemRef(itemRef string) Element {
	return e.Attribute(attribute.ItemRef, itemRef)
}

func (e Element) ItemProp(itemProp string) Element {
	return e.Attribute(attribute.ItemProp, itemProp)
}

// Aliasing
func (e Element) ID(id string) Element {
	return e.Id(id)
}

func (e Element) ItemID(itemID string) Element {
	return e.ItemID(itemID)
}
