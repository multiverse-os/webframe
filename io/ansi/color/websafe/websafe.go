package websafe

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

// NOTE: The remainder of the colors subtracting 16 original ANSI colors, and
// the 24 grayscale colors is the standard 216 websafe colors, which is why this
// pallete will be named websafe
