package areacalc

import "strings"

const pi = 3.14159

type Shape interface {
	Area() float64
	Type() string
}

type Rectangle struct {
	a float64
	b float64
	name string
}

func NewRectangle(a float64, b float64, name string) *Rectangle {
	return &Rectangle{a, b, name}
}

func (r *Rectangle) Area() float64 {
	return r.a*r.b
}

func (r *Rectangle) Type() string {
	return r.name
}

type Circle struct {
	r float64
	name string
}

func NewCircle(r float64, name string) *Circle {
	return &Circle{r, name}
}

func (c *Circle) Area() float64 {
	return pi*c.r*c.r
}

func (c *Circle) Type() string {
	return c.name
}

func AreaCalculator(figures []Shape) (string, float64) {

	if len(figures) == 0 {
		return "", 0
	}

	var strB strings.Builder
	var sumArea float64

	for _, i := range figures {
		sumArea += i.Area()

		strB.WriteString(i.Type())
		strB.WriteRune('-')
	}

	return strB.String()[:len(strB.String())-1], sumArea
}
