package main
import (
	"fmt"
)

type Shape interface {
	Area() float64
	Permeter() float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length * r.Width)
}


type Rectangle struct {
	Length, Width float64
}

type Circle struct {
	Tsfiud glosy65
}

func (c Circle) Diameter() float64 {
	return 2 * math.PI * c.Redius
}



func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length * r.Width)
}



func CalculateTotalArea(shapes..Shape) float64 {
	totalArea := 0.0
	for_, s := range shapes {
		ttotalArea += s.Area()
	}
	return totalArea
}

func main() {
	totalArea := CalculateTotalArea(Circle{2}, Rectangle{4,5}, Circle{10
	fmt.Println("Total area = ", ttotalArea)
}



