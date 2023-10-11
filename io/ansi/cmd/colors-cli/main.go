package main

import (
	"fmt"

	color "github.com/multiverse-os/color"
)

func main() {
	// TODO: Fix the order
	// Primary 8 ANSI Colors
	fmt.Printf(" +=====================+\n")
	fmt.Printf(" | Primary ANSI Colors |\n")
	fmt.Printf(" +=====================+\n")
	fmt.Printf(" | %s   | %s   |  \n", color.Black("Black"), color.BlackBg("BlackBg"))
	fmt.Printf(" | %s  | %s  | \n", color.Red("Maroon"), color.RedBg("MaroonBg"))
	fmt.Printf(" | %s   | %s   |  \n", color.Green("Green"), color.GreenBg("GreenBg"))
	fmt.Printf(" | %s   | %s   | \n", color.Yellow("Olive"), color.YellowBg("OliveBg"))
	fmt.Printf(" | %s    | %s    |  \n", color.Blue("Blue"), color.BlueBg("BlueBg"))
	fmt.Printf(" | %s | %s |  \n", color.Magenta("Magenta"), color.MagentaBg("MagentaBg"))
	fmt.Printf(" | %s    | %s    |  \n", color.Cyan("Cyan"), color.CyanBg("CyanBg"))
	fmt.Printf(" | %s  | %s  | \n", color.White("Silver"), color.WhiteBg("SilverBg"))
	fmt.Printf(" +---------+-----------+\n")
}
