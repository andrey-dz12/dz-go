package geometry

import (
	"errors"
	"fmt"
	"math"
	"os"
)

const ErrInputText = "incorrect parameter entered"


var ErrInput = errors.New(ErrInputText)

type Circle struct {
	radius float64
}


func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}


func (c *Circle) Diameter() float64 {
	return 2 * math.Sqrt(c.Area()/math.Pi)
}


func (c *Circle) Length() float64 {
	return math.Pi * c.Diameter()
}

func (c *Circle) String() string {
	return fmt.Sprintf("Circle {Radius: %.2f}", c.radius)
}


func NewCircle() Circle {
	var rad float64

	fmt.Print("Enter radius of circle: ")
	_, err := fmt.Scan(&rad)
	if err != nil {
		fmt.Printf("Error: %v\n\n", ErrInput)
		os.Exit(1)
	}

	return Circle{rad}
}
