package model

type entity interface {
	components() []component
}

type component interface {

}

type MapEntity struct {
	Symbol *PrintableEntity
	X, Y int
}

type PrintableEntity struct {
	Glyph rune
}