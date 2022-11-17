package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func main() {
	var m sync.Mutex
	c := Counter{m, 0}
	ch := make(chan int)

	go foo(&c, ch, 1)
	go foo(&c, ch, 2)
	go foo(&c, ch, 3)
	go foo(&c, ch, 4)

	fmt.Println(<-ch)
}

func foo(c *Counter, ch chan int, i int) {
	for {
		fmt.Println(i)
		if c.value >= 100 {
			fmt.Println(c.value)
			ch <- 100
		}

		c.Increment()
	}
}
