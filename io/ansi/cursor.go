package ansi

// TODO: This hasn't been merged in yet, this overlaps with cursor.go

//var dsrPattern = regexp.MustCompile(`\x1b\[(\d+);(\d+)R$`)

type Cursor struct {
	Terminal *Terminal
	Position *Coordinate
	Symbol   string
	Blinking bool
	Hidden   bool
}

// TODO: These need to update the Cursor.Position after changing it so we can
// have that local cache for simple checking.
// Relative Movement
// CSI, or "Control Sequence Introducer" = "ESC+["
//func (self Cursor) Up(row uint)          { self.Terminal.Print(string.CSI("A", row)) }
//func (self Cursor) Down(row uint)        { self.Terminal.Print(CSI("B", row)) }
func (self Cursor) Forward(column uint)  { self.Terminal.CSI("C", column) }
func (self Cursor) Backward(column uint) { self.Terminal.CSI("D", column) }

func (self Cursor) NextLine(row uint) { self.Terminal.CSI("E", row) }
func (self Cursor) PrevLine(row uint) { self.Terminal.CSI(row, "F") }

func (self Cursor) MoveTo(row, column uint) { self.Terminal.CSI("G", row, column) }
func (self Cursor) To(row, column uint)     { self.Terminal.CSI("H", row, column) }
func (self Cursor) Home()                   { self.Terminal.CSI("H") }

func (self Cursor) IsAtStartOfLine() bool { return self.Position.X == 0 }
func (self Cursor) IsAtEndOfLine() bool   { return self.Position.X == self.Terminal.Width() }

// TODO: These should be method of terminal not cursor
func (self Cursor) ClearScreenPartialForward()  { self.Terminal.DSR("J") }
func (self Cursor) ClearScreenPartialBackward() { self.Terminal.DSR("J", 1) }
func (self Cursor) ClearScreen()                { self.Terminal.DSR("J", 2) }

func (self Cursor) ClearLinePartialForward()  { self.Terminal.DSR("K") }
func (self Cursor) ClearLinePartialBackward() { self.Terminal.DSR("K", 1) }
func (self Cursor) ClearLine()                { self.Terminal.DSR("K", 2) }

// TODO: Need a better way to handle initializing
func (self Cursor) HorizontalAbsolute(x int) { self.Terminal.DSR("[", x, "G") }
func (self Cursor) Show()                    { self.Terminal.DSR("?25h") }
func (self Cursor) Hide()                    { self.Terminal.DSR("?25l") }
func (self Cursor) Move(x int, y int)        { self.Terminal.DSR("[", x, ";", y, "f") }
func (self Cursor) Save()                    { self.Terminal.DSR("[7") }
func (self Cursor) Restore()                 { self.Terminal.DSR("[8") }

//func (self Cursor) MoveNextLine() {
//	if coordinate.Y == terminalSize.Y {
//		fmt.Fprintln(self.Terminal.Print())
//	}
//	self.NextLine(1)
//}

// ANSI escape sequence for DSR - Device Status Report
// https://en.wikipedia.org/wiki/ANSI_escape_code#CSI_sequences
//func (self Cursor) Location() (Coordinate, error) {
//	self.Terminal.Print("[6n")
//	// There may be input in Stdin prior to CursorLocation so make sure we don't
//	// drop those bytes.
//	var location []int
//	var match string
//	for location == nil {
//		// Reports the cursor position (CPR) to the application as (as though typed at
//		// the keyboard) ESC[n;mR, where n is the row and m is the column.
//		reader := bufio.NewReader(self.In)
//		text, err := reader.ReadSlice('R')
//		if err != nil {
//			return nil, err
//		}
//
//		location = dsrPattern.FindStringIndex(string(text))
//		if location == nil {
//			// Stdin contains R that doesn't match DSR.
//			buffer.Write(text)
//		} else {
//			buffer.Write(text[:location[0]])
//			match = string(text[location[0]:location[1]])
//		}
//	}
//
//	matches := dsrPattern.FindStringSubmatch(string(match))
//	if len(matches) != 3 {
//		return nil, fmt.Errorf("incorrect number of matches: %d", len(matches))
//	}
//
//	column, err := strconv.Atoi(matches[2])
//	if err != nil {
//		return nil, err
//	}
//
//	row, err := strconv.Atoi(matches[1])
//	if err != nil {
//		return nil, err
//	}
//	self.Position = Coordinate{
//		X: row,
//		Y: column,
//	}
//	return self.Position, nil
//}

func (self *Cursor) Left(amount int) {
	if amount == 1 {
		//self.CSI("D")
		self.Terminal.CSI(CursorBack.Params())
	} else if amount > 1 {
		//self.CSI("%dD", amount)
		self.Terminal.CSI(CursorBack.Params(amount))
	} else if amount < 0 {
		self.Right(-amount)
	}
}

func (self *Cursor) Right(amount int) {
	if amount == 1 {
		//self.CSI("C")
		self.Terminal.CSI(CursorForward.Params())
	} else if amount > 1 {
		//self.CSI("%dC", amount)
		self.Terminal.CSI(CursorForward.Params(amount))
	} else if amount < 0 {
		self.Left(-amount)
	}
}

func (self *Cursor) Down(amount int) {
	if amount == 1 {
		//self.CSI("B")
		self.Terminal.CSI(CursorDown.Params())
	} else if amount > 1 {
		//self.CSI("%dB", amount)
		self.Terminal.CSI(CursorDown.Params(amount))
	} else if amount < 0 {
		self.Up(-amount)
	}
}

// Nothing for zero
func (self *Cursor) Up(amount int) {
	if amount == 1 {
		//self.CSI("A")
		self.Terminal.CSI(CursorUp.Params())
	} else if amount > 1 {
		//self.CSI("%dA", amount)
		self.Terminal.CSI(CursorUp.Params(amount))
	} else if amount < 0 {
		self.Down(-amount)
	}
}
