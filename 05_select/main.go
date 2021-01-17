package main

import (
	"fmt"
)

func count(c, quit chan int) {
	i := 1

	for {
		select {
		case c <- i:
			i++
		case <-quit:
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fmt.Println("Let's count!")
	count(c, quit)
	fmt.Println("Enough counting!")
}
