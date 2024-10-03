package main

import "testing"

func TestAdd(t *testing.T) {
	myArray := make([]int, 5)
	myArray[0] = 4
	myArray[1] = 3
	myArray[2] = 5
	myArray[3] = 3
	myArray[4] = 5

	got := removeDuplicates(myArray)
	want := []int{4, 3, 5}

	for intiger := range len(got) {
		if got[intiger] != want[intiger] {
			t.Errorf("got %q, wanted %q", got[intiger], want[intiger])
		}
	}
}
