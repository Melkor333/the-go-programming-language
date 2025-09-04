package main

import (
	"fmt"
	"3dmagic/obj"
)

func main() {
	fmt.Println(`<svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">`)
	p := obj.RegularPolygon(3, 20., 50.)
	fmt.Printf(`<path  style="fill:none;stroke:green;stroke-width:1" d="M%v %v`, p[0].X, p[0].Y)
	for _, i := range p[1:] { 
		fmt.Printf("L%v %v ", i.X, i.Y)
	}
	fmt.Println("Z\" />")
	fmt.Println("</svg>")
}
