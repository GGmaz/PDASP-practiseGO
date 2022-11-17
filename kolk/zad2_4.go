package main

import "fmt"

func main() {
	ch := make(chan int)

	go findfib(15, ch)

loop:
	for {
		select {
		case val, ok := <-ch:
			if !ok {
				break loop
			}
			fmt.Printf("%d ", val)
		}
	}

}

func findfib(n int, iCh chan int) {
	defer close(iCh)
	a := 0
	b := 1
	i := 2
	iCh <- 0
	iCh <- 1
	for i < n {
		b = a + b
		a = b - a
		iCh <- b
		i++
	}
}
