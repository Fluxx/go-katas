package main

import (
	"fmt"
)

func fibonacci(n int, channel chan int) {
	previous, current := 1, 1
	for i := 0; i < n; i++ {
		channel <- previous
		previous, current = current, previous+current
	}
	close(channel)
}

func main() {
	channel := make(chan int, 10)
	go fibonacci(cap(channel), channel)
	for i := range channel {
		fmt.Println(i)
	}
}
