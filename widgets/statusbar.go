package widgets

import (
	"groggy/world"
	"github.com/nsf/termbox-go"
)

func DrawStatusBar(x, y int, worldState *world.WorldState) {
	printString(x, y, worldState.PlayerCharacter.Name, termbox.ColorWhite)
}