package text

// type Runes string
type InputString string

const (
	alpha   = `abcdefghijklmnopqrstuvwxyz`
	numeric = `0123456789`
)

// TODO Use above runesets to provide basic Numericality and IsAlpha and
// IsAlphanumeric

// TODO: Review multiverse-os validation package for some ideas and maybe just
// use that.

// TODO: This will likely be moved into an active_record type package under
// db for working with our databases in a smart way

// but working on forms has me wanting to outline bits of it

type ValidationType int

// TODO: We are initializing a lot of const maybe we would be better off using
// uint8 or something for our type switches but overall they still do not
// use that much memory, and we dont use global variables otherwise so it should
// not be too memory ehavy but we can benchmark soon as we get our testing
// going
const (
	LengthValidation       ValidationType = iota
	NumericValidation                     // options include equals, lessThan, greaterThan, notEqual
	AlphaValidation                       // options include specifying specific runes/characters
	AlphanumericValidation                // options include specifying allowed chars/nums/runes
	RegexValidation                       // specify the regex to match against

	EmailValidation // only real validation is done via checking smtp/or emailing
	URLValidation   // or maybe local/remote should be options of url validation
	//LocalURLValidation
	//RemoteURLValidation

)

func (self ValidationType) String() string {
	switch self {
	case LengthValidation:
		return "length"
	case NumericValidation:
		return "numeric"
	case AlphaValidation:
		return "alpha"
	case AlphanumericValidation:
		return "alphanumeric"
	case RegexValidation:
		return "regex"
	case EmailValidation: // type of regex unless we get real
		return "email"
	case URLValidation:
		return "url"
	default: // InvalidType
		return ""
	}
}

// Thse will be just bools, then error messages can be handled in a single
// function. That checks which validations need to be preformed, checked
// with these, then errors applied.
func (self InputString) IsBetween(minimum, maximum int) bool {
	return (minimum <= len(string(self)) && len(string(self)) <= maximum)
}
