package sortfunc


func InsertionSort(unsortedSlice *[]int64) []int64 {

	sortedSlice := make([]int64, len(*unsortedSlice))
	copy(sortedSlice, *unsortedSlice)

	for i := 1; i < len(sortedSlice); i++ {
		for j := i; j != 0 && sortedSlice[j] < sortedSlice[j-1]; j-- {
			sortedSlice[j-1], sortedSlice[j] = sortedSlice[j], sortedSlice[j-1]
		}
	}

	return sortedSlice
}

