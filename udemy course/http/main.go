package main

import "fmt"

type shape interface {
	getArea() float64
}

func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t Traingle) getArea() float64 {
	return 0.5 * t.base * t.heigth
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())

}

type Square struct {
	sideLength float64
}
type Traingle struct {
	heigth float64
	base   float64
}

func main() {

	sq := Square{
		sideLength: 8.0,
	}

	tr := Traingle{
		heigth: 4,
		base:   5,
	}
	printArea(sq)

	printArea(tr)

}
