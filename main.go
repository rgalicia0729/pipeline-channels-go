package main

import "fmt"

func generateListInt(in chan<- int) {
	for i := 1; i <= 10; i++ {
		in <- i
	}
	close(in)
}

func doubleListInt(in <-chan int, out chan<- int) {
	for value := range in {
		out <- value * 2
	}

	close(out)
}

func printListInt(out chan int) {
	for value := range out {
		fmt.Println(value)
	}
}

func main() {
	in := make(chan int)
	out := make(chan int)

	go generateListInt(in)
	go doubleListInt(in, out)

	printListInt(out)
}
