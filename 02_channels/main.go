package main

import (
	"fmt"
	"time"
)

func count(c chan int) {
	for i := 1; true; i++ {
		c <- i
		time.Sleep(time.Second)
	}
}

func main() {
	c := make(chan int)
	go count(c)
	fmt.Println("Let's count!")
	for i := 0; i < 3; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Enough counting!")
}
