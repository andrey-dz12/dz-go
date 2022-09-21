package separator

import (
	"fmt"
)

func getParts(num int) [3]int {
	h := num / 100
	d := (num % 100) / 10
	u := (num % 100) % 10
	return [3]int{h, d, u}
}

func DividedNum() {
	var num int

	fmt.Print("Enter a number for division by parts: ")
	_, err := fmt.Scanf("%d", &num)

	if (err == nil) && (num >= 100) && (num <= 999) {
		parts := getParts(num)
		fmt.Printf("Hundreds: %d\nDozens: %d\nUnits: %d\n\n", parts[0], parts[1], parts[2])
	} else {
		fmt.Println("Incorrect number! (100 - 999)")
	}
}
