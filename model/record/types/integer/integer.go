package integer

type Type uint8

const (
	Signed Type = iota
	Unsigned
)

type Size uint8

const (
	Size8BitIs256 Size = iota
	Size16BitIs
	ThirtyTwoBit
	SixtyFourBit
)

// Aliasing Size
const (
	Bit8  = EightBit
	Bit16 = SixteenBit
	Bit32 = ThirtyTwoBit
	Bit64 = SixtyFourBit
)

type Integer struct {
}
