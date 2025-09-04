// function to render all ceoordinates in a raster 0-1
package obj

import (
	"math"
)

type Vertex struct{
	X, Y, Z float64
	Edges []*Vertex
}

type Coord struct{
	X, Y float64
}

// coordinates, but they are supposed to be connected!
type Coords []*Coord

func RegularPolygon(edgeCount int, rad, center float64) Coords {
	o := make(Coords, edgeCount)
	// glichschenkligs 3egg, berechnemer opposite side
	for i := range edgeCount {
		// get degree and convert to rad
		degree := 360. / float64(edgeCount) * float64(i) * math.Pi / 180
		x := math.Cos(degree)
		y := math.Sin(degree)
		// resize and shift around center
		x = (x * rad) + center
		y = (y * rad) + center
		o[i] = &Coord{x, y}
	}
	return o
}
