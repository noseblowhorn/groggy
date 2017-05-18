package main

//import "github.com/jroimartin/gocui"
import "github.com/nsf/termbox-go"

import (
	"groggy/world"
	"groggy/model"
	"groggy/screens"
)

func temporaryInitializeWorld() *world.WorldState {
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

	return worldState
}

func initializeTermbox() {
	termbox.Init()
	termbox.HideCursor()
}

func main() {
	worldState := temporaryInitializeWorld()

	initializeTermbox()

	screens.MainMapLoop(worldState)

	termbox.Close()
}
