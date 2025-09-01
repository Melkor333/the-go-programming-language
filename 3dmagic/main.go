package main

import (
	"fmt"
)

func main() {
	fmt.Println(`<svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">`)
	x := polygon(3, 20., 50.)
	fmt.Printf(`<path  style="fill:none;stroke:green;stroke-width:1" d="M%v %v`, x[0].x, x[0].y)
	for _, i := range x[1:] { 
		fmt.Printf("L%v %v ", i.x, i.y)
	}
	fmt.Println("Z\" />")
	fmt.Println("</svg>")

}
