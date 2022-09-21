package geometry

import (
	"fmt"
	"os"

)

type Rectangle struct {
	w, h float64
}


func (r *Rectangle) Area() float64 {
	return r.w * r.h
}


func (r *Rectangle) String() string {
	return fmt.Sprintf("Rectangle {Width: %.2f, Height: %.2f}", r.w, r.h)
}

func NewRectangle() Rectangle {
	var w, h float64

	fmt.Print("Enter width of rectangle: ")
	_, err1 := fmt.Scanln(&w)
	if err1 != nil {
		fmt.Printf("Error: %v\n\n", ErrInput)
		os.Exit(0)
	}

	fmt.Print("Enter heigth of rectangle: ")
	_, err2 := fmt.Scanln(&h)
	if err2 != nil {
		fmt.Printf("Error: %v\n\n", ErrInput)
		os.Exit(1)
	}

	return Rectangle{w, h}
}
