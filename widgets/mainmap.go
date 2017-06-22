package widgets

import (
	"groggy/world"
	"github.com/nsf/termbox-go"
)

func DrawMainMap(x int, y int, worldstate *world.WorldState) {
	level := worldstate.CurrentLevel
	player := worldstate.PlayerCharacter.MapEntity

	for i := 0;i < level.Width;i++ {
		for j := 0;j < level.Height;j++ {
			if (level.SeenMask[i][j]) {
				termbox.SetCell(i + x, j + y, level.Tiles[i][j].Glyph.Glyph, termbox.ColorWhite, termbox.ColorBlack)
			} else {
				termbox.SetCell(i + x, j + y, ' ', termbox.ColorBlack, termbox.ColorBlack)
			}
		}
	}

	termbox.SetCell(player.X + x, player.Y + y, player.Symbol.Glyph, termbox.ColorBlue, termbox.ColorBlack)
}
