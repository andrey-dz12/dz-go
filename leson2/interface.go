package main

import (
	"fmt"

	"D:\git\dz-go\leson2\geometry"
)

type Sizer interface {
	Area() float64
}

type Shaper interface {
	Sizer
	fmt.Stringer
}


func compareShapesAreas(c geometry.Circle, r geometry.Rectangle) {
	printArea(&c)
	printArea(&r)
	l := less(&c, &r)
	fmt.Printf("%+v is the smallest\n\n", l)
}


func less(s1, s2 Sizer) Sizer {
	if s1.Area() < s2.Area() {
		return s1
	}
	return s2
}

func printArea(s Shaper) {
	fmt.Printf("area of %s is %.2f\n", s.String(), s.Area())
}

