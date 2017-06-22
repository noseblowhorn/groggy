package world

import "groggy/model"
import "math"

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
	VisibleMask [][]bool
	SeenMask [][]bool
	Width  int
	Height int
}

type Point struct {
	X, Y int
}

func newPoint(x, y int) *Point {
	p := new(Point)
	p.X = x
	p.Y = y

	return p
}

func (l *Level) IsTilePassable(x, y int) bool {
	if x < 0 || y < 0 || x > l.Width - 1 || y > l.Height - 1 {
		return false
	}
	return l.Tiles[x][y].Passable
}

func (l *Level) IsTileTransparent(x, y int) bool {
	return l.Tiles[x][y].Transparent
}

func createEdgePointSlice(width, height int) []Point {
	edgePoints := make([]Point, 2 * width + 2 * height)
	i := 0
	for j := 0;j < width;j++ {
		edgePoints[i] = *newPoint(j, 0)
		i++
		edgePoints[i] = *newPoint(j, height - 1)
		i++
	}
	for j := 0;j < height;j++ {
		edgePoints[i] = *newPoint(0, j)
		i++
		edgePoints[i] = *newPoint(width - 1, j)
		i++
	}

	return edgePoints
}

func sgn(x float64) float64 {
	if (x < 0) {
		return -1.0
	} else {
		return 1.0
	}
}

func (l *Level) CalculateVisibility(observerX, observerY int) {
	edgePoints := createEdgePointSlice(l.Width, l.Height)

	visibilityMap := make([][]bool, l.Width)
	for i := 0;i < l.Width;i++ {
		visibilityMap[i] = make([]bool, l.Height)
	}

	for _, point := range edgePoints {
		var deltaX, deltaY float64

		x := float64(observerX)
		y := float64(observerY)

		distX := float64(point.X - observerX)
		distY := float64(point.Y - observerY)

		if (math.Abs(distX) < math.Abs(distY)) {
			deltaY = sgn(distY)
			deltaX = distX / distY
		} else {
			deltaX = sgn(distX)
			deltaY = distY / distX
		}

		for ;y > -1 && y < float64(l.Height) && x > -1 && x < float64(l.Width); {
			visibilityMap[int(x)][int(y)] = true
			l.SeenMask[int(x)][int(y)] = true
			if !l.IsTileTransparent(int(x), int(y)) {
				break
			}
			x += deltaX
			y += deltaY
		}
	}

	l.VisibleMask = visibilityMap
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
			tiles[i][j] = seedTile
		}
	}

	l.Tiles = tiles

	l.SeenMask = make([][]bool, l.Width)
	for i := 0;i < l.Width;i++ {
		l.SeenMask[i] = make([]bool, l.Height)
	}

	return l
}