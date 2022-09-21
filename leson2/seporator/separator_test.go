package separator

import (
	"reflect"
	"testing"
)


func TestDividedNum(t *testing.T) {
	var dividedTest = []struct {
		in  int
		out [3]int
	}{
		{123, [3]int{1, 2, 3}},
		{546, [3]int{5, 4, 6}},
		{111, [3]int{1, 1, 1}},
		{956, [3]int{9, 5, 6}},
		{762, [3]int{7, 6, 2}},
	}

	for _, tc := range dividedTest {
		got := getParts(tc.in)
		if !reflect.DeepEqual(tc.out, got) {
			t.Fatalf("Expected: %v, got: %v", tc.out, got)
		}
	}
