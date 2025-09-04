package main

import (
	"fmt"
	"3dmagic/obj"
	//"3dmagic/coord"
	"3dmagic/svg"
)

func main() {
	fmt.Println(`<svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">`)
	p := obj.RegularPolygon(3, 20., 50.)
	p2 := obj.RegularPolygon(7, 20., 50.)
	fmt.Println((&svg.Path{p, true}).String())
	fmt.Println((&svg.Path{p2, true}).String())

	fmt.Println("</svg>")
}
