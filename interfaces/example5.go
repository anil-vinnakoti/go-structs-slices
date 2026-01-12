package main

import "fmt"

type Store interface{
	Save(string) error
}

type DBStore struct{}
func (DBStore)Save(s string)error{
	fmt.Println("saved to DB:", s)
	return  nil
}

type MockStore struct{}
func (MockStore) Save(s string)error{
	fmt.Println("Mock save:", s)
	return  nil
}

func process(store Store){
	store.Save("Data...")
}

// Example 5️⃣ Interface for Testing (real-world)
func example5() {
	process(DBStore{})
	process(MockStore{})
	
}