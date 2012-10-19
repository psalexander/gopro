// test project main.go
package main

import (
	"fmt"
	"mypak"
)

func main() {
	fmt.Println("Hello World!", mypak.Test)
	fmt.Print("test")
	fmt.Printf("%s\n", mypak.Test)

	x := 10
	if x > 0 {
		fmt.Print("x>0")
	} else {
		fmt.Print("x<0")
	}
	myfunc3()
}

func myfunc() {
	i := 0
Here:
	fmt.Println(i)
	i++
	if i == 50000 {

	} else {
		goto Here
	}
}
func myfunc2() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	at := "this is a string"
	astr := []rune(at)
	for i, j := 0, len(astr)-1; i < j; i, j = i+1, j-1 {
		astr[i], astr[j] = astr[j], astr[i]
	}
	fmt.Println(string(astr))
}

func myfunc3() {
	list := []string{"a", "b", "c", "d", "e", "f"}
	for k, v := range list {
		fmt.Printf("%d  =   %s\n", k, v)
	}
}

func myfun4() {
	for pos,char:=range ""
}
