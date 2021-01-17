package main

import (
	"fmt"
	"time"
)

func main() {
	c := time.Tick(time.Second)
	quit := time.After(time.Second * 3)
	i := 1

	fmt.Println("Let's count!")
	for {
		select {
		case <-c:
			fmt.Println(i)
			i++
		case <-quit:
			fmt.Println("Enough counting!")
			return
		default:
			fmt.Print(".")
			time.Sleep(time.Second / 2)
		}
	}
}
