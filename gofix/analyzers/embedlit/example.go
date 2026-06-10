// `embedlit`: Go 1.27 lets you initialize fields promoted from an embedded
// struct directly in the outer composite literal, dropping the nested
// `U: U{...}` layer.
package embedlit

type Coord struct {
	X, Y int
}

type Marker struct {
	Coord
	Label string
}

var Origin = Marker{
	Coord: Coord{X: 0, Y: 0},
	Label: "origin",
}
