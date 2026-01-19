package main

import "fmt"

// A function is a closure if it refers to variables defined outside of it,
// and those variables stay alive as long as the function does.

func createGiftCard() func(int) int {
	amount := 100 // outer scope variable

	return func(debit int) int {
		amount -= debit // ← uses outer variable
		return amount
	}
}

func main() {
	gc1 := createGiftCard()
	gc2 := createGiftCard()

	fmt.Println(gc1(15))
	fmt.Println(gc2(10))
}
