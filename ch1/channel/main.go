package main

import (
	"fmt"
)

// 單向 Channel
func printCount(c <-chan int) {
	num := 0
	for num >= 0 {
		num = <-c
		fmt.Println(num)
	}
}

func main() {
	c := make(chan int)
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, -1}

	go printCount(c)

	for _, n := range a {
		c <- n
	}
	// time.Sleep(time.Microsecond * 1)
	fmt.Println("End of main")
}
