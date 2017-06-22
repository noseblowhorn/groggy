package world

import "groggy/model"
import "math/rand"

type LevelGenerator interface {
	generate() *Level
}



func GenerateBigEmptyLevel() *Level {
	floorSymbol := new(model.PrintableEntity)
	floorSymbol.Glyph = '.'

	seedTile := new(Tile)
	seedTile.Passable = true
	seedTile.Transparent = true
	seedTile.Glyph = floorSymbol

	level := NewLevel(80, 20, seedTile)

	return level
}

func GenerateCavesLevel() *Level {
	floorSymbol := new(model.PrintableEntity)
	floorSymbol.Glyph = '.'

	wallSymbol := new(model.PrintableEntity)
	wallSymbol.Glyph = '#'

	seedTile := new(Tile)
	seedTile.Passable = false
	seedTile.Transparent = false
	seedTile.Glyph = wallSymbol

	floorTile := new(Tile)
	floorTile.Passable = true
	floorTile.Transparent = true
	floorTile.Glyph = floorSymbol

	level := NewLevel(80, 20, seedTile)

	//x := rand.Int() % 40 + 20
	//y := rand.Int() % 10 + 5

	x := 10
	y := 10

	totalIterations := 0

	for i := 0;; {
		var dx, dy int

		if (level.Tiles[x][y] == seedTile) {
			i++
			level.Tiles[x][y] = floorTile
		}

		if (rand.Int() % 2 == 0) {
			dx = rand.Int() % 3 - 1
			dy = 0
		} else {
			dx = 0
			dy = rand.Int() % 3 - 1
		}

		if (x < 5 && dx == -1) {
			if (rand.Int() % 100 < (6 - x) * 20) {
				dx = 0
			}
		}
		if (y < 2 && dy == -1) {
			if (rand.Int() % 100 < (3 - y) * 50) {
				dy = 0
			}
		}
		if (x > 75 && dx == 1) {
			if (rand.Int() % 100 < (x - 73) * 20) {
				dx = 0
			}
		}
		if (y > 17 && dy == 1) {
			if (rand.Int() % 100 < (y - 16) * 50) {
				dy = 0
			}
		}

		x += dx
		y += dy

		totalIterations++

		if (i > 600 || totalIterations > 10000) {
			break
		}
	}

	return level
}