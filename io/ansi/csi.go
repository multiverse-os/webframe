package ansi

import (
	"fmt"
	. "math"
)

// NOTE: This is the developing API, track changes here so its easy to see for
// other developers and be present for documentation writing
//
//     CSI(CursorUp.Params())
//     CSI(SGR.Params(11))
//
//

type ControlFunction int

const (
	CursorUp ControlFunction = iota
	CursorDown
	CursorForward
	CursorBack
	CursorNextLine
	CursorPreviousLine
	CursorHorizontalAbsolute
	CursorPosition
	EraseInDisplay
	EraseInLine
	ScrollUp
	ScrollDown
	SaveCursorPosition
	RestoreCursorPosition
	SelectGraphicRendition
)

// Aliasing
const (
	SGR = SelectGraphicRendition
)

// https://bluesock.org/~willg/dev/ansi.html#sequences
// This resource is GREAT

// SET CURSOR
// 	Set cursor position	origin is at (1, 1), not (0, 0).

// TODO: Would be nice to interface{} then handle both string and int input
// params because it shouldn't matter, but in the end int ones will be best
// since we can check against ranges
func (self ControlFunction) Params(params ...int) (output string) {
	// TODO: Should actually verify params are valid-- for example, SGR is limited
	// to specific ranges

	// TODO: hrmm some have ? prefix
	for index, param := range params {
		if index == len(params) {
			output += fmt.Sprintf("%v", param)
		} else {
			output += fmt.Sprintf("%v;", param)
		}
		// TODO: Last one actually doesnt have a ; behind it, we want equivilent to
		// a strings.Join(";", params) which would not end it with a ;, whereas now
		// we would have ; at the end. ⌁
	}
	return output
	switch self {
	case CursorUp:
		return fmt.Sprintf("%vA", output)
	case CursorDown:
		return fmt.Sprintf("%vB", output)
	case CursorForward:
		return fmt.Sprintf("%vC", output)
	case CursorBack:
		return fmt.Sprintf("%vD", output)
	case CursorNextLine:
		return fmt.Sprintf("%vE", output)
	case CursorPreviousLine:
		return fmt.Sprintf("%vF", output)
	case CursorHorizontalAbsolute:
		return fmt.Sprintf("%vH", output)
	case CursorPosition:
		return fmt.Sprintf("%vI", output)
	case EraseInDisplay:
		return fmt.Sprintf("%vJ", output)
	case EraseInLine:
		return fmt.Sprintf("%vK", output)
	case ScrollUp:
		return fmt.Sprintf("%vS", output)
	case ScrollDown:
		return fmt.Sprintf("%vT", output)
	case SaveCursorPosition:
		return fmt.Sprintf("%vs", output)
	case RestoreCursorPosition:
		return fmt.Sprintf("%vu", output)
	default: // SelectGraphicRendition
		return fmt.Sprintf("%vm", output)
	}
}

func (self *Terminal) CSI(a ...interface{}) {
	fmt.Fprintf(&self.stdout, "\x1b[", a...) // byte(27)
}

func (self *Terminal) MoveTo(position Coordinate) { self.CSI("%d;%dH", position.Y+1, position.X+1) }
func (self *Terminal) Clear()                     { self.CSI("2J") }
func (self *Terminal) AltScreen()                 { self.CSI("?1049h") }
func (self *Terminal) MainScreen()                { self.CSI("?1049l") }
func (self *Terminal) ShowCursor()                { self.CSI("?25h") }
func (self *Terminal) HideCursor()                { self.CSI("?25l") }

// NOTE: I like this
func (self *Terminal) Color(code uint8) { self.CSI("%dm", int(code)+30) }

// TODO: Not a great name
func (self *Terminal) PlotChar(position Coordinate, char rune) {
	if position.X < 0 || position.Y < 0 || position.X >= self.width || position.Y >= self.height {
		return
	}
	self.MoveTo(position)
	self.Write(string(char))
}

// Draw vector line using coordinates between -1.0 and +1.0
// It scales to the screen size
func (self *Terminal) DrawLine(start, end Coordinate, char rune) {
	// Center and scale coordinates
	hw := float64(self.width) / 2
	hh := float64(self.height) / 2
	startPos := Coordinate{
		X: int((float64(start.X) * hw) + hw),
		Y: int((float64(start.Y) * hh) + hh),
	}
	endPos := Coordinate{
		X: int((float64(end.X) * hw) + hw),
		Y: int((float64(end.Y) * hh) + hh),
	}
	self.PlotLine(startPos, endPos, char)
}

func (self *Terminal) ClearLine(start, end Coordinate) { self.PlotLine(start, end, ' ') }

// Draw a line using absolute character positions
func (self *Terminal) PlotLine(start, end Coordinate, char rune) {
	x0 := float64(start.X)
	y0 := float64(start.Y)
	x1 := float64(end.X)
	y1 := float64(end.Y)
	px := int(x0)
	py := int(y0)

	dx := Abs(x1 - x0)
	dy := Abs(y1 - y0)

	var sx, sy, err float64
	if x0 < x1 {
		sx = 1
	} else {
		sx = -1
	}
	if y0 < y1 {
		sy = 1
	} else {
		sy = -1
	}

	if dx > dy {
		err = dx / 2.0
	} else {
		err = -dy / 2.0
	}

	// Move to start of line
	self.MoveTo(Coordinate{
		X: int(x0),
		Y: int(y0),
	})
	color := uint8(0)
	for {
		symbol := char
		if symbol != ' ' {
			// Draw dots at the ends
			if (int(x0) == start.X && int(y0) == start.Y) || (int(x0) == end.X && int(y0) == end.Y) {
				color = 5
				self.Color(color)
				symbol = '♦'
			} else if color == 5 {
				color = 3
				self.Color(color)
			}
		}
		if x0 >= 0 && y0 >= 0 && int(x0) < self.width && int(y0) < self.height {
			if px < 0 || py < 0 || int(px) >= self.width || int(py) >= self.height {
				// Prev location was off screen, so we have to move the cursor
				self.MoveTo(Coordinate{
					X: int(x0),
					Y: int(y0),
				})
			}
			self.WriteRune(symbol)
		}

		if x0 == x1 && y0 == y1 {
			// The end
			break
		}

		px = int(x0)
		py = int(y0)

		e2 := err
		if e2 > -dx {
			err -= dy
			x0 += sx
			self.cursor.Right(int(x0) - px - 1)
		} else {
			self.cursor.Left(1)
		}
		if e2 < dy {
			err += dx
			y0 += sy
			self.cursor.Down(int(y0) - py)
		}
	}
}
