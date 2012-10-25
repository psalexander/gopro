// goroutine project main.go
package main

import (
	"fmt"
	"time"
)

var c chan int

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready")
	c <- 1
}

func main() {
	/*
		c = make(chan int)
		go ready("Tea", 2)
		go ready("Coffee", 1)
		go ready("Coffede", 4)
		go ready("Coffe2e", 2)
		go ready("Coff2e", 4)
		go ready("Coff12de", 2)
		fmt.Println("I'm waiting,but not too long")
		//<-c
		//<-c

			i := 0
		L:
			for {
				select {
				case <-c:
					i++
					if i > 5 {
						break L
					}
				}
			}
	*/
	ch := make(chan int)
	go shower(ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	x := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(<-x)
	}
}

func shower(c <-chan int) {
	for {
		j := <-c
		fmt.Printf("%d\n", j)
	}
}
func dup3(in <-chan int) (<-chan int, <-chan int, <-chan int) {
	a, b, c := make(chan int, 2), make(chan int, 2), make(chan int, 2)
	go func() {
		for {
			x := <-in
			a <- x
			b <- x
			c <- x
		}
	}()
	return a, b, c
}
func fib() <-chan int {
	x := make(chan int, 2)
	a, b, out := dup3(x)
	go func() {
		x <- 0
		x <- 1
		<-a
		for {
			x <- <-a + <-b
		}
	}()
	return out
}
