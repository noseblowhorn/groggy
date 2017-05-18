package main

//import "github.com/jroimartin/gocui"
import "github.com/nsf/termbox-go"

import (
	"groggy/world"
	"groggy/model"
	"groggy/widgets"
)

func main() {
	playerSymbol := new(model.PrintableEntity)
	playerSymbol.Glyph = '@'

	playerPrintable := new(model.MapEntity)
	playerPrintable.Symbol = playerSymbol
	playerPrintable.X = 10
	playerPrintable.Y = 10

	player := new (world.Player)
	player.Name = "Brajanek"
	player.MapEntity = playerPrintable

	worldState := new (world.WorldState)
	worldState.PlayerCharacter = player
	worldState.CurrentLevel = world.GenerateBigEmptyLevel()

	level := worldState.CurrentLevel

	termbox.Init()
	termbox.HideCursor()

	for true {
		widgets.DrawMainMap(0, 2, worldState)
		widgets.DrawStatusBar(0, 23, worldState)

		termbox.Flush()

		event := termbox.PollEvent()

		if event.Key == termbox.KeyEsc {
			termbox.HideCursor()
			break
		}

		switch event.Key {
			case termbox.KeyArrowLeft:
				if player.MapEntity.X > 0 {
					player.MapEntity.X -= 1
				}
			case termbox.KeyArrowRight:
				if player.MapEntity.X < level.Width - 1 {
					player.MapEntity.X += 1
				}
			case termbox.KeyArrowUp:
				if player.MapEntity.Y > 0 {
					player.MapEntity.Y -= 1
				}
			case termbox.KeyArrowDown:
				if player.MapEntity.Y < level.Height - 1 {
					player.MapEntity.Y += 1
				}
		}
	}

	termbox.Close()
}
