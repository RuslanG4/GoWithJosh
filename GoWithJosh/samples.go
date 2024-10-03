package main

import "fmt"

func romanNumeral( input string )( int ){

	MAX_ROMAN_NUMERALS := int(7)
	romanNumeralStr:= [7]rune{'I', 'V', 'X', 'L', 'C', 'D', 'M'}
	romanNumeralValues:=[7]int{ 1, 5, 10, 50, 100, 500, 1000 }
	
	var previousRomanNumeral rune

	year :=int(0)

	for i := len(input) -1; i >=0 ; i--{

		for j := int(0); j < MAX_ROMAN_NUMERALS; j++{
			
			if romanNumeralStr[j] == rune(input[i]){

				if romanNumeralValues[j] < year && romanNumeralStr[j] != previousRomanNumeral{
					year-=romanNumeralValues[j] // 
				}else{
					year+=romanNumeralValues[j]
				}
				previousRomanNumeral = romanNumeralStr[j]
				
			}
		}

	}

	
	return year
}


func removeDuplicates(array []int) []int {
	var tempArray []int
	var dupeArray []int

	for i := range len(array) {
		current := array[i]
		dupeTaken := false
		for k := 0; k < len(array); k++ {
			//if lhs value == rhs value, the position of the values are not the same ( don't get yourself )
			if current == array[k] && i != k {
				for dupes := range len(dupeArray) {
					if dupeArray[dupes] == array[k] {
						dupeTaken = true
					}
				}
				if !dupeTaken {
					tempArray = append(tempArray, array[i])
					dupeArray = append(dupeArray, array[i])
					dupeTaken = true
				}
				//fmt.Printf("len=%d cap=%d %v\n", len(dupeArray), cap(dupeArray), dupeArray)
			}
		}
		if !dupeTaken {
			tempArray = append(tempArray, array[i])
		}

	}
	fmt.Printf("len=%d cap=%d %v\n", len(tempArray), cap(tempArray), tempArray)
	return tempArray
}

func main() {
	myArray := make([]int, 5)
	myArray[0] = 4
	myArray[1] = 3
	myArray[2] = 5
	myArray[3] = 3
	myArray[4] = 5
	//fmt.Println(myArray)
	removeDuplicates(myArray)
	//fmt.Println(myArray)

}
