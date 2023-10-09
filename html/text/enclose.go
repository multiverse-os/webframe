package text

import (
	"strings"

	entity "github.com/multiverse-os/maglev/html/entity"
)

type EnclosureComponent bool

const (
	OpenComponent  EnclosureComponent = true
	CloseComponent EnclosureComponent = false
)

type Enclosure struct {
	enclosureType EnclosureType
	padding       int
	iterations    int
	content       string
}

type EnclosureType int

const (
	SpaceType EnclosureType = iota
	//Brackets
	SquareBracketType
	AngleBracketType
	ParenthesesType
	CurlyBraceType // Are there other braces?
	// Quotes
	SingleQuoteType
	DoubleQuoteType
	GuillemetsType // «» Quote type used in some languages
	// Symbol
	UnderscoreType
	AsterixType
	// Special
	HTMLCommentType
)

// Aliasing
const (
	ParenType = ParenthesesType
)

func (self Enclosure) String() string {
	return strings.Repeat(self.Open(), self.iterations) +
		strings.Repeat(entity.Space.String(), self.padding) +
		self.content +
		strings.Repeat(entity.Space.String(), self.padding) +
		strings.Repeat(self.Close(), self.iterations)
}

func (self EnclosureType) Enclose(content string) Enclosure {
	return Enclosure{
		enclosureType: self,
		content:       content,
		padding:       0,
		iterations:    1,
	}
}

func (self Enclosure) Padding(size int) Enclosure {
	if 0 <= size && size <= 32 {
		self.padding = size
	} else {
		self.padding = 0
	}
	return self
}

func (self Enclosure) Iterations(x int) Enclosure {
	if 0 < x && x <= 32 {
		self.iterations = x
	} else {
		self.iterations = 1
	}
	return self
}

func (self Enclosure) Open() string {
	return strings.Repeat(self.enclosureType.String(OpenComponent), self.iterations)
}

func (self Enclosure) Close() string {
	return strings.Repeat(self.enclosureType.String(CloseComponent), self.iterations)
}

func (self EnclosureType) String(component EnclosureComponent) string {
	switch self {
	case SquareBracketType, AngleBracketType, ParenthesesType,
		CurlyBraceType, GuillemetsType:
		switch self {
		case SquareBracketType:
			if component {
				return entity.OpenSquareBracket.String()
			} else {
				return entity.CloseSquareBracket.String()
			}
		case AngleBracketType:
			if component {
				return entity.OpenAngleBracket.String()
			} else {
				return entity.CloseAngleBracket.String()
			}
		case ParenthesesType:
			if component {
				return entity.OpenParen.String()
			} else {
				return entity.CloseParen.String()
			}
		case CurlyBraceType:
			if component {
				return entity.OpenCurlyBrace.String()
			} else {
				return entity.CloseCurlyBrace.String()
			}
		case GuillemetsType:
			if component {
				return entity.OpenGuillemet.String()
			} else {
				return entity.CloseGuillemet.String()
			}
		default:
			return entity.EmptyString
		}
		// Type with same open and close runes
	case SingleQuoteType:
		return entity.SingleQuote.String()
	case DoubleQuoteType:
		return entity.DoubleQuote.String()
	case AsterixType:
		return entity.Asterix.String()
	case HTMLCommentType:
		if component {
			return entity.OpenAngleBracket.String() + entity.Bang.String() + strings.Repeat(entity.Dash.String(), 2) + entity.Space.String()
		} else {
			return entity.Space.String() + strings.Repeat(entity.Dash.String(), 2) + entity.Bang.String() + entity.CloseAngleBracket.String()
		}
	default: // SpaceType
		return entity.Space.String()
	}
}
