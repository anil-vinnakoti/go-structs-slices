package main

import "fmt"

func main() {
	// nil map
	var nilMap map[string]int
	fmt.Println(len(nilMap))
	fmt.Println(nilMap["age"])		// gives 0 as it is nil
	// nilMap["age"] = 27 			// causes pani as we cant assign values to nil map

	// empty map
	// emptyMap1 := map[string]int{}
	// emptyMap2 := make(map[string]int)
	// var emptyMap3 map[string]int  = map[string]int{}

	nameAge := map[string]int{
		"four": 25,
	}
	nameAge["one"] = 21
	nameAge["two"] = 27
	nameAge["three"] = 22

	fmt.Println(len(nameAge))
	fmt.Println(nameAge["one"])

	// comma ok idiom
	value, ok := nameAge["fur"]
	if !ok {
		fmt.Println("value not found in map")
	}
	fmt.Println(value)

	//over writing a key
	nameAge["four"] = 33
	fmt.Println(nameAge)

	data := map[int][]int{
		1:[]int{1,2,3,},
		2:{4,5,6},
	}

	fmt.Println(data)
	
}