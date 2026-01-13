package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	ID       int      `json:"user_id"`
	Name     string   `json:"name,omitempty"`
	Age      int      `json:"age"`
	Password string   `json:"-"`
	Roles    []string `json:"roles"`
}

type Person struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Active bool `json:"is_active"`
}

func main() {
	
	// marshal
	marshallDataJson := marshalData()
	fmt.Println(marshallDataJson)

	// unmarshall
	unmarshalDataStruct := unmarshalData()
	fmt.Printf("%+v\n", unmarshalDataStruct)

	// JSON encoding(creates a json file)
	jsonEncoding()

	// JSON decoding(read from a json file)
	jsonDecoding()


}

func marshalData()string{
	user := User{
		ID:       1,
		Name:     "Anil",
		Age:      27,
		Password: "my-password",
		Roles:    []string{"admin", "platform-admin"},
	}

	j, err := json.Marshal(user)
	if err != nil {
		fmt.Println("error while marshalling data:", err)
	}
	return  string(j)
}

func unmarshalData()User{
	jsonData := `{"user_id":1,"name":"Anil","age":27,"roles":["admin","platform-admin"]}`

	var user User

	err := json.Unmarshal([]byte(jsonData), &user)
	if err != nil{
		fmt.Println("error unmarshalling json:", err)
	}
	return  user
}

func jsonEncoding(){
	person := Person{Name: "Annie", Age: 26,Active: true}

	jsonFile, createFileErr := os.Create("output.json")
	if createFileErr != nil{
		fmt.Println("error occured while creating file:", createFileErr)
	}
	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	encodingErr := encoder.Encode(person)
	if encodingErr != nil{
		fmt.Println("error encoding person:", encodingErr)
	}
}

func jsonDecoding(){
	var person Person
	openedFile, fileOpenErr := os.Open("dummy.json")
	if fileOpenErr != nil{
		fmt.Println("error while opening file:", fileOpenErr)
	}
	defer openedFile.Close()

	decoder := json.NewDecoder(openedFile)
	err := decoder.Decode(&person)
	if err != nil{
		fmt.Println("error occured while decoding person:", err)
	}

	fmt.Printf("%+v\n", person)
}