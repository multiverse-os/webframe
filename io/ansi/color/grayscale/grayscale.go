package grayscale

//  8-bit
//
//  As 256-color lookup tables became common on graphic cards, escape sequences were added to select from a pre-defined set of 256 colors:[citation needed]
//
//    ESC[ 38;5;⟨n⟩ m Select foreground color
//    ESC[ 48;5;⟨n⟩ m Select background color
//      0-  7 :  standard colors (as in ESC [ 30–37 m)
//      8- 15 :  high intensity colors (as in ESC [ 90–97 m)
//      16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
//    232-255 :  grayscale from black to white in 24 steps

const (
	//Black  = 40
	Gray1  = 232
	Gray2  = 233
	Gray3  = 234
	Gray4  = 235
	Gray5  = 236
	Gray6  = 237
	Gray7  = 238
	Gray8  = 239
	Gray9  = 240
	Gray10 = 241
	Gray11 = 242
	Gray12 = 243
	Gray13 = 244
	Gray14 = 245
	Gray15 = 246
	Gray16 = 247
	Gray17 = 248
	Gray18 = 249
	Gray19 = 250
	Gray20 = 251
	Gray21 = 252
	Gray22 = 253
	Gray23 = 254
	Gray24 = 255
	//White  = 97
)

// Grayscale With Color Names; 4 Step-wise
const (
	DarkestGray  = Gray1
	DarkerGray   = Gray4
	DarkGray     = Gray8
	Gray         = Gray12
	LightGray    = Gray16
	LighterGray  = Gray20
	LightestGray = Gray24
)
