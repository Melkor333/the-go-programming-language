// function to render all ceoordinates in a raster 0-1
package main

import (
	"math"
)

type coord struct{
	x, y float64
}


func polygon(edgeCount int, rad, center float64) []coord {
	o := make([]coord, edgeCount)
	// glichschenkligs 3egg, berechnemer opposite side
	for i := range edgeCount {
		// get degree and convert to rad
		degree := 360. / float64(edgeCount) * float64(i) * math.Pi / 180
		x := math.Cos(degree)
		y := math.Sin(degree)
		// resize and shift around center
		x = (x * rad) + center
		y = (y * rad) + center
		o[i] = coord{x, y}
	}
	return o
}
