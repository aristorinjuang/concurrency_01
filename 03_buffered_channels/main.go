package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)
	initCount := 3
	printCount := 3

	for i := 1; i <= initCount; i++ {
		c <- i
		time.Sleep(time.Second)
	}

	fmt.Println("Let's count!")
	for i := 0; i < printCount; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Enough counting!")
}
