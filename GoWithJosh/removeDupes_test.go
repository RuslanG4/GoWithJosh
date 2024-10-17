package main

import "testing"

//testing that the expected outcome of removeDuplicates array function is correct
func ExpectedOutcomeTest(t *testing.T) {
	input := []int{3, 3, 4, 4, 5}
    expected := []int{3, 4, 5}

	got, err := removeDuplicates(input)
	//error occuring in conditions
	if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }
	//verify each array position with expected
	for intiger := range len(got) {
		if got[intiger] != expected[intiger] {
			t.Errorf("got %q, wanted %q", got[intiger], expected[intiger])
		}
	}
}// end of test

//testing the length of the array
func ValidArrayLengthTest(t* testing.T){
	// array is >1 in length
	invalidArray := []int{1}
	if arrayIsValid(invalidArray) {
	 	t.Errorf("expected false for array is <=1, got true")
	}
	// array is < 3*10^4 in length
	largeArray := make([]int, 3*10000+1)
	if arrayIsValid(largeArray) {
		t.Errorf("expected false for array is > 3*10^4, got true")
	}
	// array is valid in length
	validArray := []int{1, 2, 3}
	if !arrayIsValid(validArray) {
		t.Errorf("expected true for valid array, got false")
	}
}//end of test

//test that the individual values are -100 <= nums[i] <= 100
func ArrayMembersAreInCorrectRange(t* testing.T){
	// array members are invalid >100
	invalidHighArray := []int{101,102,103}
	if arrayValuesAreInRange(invalidHighArray) {
	 	t.Errorf("expected false for <= 100, got true")
	}
	// array members are invalid < -100
	invalidLowArray := []int{-101,10,15}
	if arrayValuesAreInRange(invalidLowArray) {
	 	t.Errorf("expected false for <= -100, got true")
	}
	// array members are valid in value
	validArray := []int{-100, 20, 100}
	if !arrayValuesAreInRange(validArray) {
		t.Errorf("expected true for valid array, got false")
	}
}//end of test

//test array is in a non decreasing order
func ArrayIsInNonDecreasingOrder(t* testing.T){
	// valid outcome of increasing order
    sortedArray := []int{5, 6, 6, 7, 8, 8}
    if !arraySortedInNonDecreaseOrder(sortedArray) {
        t.Errorf("expected true for sorted array, got false")
    }
	// array members are non increasing
	unsortedArray := []int{53, 32, 61, 11, 17}
    if arraySortedInNonDecreaseOrder(unsortedArray) {
        t.Errorf("expected false for unsorted array, got true")
    }
}//end of test

