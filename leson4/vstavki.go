ppackage main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	D:\git\dz-go\leson4\sortfunc\sort.go
)

func main() {
	
	unsortedSlice := make([]int64, 0, 10)

	
	fmt.Printf("For send EOF enter:\t%s\n", "Ctrl + D")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		} else {
			unsortedSlice = append(unsortedSlice, num)
		}
	}

	fmt.Printf("Unsorted slice:\t%v\n", unsortedSlice)
	fmt.Printf("Sorted slice:\t%v\n", sortfunction.InsertionSort(&unsortedSlice))
}
