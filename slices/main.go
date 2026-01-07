package main

import (
	"fmt"
)

func main(){
	// nil Slice
	var nilSlice []int
	fmt.Println(nil == nilSlice)
	
	// empty slice
	emptySlice := []int{}
	fmt.Println(nil == emptySlice)

	// using make
	sliceOne := make([]int, 0)						// make(type, length, capacity?)
	fmt.Println(sliceOne)
	
	// append
	exampleSlice := []int{5,4,8,7}
	sliceOne = append(sliceOne, 15, 20)				// append returns a new slice after inserting the value to the slice
	sliceOne = append(sliceOne, exampleSlice...)	// we can spread a slice using dot notation
	fmt.Println(sliceOne)

	// slicing a slice
	// [:]   => index 0 to last
	// [1:]  => index 1 to last
	// [:6]  => index 0 to 5 (excludes the mentioned index)
	// [3:7] => index 3 to 6
	slicedSlice := exampleSlice[:2]
	fmt.Println(slicedSlice)

	// passing capacity to slice with slice expression
	slicedTwo := exampleSlice[:3:4]		//[start : end : max(capacity)]
	fmt.Println(slicedTwo)

	// copy function
	base := []int{1, 2, 3, 4}

	safe := make([]int, 2)
	copy(safe, base[:2])

	safe[0] = 99

	fmt.Println(base) // [1 2 3 4]
	fmt.Println(safe) // [99 2]


	// NOTE
	// Append beyond capacity: If a slice has no remaining capacity, append allocates a new backing array, so changes do not affect other slices.
	// Append within capacity: If a slice still has available capacity, append reuses the same backing array, causing a ripple effect on related slices.

	// NOTE
	// Copy allocates a new backing array, the slice no longer shares memory, so changes can’t ripple
}