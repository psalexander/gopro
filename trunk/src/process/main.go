// process project main.go
package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	fmt.Println("length:", flag.NArg())
	var a NameAge
	var i int
	d := &i
	c := &a
	fmt.Println(d)
	fmt.Println(c)
	fmt.Printf("222%s", a.getName())
}
