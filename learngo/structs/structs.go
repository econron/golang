package structs

import "math"

type Rectangle struct {
	Width float64
	Height float64
}

type Circel struct {
	Radius float64
}

type Shape interface {
	Area() float64
}


func Perimenter(r Rectangle) float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circel) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}