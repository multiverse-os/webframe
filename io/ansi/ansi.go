package ansi

import (
	"os"
	"syscall"
	"unsafe"
)

// TODO: Shouldn't this use the below function
func (self *Terminal) TermIOs() syscall.Termios {
	var termios syscall.Termios
	self.SetTermIOs(termios)
	//syscall.Syscall(
	//	syscall.SYS_IOCTL,
	//	os.Stdin.Fd(),
	//	syscall.TCGETS,
	//	uintptr(unsafe.Pointer(&termios)),
	//)
	return termios
}

func (self *Terminal) SetTermIOs(termios syscall.Termios) {
	syscall.Syscall(
		syscall.SYS_IOCTL,
		os.Stdin.Fd(),
		syscall.TCSETS,
		uintptr(unsafe.Pointer(&termios)),
	)
}

func (self *Terminal) RawMode() {
	self.DisableEcho()
	self.DisableCanonical()
	self.DisableFlowControl()
	self.DisableInputProcessing()
	self.DisableCRtoNL()
}

func (self *Terminal) NormalMode() {
	self.EnableCRtoNL()
	self.EnableInputProcessing()
	self.EnableFlowControl()
	self.EnableCanonical()
	self.EnableEcho()
}

func (self *Terminal) SetLFlag(flag uint32) {
	termios := self.TermIOs()
	termios.Lflag |= flag
	self.SetTermIOs(termios)
}

func (self *Terminal) UnsetLFlag(flag uint32) {
	termios := self.TermIOs()
	termios.Lflag &^= flag
	self.SetTermIOs(termios)
}

func (self *Terminal) SetIFlag(flag uint32) {
	termios := self.TermIOs()
	termios.Iflag |= flag
	self.SetTermIOs(termios)
}

func (self *Terminal) UnsetIFlag(flag uint32) {
	termios := self.TermIOs()
	termios.Iflag &^= flag
	self.SetTermIOs(termios)
}

func (self *Terminal) EnableFlowControl()     { self.SetIFlag(syscall.IXON) }
func (self *Terminal) EnableCRtoNL()          { self.SetIFlag(syscall.ICRNL) }
func (self *Terminal) EnableCanonical()       { self.SetLFlag(syscall.ICANON) }
func (self *Terminal) EnableInputProcessing() { self.SetLFlag(syscall.IEXTEN) }
func (self *Terminal) EnableCtrlC()           { self.SetLFlag(syscall.ISIG) }
func (self *Terminal) EnableEcho()            { self.SetLFlag(syscall.ECHO) }

func (self *Terminal) DisableFlowControl()     { self.UnsetIFlag(syscall.IXON) }
func (self *Terminal) DisableCRtoNL()          { self.UnsetIFlag(syscall.ICRNL) }
func (self *Terminal) DisableCanonical()       { self.UnsetLFlag(syscall.ICANON) }
func (self *Terminal) DisableInputProcessing() { self.UnsetLFlag(syscall.IEXTEN) }
func (self *Terminal) DisableCtrlC()           { self.UnsetLFlag(syscall.ISIG) }
func (self *Terminal) DisableEcho()            { self.UnsetLFlag(syscall.ECHO) }

// Enables mouse position tracking
func (self *Terminal) EnableMouse() {
	self.CSI("?1000h") // Enable VT200
	self.CSI("?1002h") // Enable "button" Xterm mouse events
	//self.CSI("?1015h") // Enable urxvt extended mouse positions (for > 223 cells)
	self.CSI("?1006h") // Enable SGR extended mouse positions (for > 223 cells)
}

func (self *Terminal) EnableMouseMove() {
	self.CSI("?1003h") // Enable "any" Xterm mouse events (i.e. mouse move without buttons)
}

// Disables mouse position tracking
func (self *Terminal) DisableMouse() {
	self.CSI("?1006l") // Disable SGR extended mouse positions
	//t.CSI("?1015h")    // Disable urxvt extended mouse positions
	self.CSI("?1002l") // Disable "button" Xterm mouse events
	self.CSI("?1000l") // Disable VT200
}

func (self *Terminal) DisableMouseMove() {
	self.CSI("?1003l") // Disable "any" Xterm mouse events
}

//func sgr(params ...int) string {
//	var mergedParams string
//	for index, param := range params {
//		if (index + 1) == len(params) {
//			mergedParams += fmt.Sprintf("%d", param)
//		} else {
//			mergedParams += fmt.Sprintf("%d;", param)
//		}
//	}
//	return fmt.Sprintf("\x1b[%sm", mergedParams)
//}

//func color3Bit(c uint8) string      { return fmt.Sprintf("%v%dm", byte(27), c+30) }
//func color3BitBg(c uint8) string    { return fmt.Sprintf("%v%dm", byte(27), c+40) }
//func color8Bit(c uint8) string      { return fmt.Sprintf("%v38;5;%dm", byte(27), c) }
//func color8BitBg(c uint8) swtring    { return fmt.Sprintf("%v48;5;%dm", byte(27), c) }
//func color256(r, g, b uint8) string   { return fmt.Sprintf("%v38;2;%d;%d;%dm", byte(27), r, g, b) }
//func color256Bg(r, g, b uint8) string { return fmt.Sprintf("%v48;2;%d;%d;%dm", byte(27), r, g, b) }

//var (
//	reset = sgr(0)
//
//	resetStyle = sgr(11)
//
//	resetColor      = sgr(39)
//	resetBackground = sgr(49)
//)
//
//var (
//	black   = sgr(30)
//	red     = sgr(31)
//	green   = sgr(32)
//	yellow  = sgr(33)
//	blue    = sgr(34)
//	magenta = sgr(35)
//	cyan    = sgr(36)
//	white   = sgr(37)
//
//	lightBlack   = sgr(90)
//	lightRed     = sgr(91)
//	lightGreen   = sgr(92)
//	lightYellow  = sgr(93)
//	lightBlue    = sgr(94)
//	lightMagenta = sgr(95)
//	lightCyan    = sgr(96)
//	lightWhite   = sgr(97)
//
//	blackBackground   = sgr(40)
//	redBackground     = sgr(41)
//	greenBackground   = sgr(42)
//	yellowBackground  = sgr(43)
//	blueBackground    = sgr(44)
//	magentaBackground = sgr(45)
//	cyanBackground    = sgr(46)
//	whiteBackground   = sgr(47)
//
//	lightBlackBackground   = sgr(100)
//	lightRedBackground     = sgr(101)
//	lightGreenBackground   = sgr(102)
//	lightYellowBackground  = sgr(103)
//	lightBlueBackground    = sgr(104)
//	lightMagentaBackground = sgr(105)
//	lightCyanBackground    = sgr(106)
//	lightWhiteBackground   = sgr(107)
//)
//
//var (
//	Black   = func(text string) string { return black + text + reset }
//	Red     = func(text string) string { return red + text + reset }
//	Green   = func(text string) string { return green + text + reset }
//	Yellow  = func(text string) string { return yellow + text + reset }
//	Blue    = func(text string) string { return blue + text + reset }
//	Magenta = func(text string) string { return magenta + text + reset }
//	Cyan    = func(text string) string { return cyan + text + reset }
//	White   = func(text string) string { return white + text + reset }
//
//	LightBlack   = func(text string) string { return lightBlack + text + reset }
//	LightRed     = func(text string) string { return lightRed + text + reset }
//	LightGreen   = func(text string) string { return lightGreen + text + reset }
//	LightYellow  = func(text string) string { return lightYellow + text + reset }
//	LightBlue    = func(text string) string { return lightBlue + text + reset }
//	LightMagenta = func(text string) string { return lightMagenta + text + reset }
//	LightCyan    = func(text string) string { return lightCyan + text + reset }
//	LightWhite   = func(text string) string { return lightWhite + text + reset }
//
//	BlackBg   = func(text string) string { return blackBackground + text + resetBackground }
//	RedBg     = func(text string) string { return redBackground + text + resetBackground }
//	GreenBg   = func(text string) string { return greenBackground + text + resetBackground }
//	YellowBg  = func(text string) string { return yellowBackground + text + resetBackground }
//	BlueBg    = func(text string) string { return blueBackground + text + resetBackground }
//	MagentaBg = func(text string) string { return magentaBackground + text + resetBackground }
//	CyanBg    = func(text string) string { return cyanBackground + text + resetBackground }
//	WhiteBg   = func(text string) string { return whiteBackground + text + resetBackground }
//
//	LightBlackBg   = func(text string) string { return lightBlackBackground + text + resetBackground }
//	LightRedBg     = func(text string) string { return lightRedBackground + text + resetBackground }
//	LightGreenBg   = func(text string) string { return lightGreenBackground + text + resetBackground }
//	LightYellowBg  = func(text string) string { return lightYellowBackground + text + resetBackground }
//	LightBlueBg    = func(text string) string { return lightBlueBackground + text + resetBackground }
//	LightMagentaBg = func(text string) string { return lightMagentaBackground + text + resetBackground }
//	LightCyanBg    = func(text string) string { return lightCyanBackground + text + resetBackground }
//	LightWhiteBg   = func(text string) string { return lightWhiteBackground + text + resetBackground }
//)
//
//// Aliasing
//var (
//	Maroon      = Red
//	Olive       = Yellow
//	Silver      = White
//	Lime        = LightYellow
//	SkyBlue     = LightBlue
//	Fuchsia     = LightMagenta
//	LightPurple = LightMagenta
//	Aqua        = LightCyan
//	Purple      = Magenta
//	Gray        = White
//
//	FuchsiaBg     = LightMagentaBg
//	SkyBlueBg     = LightBlueBg
//	LimeBg        = LightYellowBg
//	AquaBg        = LightCyanBg
//	PurpleBg      = MagentaBg
//	OliveBg       = LightYellowBg
//	MaroonBg      = RedBg
//	LightPurpleBg = LightMagentaBg
//	GrayBg        = WhiteBg
//)
