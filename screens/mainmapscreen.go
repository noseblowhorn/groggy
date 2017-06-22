package screens

import (
	"groggy/world"
	"github.com/nsf/termbox-go"
	"groggy/widgets"
)

func mainMapProcessInput(worldState *world.WorldState, event termbox.Event) {
	level := worldState.CurrentLevel
	player := worldState.PlayerCharacter

	switch event.Key {
	case termbox.KeyArrowLeft:
		if level.IsTilePassable(player.MapEntity.X - 1, player.MapEntity.Y) {
			player.MapEntity.X -= 1
		}
	case termbox.KeyArrowRight:
		if level.IsTilePassable(player.MapEntity.X + 1, player.MapEntity.Y) {
			player.MapEntity.X += 1
		}
	case termbox.KeyArrowUp:
		if level.IsTilePassable(player.MapEntity.X, player.MapEntity.Y - 1) {
			player.MapEntity.Y -= 1
		}
	case termbox.KeyArrowDown:
		if level.IsTilePassable(player.MapEntity.X, player.MapEntity.Y + 1) {
			player.MapEntity.Y += 1
		}
	}
}

func MainMapLoop(worldState *world.WorldState) {
	for {
		worldState.CurrentLevel.CalculateVisibility(
			worldState.PlayerCharacter.MapEntity.X,
			worldState.PlayerCharacter.MapEntity.Y)

		widgets.DrawMainMap(0, 2, worldState)
		widgets.DrawStatusBar(0, 23, worldState)

		termbox.Flush()

		event := termbox.PollEvent()

		if event.Key == termbox.KeyEsc {
			break
		}

		mainMapProcessInput(worldState, event)
	}
}