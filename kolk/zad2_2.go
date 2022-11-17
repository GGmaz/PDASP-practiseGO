package main

import (
	"fmt"
)

func findMin(s []int, c chan int) {
	min := 110
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	c <- min
}

func main() {

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	c := make(chan int) // need to create channel
	go findMin(s[:len(s)/4], c)
	go findMin(s[5:10], c)
	go findMin(s[10:15], c)
	go findMin(s[15:20], c)
	x, y, z, v := <-c, <-c, <-c, <-c // receive from c
	fmt.Println(x, y, z, v)
}
