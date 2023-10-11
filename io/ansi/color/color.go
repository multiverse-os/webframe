package color

// NOTE How it works now:
// ansi.Color(Bright(Blue)).Text("test")
// ansi.Color(Blue).Text("test")
// ansi.Color(Bright(Blue)).Background("test")
// ansi.Color(Blue).Background("test")

// NOTE: This is what we want:
// color.Blue("test")
// color.Foreground.Red()
// color.Background.Red()
// color.Red()

//var Reset string = fmt.Sprintf("%s", color(0))
//
//func (self color) Open() string {
//	return ansi(color, self.palette().add(39))
//}
//
//func Bg(code int) int { return (code + 10) }
//
//// TODO: 39 is default foreground reset
////       49 is default background reset
//func (self color) Close() string {
//	//return ansi(sgr.String(), self.palette().add(49))
//	// TODO: Would be better to close the corrrect tag so we can maintain
//	// different tags together
//	switch {
//	case (code >= 40 && code < 49), (code >= 100 && code < 107):
//		return color(params(49))
//	default:
//		return color(params(39))
//	}
//}
//
//func (self color) Text(text string) string {
//	return fmt.Sprintf("%s%d%s", self.Open(), text, self.Close())
//}
//
//func (self color) Background(text string) string {
//	return fmt.Sprintf("%s%d%s", self.add(background), text, self.add(background).Off())
//}
//
//// Alias functions
//func (self color) Fg(text string) string { return self.Text(text) }
//func (self color) Bg(text string) string { return self.Background(text) }
//
//func ansi(controlCode code, codes ...code) string {
//	return fmt.Sprintf("%s%s%s", controlCode, params(codes))
//}
