package main

import (
	"fmt"
	"math"
)

// interface becomes accessible when is named from capital letter
type Geometry interface {
	area() float64
	perimeter() float64
}

type Rect struct{ width, height float64 }
type Circle struct{ radius float64 }

func (a Rect) area() float64   { return a.height * a.width }
func (a Circle) area() float64 { return math.Pi * a.radius * a.radius }

func (a Rect) perimeter() float64   { return 2 * (a.height + a.width) }
func (a Circle) perimeter() float64 { return 2 * math.Pi * a.radius }
func (a Circle) diameter() float64  { return 2 * a.radius }

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Printf("  area: %3.2f, perimeter: %3.2f\n", g.area(), g.perimeter())
}

func main() {
	r := Rect{width: 3, height: 4}
	c := Circle{radius: 5}

	// if instances have functions requested by interface then instance can be used "instead" (in place) of the imterface":
	measure(r)
	measure(c)
	MyPrintln("Rectangle:", r, "\nCircle:", c)
	PrintType(9)
	PrintType("Text")
	PrintType(3.14)
}

func MyPrintln(pars ...interface{}) {
	for i, v := range pars {
		if i > 0 {
			fmt.Print(" ")
		}

		fmt.Printf("%T:%v", v, v)
	}
	fmt.Println()
}

func PrintType(i interface{}) {
	fmt.Printf("Type of %v is %T\n", i, i)
}
