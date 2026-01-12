package main

import "fmt"

type Counter interface {
	Inc()
}

type MyCounter struct {
	count int
}

func (c *MyCounter) Inc() {
	c.count++
}

// Example 3️⃣ Pointer Receiver & Interface
func example3() {
	var c Counter

	c = &MyCounter{}
	c.Inc()
	c.Inc()

	fmt.Println(c.(*MyCounter).count) // 2
}
