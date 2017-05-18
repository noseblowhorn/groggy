package widgets

import "github.com/nsf/termbox-go"

func printString(x, y int, s string, fg termbox.Attribute) {
	printStringWithBackground(x, y, s, fg, termbox.ColorBlack)
}

func printStringWithBackground(x, y int, s string, fg termbox.Attribute, bg termbox.Attribute) {
	for i, r := range []rune(s) {
		termbox.SetCell(x + i, y, r, fg, bg)
	}
}