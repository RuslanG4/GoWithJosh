package main

import "testing"

func RomanTest(t *testing.T) {
	myArray := make([]string, 1)
	myArray[0] = "III"
	myArray[1] = "IV"
	myArray[0] = "M"

	got :=  make([]int, 1)
	got[0] = romanNumeral(myArray[0])
	want := []int{3, 4, 1000}

	for intiger := range len(got) {
		if got[intiger] != want[intiger] {
			t.Errorf("got %q, wanted %q", got[intiger], want)
		}
	}
}
