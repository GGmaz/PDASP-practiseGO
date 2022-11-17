package main

import (
	"fmt"
	"math"
)

type Trougao struct {
	a float64
	b float64
	c float64
}

func main() {
	t := Trougao{34, 45.7, 38}

	fmt.Println(povrsinaTrougla(t))
	fmt.Println(t.a)

	arr := []int{1, 2, 3, 76, 5, 767, 55, 4, 2, 33, 4, 1}
	//sort.Ints(arr)
	sortArray(arr)
	fmt.Println(arr)
}

func sortArray(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func povrsinaTrougla(t Trougao) float64 {
	s := (t.a + t.b + t.c) / 2
	t.a = 30
	return math.Sqrt(s * (s - t.a) * (s - t.b) * (s - t.c))
}
