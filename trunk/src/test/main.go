// test project main.go
package main

import (
	"fmt"
	//"mypak"
)

func main() {
	sliceTest()
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

}

func sliceTest() {
	d := make([]int, 10)
	a := [...]int{1, 2, 3, 4, 5}
	b := a[2:4]
	b1 := a[1:3]
	c := a[1:]
	d1 := a[:4]
	fmt.Println(len(d1))
	fmt.Println(cap(b1))
	fmt.Println("c", c)
	e := append(d, b...)
	f := append(d, 2, 3, 5)
	fmt.Println("e", e)
	fmt.Println(f)

	fmt.Println(cap(c))

	var s = make([]int, 6)
	fmt.Println(s)
	n1 := copy(c, b[0:3])
	fmt.Println(n1)
	fmt.Println(c)
}

func mapdemo() {
	monthdays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31, // 逗号是必须的
	}

	fmt.Println(monthdays)
	years := 0
	for _, days := range monthdays {
		years += days
	}
	fmt.Println(years)

	monthdays["undecl"] = 30
	monthdays["Feb"] = 29
	fmt.Println(monthdays)
	fmt.Println(monthdays["test"])

	var value int
	var present bool
	value, present = monthdays["Fed"]
	fmt.Printf("monthdays['Fed'] 's present is %b, it's value=%d\n", present, value)

	v, ok := monthdays["Sep"]
	fmt.Println(v, ok)

	delete(monthdays, "undecl")

	fmt.Println(monthdays)
}

func practiceOne4One() {
	//for
	for j := 0; j < 10; j++ { //captor1,1
		fmt.Printf("j=%d\n", j)
	}
	//goto
	var i int = 0
flag:
	if i < 10 {
		fmt.Printf("i=%d\n", i)
		i++
		goto flag
	}
	//range
	ary := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for fg := range ary {
		fmt.Printf("fg=%d\n", fg)
	}
}

func practiceOne4Two() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}

	}
}

func practiceOne4Three() {
	for i := 1; i <= 100; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("A")
		}
		fmt.Print("\n")
	}
}

func practiceOne4Four() {
	s := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("S 's length is %d\n", len(s))
	for _, str := range s {
		fmt.Print(str)
	}
}

func practiceOne4Five() {
	s := "asSASA ddd dsjkdsjs dk"
	strMap := make(map[string]int)
	for _, str := range s {
		v, ok := strMap[string(str)]
		if ok {
			strMap[string(str)] = v + 1
		} else {
			strMap[string(str)] = 1
		}
	}

	fmt.Println(strMap)
}
