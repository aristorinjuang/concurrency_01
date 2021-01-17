package main

import (
	"fmt"
	"time"
)

func count() {
	for i := 1; true; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {
	go count()
	fmt.Println("Let's count!")
	time.Sleep(time.Second * 3)
	fmt.Println("Enough counting!")
}
