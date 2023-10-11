package ansi

type Coordinate struct {
	X int
	Y int
	// TODO
	// Z only makes sense if we are talking about multiple layers, but in those
	// cases we just usually keep a shadow buffer
	//Z int
}
