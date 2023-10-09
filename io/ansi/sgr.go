package ansi

import (
	"fmt"
)

func sgr(params ...int) string {
	var mergedParams string
	for index, param := range params {
		if (index + 1) == len(params) {
			mergedParams += fmt.Sprintf("%d", param)
		} else {
			mergedParams += fmt.Sprintf("%d;", param)
		}
	}
	return fmt.Sprintf("\x1b[%sm", mergedParams)
}

func color256(r, g, b uint8) string   { return fmt.Sprintf("%v38;2;%d;%d;%dm", byte(27), r, g, b) }
func color256Bg(r, g, b uint8) string { return fmt.Sprintf("%v48;2;%d;%d;%dm", byte(27), r, g, b) }

var (
	reset           = sgr(0)
	resetBackground = sgr(49)

	black   = sgr(30)
	red     = sgr(31)
	green   = sgr(32)
	yellow  = sgr(33)
	blue    = sgr(34)
	magenta = sgr(35)
	cyan    = sgr(36)
	white   = sgr(37)

	lightBlack   = sgr(90)
	lightRed     = sgr(91)
	lightGreen   = sgr(92)
	lightYellow  = sgr(93)
	lightBlue    = sgr(94)
	lightMagenta = sgr(95)
	lightCyan    = sgr(96)
	lightWhite   = sgr(97)

	blackBackground   = sgr(40)
	redBackground     = sgr(41)
	greenBackground   = sgr(42)
	yellowBackground  = sgr(43)
	blueBackground    = sgr(44)
	magentaBackground = sgr(45)
	cyanBackground    = sgr(46)
	whiteBackground   = sgr(47)

	lightBlackBackground   = sgr(100)
	lightRedBackground     = sgr(101)
	lightGreenBackground   = sgr(102)
	lightYellowBackground  = sgr(103)
	lightBlueBackground    = sgr(104)
	lightMagentaBackground = sgr(105)
	lightCyanBackground    = sgr(106)
	lightWhiteBackground   = sgr(107)
)

var (
	Black   = func(text string) string { return black + text + reset }
	Red     = func(text string) string { return red + text + reset }
	Green   = func(text string) string { return green + text + reset }
	Yellow  = func(text string) string { return yellow + text + reset }
	Blue    = func(text string) string { return blue + text + reset }
	Magenta = func(text string) string { return magenta + text + reset }
	Cyan    = func(text string) string { return cyan + text + reset }
	White   = func(text string) string { return white + text + reset }

	LightBlack   = func(text string) string { return lightBlack + text + reset }
	LightRed     = func(text string) string { return lightRed + text + reset }
	LightGreen   = func(text string) string { return lightGreen + text + reset }
	LightYellow  = func(text string) string { return lightYellow + text + reset }
	LightBlue    = func(text string) string { return lightBlue + text + reset }
	LightMagenta = func(text string) string { return lightMagenta + text + reset }
	LightCyan    = func(text string) string { return lightCyan + text + reset }
	LightWhite   = func(text string) string { return lightWhite + text + reset }

	BlackBg   = func(text string) string { return blackBackground + text + resetBackground }
	RedBg     = func(text string) string { return redBackground + text + resetBackground }
	GreenBg   = func(text string) string { return greenBackground + text + resetBackground }
	YellowBg  = func(text string) string { return yellowBackground + text + resetBackground }
	BlueBg    = func(text string) string { return blueBackground + text + resetBackground }
	MagentaBg = func(text string) string { return magentaBackground + text + resetBackground }
	CyanBg    = func(text string) string { return cyanBackground + text + resetBackground }
	WhiteBg   = func(text string) string { return whiteBackground + text + resetBackground }

	LightBlackBg   = func(text string) string { return lightBlackBackground + text + resetBackground }
	LightRedBg     = func(text string) string { return lightRedBackground + text + resetBackground }
	LightGreenBg   = func(text string) string { return lightGreenBackground + text + resetBackground }
	LightYellowBg  = func(text string) string { return lightYellowBackground + text + resetBackground }
	LightBlueBg    = func(text string) string { return lightBlueBackground + text + resetBackground }
	LightMagentaBg = func(text string) string { return lightMagentaBackground + text + resetBackground }
	LightCyanBg    = func(text string) string { return lightCyanBackground + text + resetBackground }
	LightWhiteBg   = func(text string) string { return lightWhiteBackground + text + resetBackground }
)

var (
	bold            = sgr(1)
	light           = sgr(2)
	italic          = sgr(3)
	underline       = sgr(4)
	slowBlink       = sgr(5)
	fastBlink       = sgr(6)
	inverse         = sgr(7)
	conceal         = sgr(8)
	crossOut        = sgr(9)
	reverse         = sgr(27)
	doubleUnderline = sgr(21)
	frame           = sgr(51)
	encircle        = sgr(52)
	overline        = sgr(53)

	Bold            = func(text string) string { return bold + text + reset }
	Thin            = func(text string) string { return light + text + reset }
	Italic          = func(text string) string { return italic + text + reset }
	Underline       = func(text string) string { return underline + text + reset }
	SlowBlink       = func(text string) string { return slowBlink + text + reset }
	FastBlink       = func(text string) string { return fastBlink + text + reset }
	Inverse         = func(text string) string { return inverse + text + reset }
	Conceal         = func(text string) string { return conceal + text + reset }
	CrossOut        = func(text string) string { return crossOut + text + reset }
	Reverse         = func(text string) string { return reverse + text + reset }
	DoubleUnderline = func(text string) string { return doubleUnderline + text + reset }
	Frame           = func(text string) string { return frame + text + reset }
	Encircle        = func(text string) string { return encircle + text + reset }
	Overline        = func(text string) string { return overline + text + reset }
)

// Aliasing
var Maroon = Red
var Olive = Yellow
var Silver = White
var Lime = LightYellow
var SkyBlue = LightBlue
var Fuchsia = LightMagenta
var LightPurple = LightMagenta
var Aqua = LightCyan
var Purple = Magenta
var Gray = White

var FuchsiaBg = LightMagentaBg
var SkyBlueBg = LightBlueBg
var SilverBg = WhiteBg
var LimeBg = LightYellowBg
var AquaBg = LightCyanBg
var PurpleBg = MagentaBg
var OliveBg = LightYellowBg
var MaroonBg = RedBg
var LightPurpleBg = LightMagentaBg
var GrayBg = WhiteBg

var Light = Thin
var Framed = Frame
var Strikethrough = CrossOut
