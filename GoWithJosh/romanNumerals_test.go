package main

import "testing"

func TestIndividualRomanNumerals(t *testing.T) {
	NUM_OF_TESTS := 7

	// input and expected out puts
	myArray := []string{ "I", "V", "X", "L", "C", "D", "M"}
	want := []int{ 1, 5, 10, 50, 100, 500, 1000}

	got :=  make([]int, NUM_OF_TESTS)

	// using function
	for i:= range len(got){
		got[i] = romanNumeral(myArray[i])
	}
	
	// compairing to expected
	for i := range len(got) {
		if got[i] != want[i] {
			t.Errorf("got %q, wanted %q", got[i], want)
		}
	}
}

func TestReversePairings(t *testing.T) {
	NUM_OF_TESTS := 6

	// input and expected out puts
	myArray := []string{ "IV", "IX", "XL", "XC", "CD", "CM"}
	want := []int{ 4, 9, 40, 90, 400, 900 }

	got :=  make([]int, NUM_OF_TESTS)

	// using function
	for i:= range len(got){
		got[i] = romanNumeral(myArray[i])
	}
	
	// compairing to expected
	for i := range len(got) {
		if got[i] != want[i] {
			t.Errorf("got %q, wanted %q", got[i], want)
		}
	}
}

func TestBigNumbers(t *testing.T){
	NUM_OF_TESTS := 1

	// input and expected out puts
	myArray := []string{ "MMCCCXXXIIIIV"}//, "IX", "XL", "XC", "CD", "CM"}
	want := []int{ 2333}//, 9, 40, 90, 400, 900 }

	got :=  make([]int, NUM_OF_TESTS)

	// using function
	for i:= range len(got){
		got[i] = romanNumeral(myArray[i])
	}
	
	// compairing to expected
	for i := range len(got) {
		if got[i] != want[i] {
			t.Errorf("got %d, wanted %d", got[i], want)
		}
	}
}

// reverse pairings 
// big numbers with reverses
// big numbers no reverses