package main

import (
	"fmt"
)

type student struct {
	firstName string
	lastname  string
	age       int
	subjects  []string
}

func main() {
	var student1 student
	fmt.Printf("%+v\n", student1)

	// initializing a struct
	// implicit
	student1 = student{"anil", "kumar", 25, []string{"maths", "science"}}
	fmt.Printf("%+v\n", student1)

	// explicit
	student2 := student{
		firstName: "vinay",
		lastname:  "kumar",
	}
	fmt.Printf("%+v\n", student2)
	fmt.Println("first name of student2:", student2.firstName)

	// rewrite or assign fields data
	student2.firstName = "Chanti"
	student2.age = 26
	student2.subjects = []string{"physics"}
	fmt.Printf("%+v\n", student2)

	// anonymous struct
	anonymousStruct := struct{
		name string
		age  int
	}{
		name: "mister",
		age:  99,
	}

	fmt.Printf("%+v\n", anonymousStruct)

}
