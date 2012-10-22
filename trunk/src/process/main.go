// process project main.go
package main

import (
	"fmt"
)

func main() {
	var p *int

	fmt.Printf("%v\n", p)

	var i int
	p = &i

	fmt.Printf("%v\n", p)

	var q *string
	var str string
	q = &str

	*q = "string"

	fmt.Println(str)
}
