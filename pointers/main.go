package main

import (
	"errors"
	"fmt"
)

type student struct{
	name string
	age int
}

func main() {
	var intPtr *int			// stores the address of an int
	// fmt.Println(*intPtr)	// throws panic as the value at that address is nil

	age := 10
	intPtr = &age			// & gives the address of age
	fmt.Println(intPtr)		// gives the address  
	fmt.Println(*intPtr)	// * gives the value stored at the address


	increment(age)			// pass by value (default behaviour of GO, everything passed by value)
	fmt.Println(age)		// 10


	incrementWithPointer(&age)	// pass by reference
	fmt.Println(age)		// 11

	// for structs
	s1 := student{name: "Anil", age: 25}

	// without pointer
	updateAge(s1, 26)		// passed by value
	fmt.Println(s1)			// age is 25
	
	// with pointer
	updateAgeWithPointer(&s1, 28)
	fmt.Println(s1)			// age is 28

	previousAge1, err1 := updateAgeTwo(&s1, 0)
	if err1 != nil{
		fmt.Println(err1)			// err
	}
	fmt.Println(previousAge1)		// previous age

	previousAge2, err2 := updateAgeTwo(&s1, 17)

	if err2 != nil{
		fmt.Println(err2)			// err
	}
	fmt.Println(previousAge2)		// updated age


}

func incrementWithPointer(num *int){
	*num++
}

func increment(num int){
	num++
}

func updateAge(stud student, newAge int){
	stud.age = newAge		
}

func updateAgeWithPointer(stud *student, newAge int){
	stud.age = newAge		// we dont have to user * in the beginning as GO internally manages for structs
}

func updateAgeTwo(stud *student, newAge int) (int, error){
	if newAge == 0{
		return stud.age, errors.New("Age cannot be 0")
	}

	perviousAge := stud.age
	stud.age = newAge
	return  perviousAge, nil
}