package html

import (
	attribute "github.com/multiverse-os/maglev/html/attribute"
)

//var NoAttributes = []Attribute{

// Attribute Helpers
// /////////////////////////////////////////////////////////////////////////////
func (self Element) Attribute(key attribute.Key, values ...string) Element {
	self.Attributes[key] = values
	return self
}

func (self Element) RemoveAttribute(key attribute.Key) Element {
	delete(self.Attributes, key)
	return self
}

// Tag chainable method attribute helpers
///////////////////////////////////////////////////////////////////////////////
// TODO: Should validate input, ensure single item ones do not have extra values
// space separated. Validate they are alphanumeric, no spaces, of a certain
// minimum and maximum length. This will be another layer to avoiding XSS/etc

// If we fail to catch it on data insertion into Golang struct, we will
// hopefully get it on render. Focusing on security at each layer we can provide
// a greater guarantee of reliability
func (self Tag) Name(name string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.Name,
		Values: []string{name},
	})
}

func (self Tag) Id(id string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.Id,
		Values: []string{id},
	})
}

func (self Tag) Class(class ...string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.Class,
		Values: class,
	})
}

func (self Tag) Style(styles ...string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.Style,
		Values: styles,
	})
}

func (self Tag) Type(option attribute.TypeOption) Element {
	// TODO: Ensure that its a valid option
	return self.Attribute(attribute.Attribute{
		Key:    attribute.Type,
		Values: []string{option.String()},
	})
}

func (self Tag) Action(action string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.Action,
		Values: []string{action},
	})
}

func (self Tag) Lang(lang string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.Lang,
		Values: []string{lang},
	})
}

func (self Tag) Charset(charset string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.Charset,
		Values: []string{charset},
	})
}

func (self Tag) Rows(rows string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.Rows,
		Values: []string{rows},
	})
}

func (self Tag) Cols(cols string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.Cols,
		Values: []string{cols},
	})
}

func (self Tag) Placeholder(placeholder string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.Placeholder,
		Values: []string{placeholder},
	})
}

func (self Tag) Content(content string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.Content,
		Values: []string{content},
	})
}

func (self Tag) Method(method string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.Method,
		Values: []string{method},
	})
}

func (self Tag) Src(src string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.Src,
		Values: []string{src},
	})
}

func (self Tag) Href(href string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.Href,
		Values: []string{href},
	})
}

func (self Tag) Value(value string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.Value,
		Values: []string{value},
	})
}

func (self Tag) Relative(relative string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.Relative,
		Values: []string{relative},
	})
}

func (self Tag) AccessKey(accessKey string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.AccessKey,
		Values: []string{accessKey},
	})
}

func (self Tag) AutoCapitalize(autoCapitalize string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.AutoCapitalize,
		Values: []string{autoCapitalize},
	})
}

func (self Tag) ContentEditable(contentEditable string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.ContentEditable,
		Values: []string{contentEditable},
	})
}

func (self Tag) Title(title string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.Title,
		Values: []string{title},
	})
}

func (self Tag) TabIndex(tabIndex string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.TabIndex,
		Values: []string{tabIndex},
	})
}

func (self Tag) InputMode(inputMode string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.InputMode,
		Values: []string{inputMode},
	})
}

func (self Tag) Is(is string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.Is,
		Values: []string{is},
	})
}

func (self Tag) Draggable(draggable string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.Draggable,
		Values: []string{draggable},
	})
}

func (self Tag) Dir(dir string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.Dir,
		Values: []string{dir},
	})
}

func (self Tag) ContextMenu(contextMenu string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.ContextMenu,
		Values: []string{contextMenu},
	})
}

func (self Tag) ItemId(itemId string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.ItemId,
		Values: []string{itemId},
	})
}

func (self Tag) ItemType(itemType string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.ItemType,
		Values: []string{itemType},
	})
}

func (self Tag) ItemScope(itemScope string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.ItemScope,
		Values: []string{itemScope},
	})
}

func (self Tag) ItemRef(itemRef string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.ItemRef,
		Values: []string{itemRef},
	})
}

func (self Tag) ItemProp(itemProp string) Element {
	return self.Attribute(attribute.Attribute{
		Key:    attribute.ItemProp,
		Values: []string{itemProp},
	})
}

// Aliasing
func (self Tag) ID(id string) Element         { return self.Id(id) }
func (self Tag) ItemID(itemId string) Element { return self.ItemId(itemId) }
func (self Tag) Source(src string) Element    { return self.Src(src) }
func (self Tag) Address(src string) Element   { return self.Src(src) }

// Chainable helpers
// /////////////////////////////////////////////////////////////////////////////
func (self Element) Name(name string) Element {
	return self.Attribute(attribute.Name, name)
}

func (self Element) Id(id string) Element {
	return self.Attribute(attribute.Id, id)
}

func (self Element) Class(class ...string) Element {
	return self.Attribute(attribute.Class, class...)
}

func (self Element) Style(style string) Element {
	return self.Attribute(attribute.Style, style)
}

func (self Element) Type(option string) Element {
	return self.Attribute(attribute.Type, option)
}

func (self Element) Action(action string) Element {
	return self.Attribute(attribute.Action, action)
}

func (self Element) Lang(lang string) Element {
	return self.Attribute(attribute.Lang, lang)
}

func (self Element) Charset(charset string) Element {
	return self.Attribute(attribute.Charset, charset)
}

func (self Element) Rows(rows string) Element {
	return self.Attribute(attribute.Rows, rows)
}

func (self Element) Cols(cols string) Element {
	return self.Attribute(attribute.Cols, cols)
}

func (self Element) Placeholder(placeholder string) Element {
	return self.Attribute(attribute.Placeholder, placeholder)
}

func (self Element) Content(content string) Element {
	return self.Attribute(attribute.Content, content)
}

func (self Element) Method(method string) Element {
	return self.Attribute(attribute.Method, method)
}

func (self Element) Src(srcValue string) Element {
	return self.Attribute(attribute.Src, srcValue)
}

func (self Element) Href(hrefValue string) Element {
	return self.Attribute(attribute.Href, hrefValue)
}

func (self Element) Value(value string) Element {
	return self.Attribute(attribute.Value, value)
}

func (self Element) Relative(relative string) Element {
	return self.Attribute(attribute.Relative, relative)
}

func (self Element) Rel(rel string) Element {
	return self.Attribute(attribute.Rel, rel)
}

func (self Element) AccessKey(accessKey string) Element {
	return self.Attribute(attribute.AccessKey, accessKey)
}

func (self Element) AutoCapitalize(autoCapitalize string) Element {
	return self.Attribute(attribute.AutoCapitalize, autoCapitalize)
}

func (self Element) ContentEditable(contentEditable string) Element {
	return self.Attribute(attribute.ContentEditable, contentEditable)
}

func (self Element) Title(title string) Element {
	return self.Attribute(attribute.Title, title)
}

func (self Element) TabIndex(tabIndex string) Element {
	return self.Attribute(attribute.TabIndex, tabIndex)
}

func (self Element) InputMode(inputMode string) Element {
	return self.Attribute(attribute.InputMode, inputMode)
}

func (self Element) Is(is string) Element {
	return self.Attribute(attribute.Is, is)
}

func (self Element) Draggable(draggable string) Element {
	return self.Attribute(attribute.Draggable, draggable)
}

func (self Element) Dir(dir string) Element {
	return self.Attribute(attribute.Dir, dir)
}

func (self Element) ContextMenu(contextMenu string) Element {
	return self.Attribute(attribute.ContextMenu, contextMenu)
}

func (self Element) ItemId(itemId string) Element {
	return self.Attribute(attribute.ItemId, itemId)
}

func (self Element) ItemType(itemType string) Element {
	return self.Attribute(attribute.ItemType, itemType)
}

func (self Element) ItemScope(itemScope string) Element {
	return self.Attribute(attribute.ItemScope, itemScope)
}

func (self Element) ItemRef(itemRef string) Element {
	return self.Attribute(attribute.ItemRef, itemRef)
}

func (self Element) ItemProp(itemProp string) Element {
	return self.Attribute(attribute.ItemProp, itemProp)
}

// Aliasing
func (self Element) ID(id string) Element {
	return self.Id(id)
}

func (self Element) ItemID(itemID string) Element {
	return self.ItemID(itemID)
}
