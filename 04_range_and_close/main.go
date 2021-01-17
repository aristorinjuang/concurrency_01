package main

import (
	"fmt"
	"time"
)

func count(c chan int) {
	for i := 1; i <= cap(c); i++ {
		c <- i
		time.Sleep(time.Second)
	}

	close(c)
}

func main() {
	c := make(chan int, 3)
	go count(c)
	fmt.Println("Let's count!")
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Enough counting!")
}
