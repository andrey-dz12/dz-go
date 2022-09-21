package sortfunction

import (
	"fmt"
	"reflect"
	"testing"
)


func TestInsertionSort(t *testing.T) {
	unsortedSlice := []int64{10, 6, 1, 3, 5, 9, 2, 4, 8, 7}
	expSlice := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if got := InsertionSort(&unsortedSlice); !reflect.DeepEqual(expSlice, got) {
		t.Errorf("Expected sorted slice %v, got %v", expSlice, got)
	}
}

func ExampleInsertionSort() {
	unsortedSlice := []int64{10, 6, 1, 3, 5, 9, 2, 4, 8, 7}
	fmt.Printf("%v", InsertionSort(&unsortedSlice))
	// Output: [1 2 3 4 5 6 7 8 9 10]
}

