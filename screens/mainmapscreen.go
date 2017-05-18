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

func MainMapLoop(worldState *world.WorldState) {
	for {
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