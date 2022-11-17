package main

import "fmt"

func hello(from string) {
	for i := 1; i < 100000000; i++ {
	}
	fmt.Println("Hello from : " + from)
}
func main() {
	hello("program")
	go hello("Go routine")
	go func(caller string) {
		fmt.Println("Anonymous f: called by " + caller)
	}("Go routine")
	fmt.Scanln()
}
