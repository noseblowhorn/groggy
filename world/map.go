package world

import "groggy/model"

type Tile struct {
	Passable    bool
	Transparent bool
	Glyph       *model.PrintableEntity
}

func (t *Tile) clone() *Tile {
	newTile := new(Tile)
	newTile.Glyph = t.Glyph
	newTile.Passable = t.Passable
	newTile.Transparent = t.Transparent

	return newTile
}

type Level struct {
	Tiles  [][]*Tile
	Width  int
	Height int
}

func NewLevel(width int, height int, seedTile *Tile) *Level {
	l := new(Level)
	l.Width = width
	l.Height = height
	tiles := make([][]*Tile, width)
	for i := 0;i < width;i++ {
		tiles[i] = make([]*Tile, height)
	}

	for i := 0;i < width;i++ {
		for j := 0;j < height;j++ {
			tiles[i][j] = seedTile.clone()
		}
	}

	l.Tiles = tiles

	return l
}