package world

import "groggy/model"

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