package main

import (
	"fmt"
	"os"

	"D:\git\dz-go\leson2\geometry"
	"D:\git\dz-go\leson2\seporator"
)

func main() {
	for stop := false; !stop; {
		switch getTask() {

		case 0:
			fmt.Println("Goodbye!")
			stop = true
		case 1:
			r := geometry.NewRectangle()
			fmt.Printf("Rectangle area: %f\n\n", r.Area())
		case 2:
			c := geometry.NewCircle()
			fmt.Printf("Circle area: %6.2f\nCircle diameter: %6.2f\nCircle length: %6.2f\n\n", c.Area(), c.Diameter(), c.Length())
		case 3:
			separator.DividedNum()
		case 4:
			c := geometry.NewCircle()
			r := geometry.NewRectangle()
			compareShapesAreas(c, r)
		default:
			fmt.Println("There is no such task, goodbye!")
			stop = true
		}
	}
}

func getTask() int {
	var taskId int
	fmt.Printf(" %s\n %s\n %s\n %s\n %s\nSelect task: ",
		"1 - Get retangle area",
		"2 - Get circle area, diameter and length",
		"3 - Parts of num (length = 3)",
		"4 - Compare areas of shapes",
		"0 - Exit")

	_, err := fmt.Scan(&taskId)
	if err != nil {
		fmt.Printf("Incorrect task Id%v!", taskId)
		os.Exit(1)
	}

	return taskId
}

