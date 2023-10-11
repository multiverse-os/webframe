package color

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

//type RGB3Bit uint8
//type RGB8Bit uint8
//
//func rgbTo3Bit(r, g, b uint8) RGB3Bit {
//	return RGB3Bit((r >> 7) | ((g >> 6) & 0x2) | ((b >> 5) & 0x4))
//}
//
//func rgbTo8Bit(r, g, b uint8) RGB8Bit {
//	return RGB8Bit(16 + 36*(r/43) + 6*(g/43) + b/43)
//}

//func color3Bit(c RGB3Bit) string      { return fmt.Sprintf("%v%dm", byte(27), c+30) }
//func color3BitBg(c RGB3Bit) string    { return fmt.Sprintf("%v%dm", byte(27), c+40) }
//func color8Bit(c RGB8Bit) string      { return fmt.Sprintf("%v38;5;%dm", byte(27), c) }
//func color8BitBg(c RGB8Bit) string    { return fmt.Sprintf("%v48;5;%dm", byte(27), c) }
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
var LimeBg = LightYellowBg
var AquaBg = LightCyanBg
var PurpleBg = MagentaBg
var OliveBg = LightYellowBg
var MaroonBg = RedBg
var LightPurpleBg = LightMagentaBg
var GrayBg = WhiteBg
