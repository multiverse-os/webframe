package maglev

type Environment uint8

const (
	Undefined Environment = iota
	Development
	Production
)

// NOTE: This may seem pointless but environmental checks will occur many times,
// and so making it a int comparsion is important.
// TODO: We need to store environment as uint8 type not as string
func (f Framework) Environment() Environment {
	return MarshalEnvStr(f.Config.Environment)
}

func (f Framework) IsEnvironment(environment Environment) bool {
	return f.Environment() == environment
}

func MarshalEnvStr(environment string) Environment {
	switch environment {
	case Production.String():
		return Production
	case Development.String():
		return Development
	default: // Undefined
		return Undefined
	}
}

func (e Environment) String() string {
	switch e {
	case Production:
		return "production"
	case Development:
		return "development"
	default: // Undefined
		return "undefined"
	}
}
