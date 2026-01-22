package main

import "fmt"

func main() {
	generatedNums := generator(1,4,8,9)
	squaredNums := squareNums(generatedNums)

	for num := range squaredNums{
		fmt.Println("squered number:", num)
	}
}

func generator(nums ...int) chan int {
	numsChannel := make(chan int)
	go func() {
		for _, num := range nums {
			numsChannel <- num
		}
		close(numsChannel)
	}()

	return numsChannel
}

func squareNums(ch <-chan int) <-chan int {
	squared := make(chan int)
	go func() {
		for num := range ch {
			squared <- num * num
		}
		close(squared)
	}()

	return  squared
}
