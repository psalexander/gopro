// interface project main.go
package main

import (
	"fmt"
	"reflect"
)

type S struct {
	i int
}

func (s *S) Get() int  { return s.i }
func (s *S) Put(v int) { s.i = v }

type I interface {
	Get() int
	Put(int)
}

func f(p I) {
	fmt.Println(p.Get())
	p.Put(10)
	fmt.Println(p.Get())
}
func main() {

	var s S
	f(&s)
	fmt.Println(g(&s))

	//var buf bytes.Buffer //创建指向接口的指针是无意义的
	//io.Copy(buf, os.Stdin)

	ints := Xi{44, 67, 3, 17, 89, 10, 73, 9, 14, 8}
	strings := Xs{"nut", "ape", "elephant", "zoo", "go"}
	Sort(ints)
	fmt.Printf("%v\n", ints)
	Sort(strings)
	fmt.Printf("%v\n", strings)

	p1 := new(Person)
	Set(p1)
	fmt.Println(p1.Name)
}

type R struct {
	i int
}

func (r *R) Get() int  { return r.i }
func (r *R) Put(i int) { r.i = i }

func ft(p I) {
	switch p.(type) { //p.(I)
	case *R:
		fmt.Println("*R")
	case *S:
		fmt.Println("*S")
	//case R:fmt.Println("R") R
	//case S:fmt.Println("S")
	default:
		fmt.Println("no type")
	}
}

func g(something interface{}) int {
	return something.(I).Get()
}

//会有运行时错误（ cannot define new methods on non-local type int）
//func (i int) Emit() {
//	fmt.Println("%d", i)
//}
//会有编译时错误（ expected (unqualified) identifier）
//func (a *net.AddrError) Emit(){
//}

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type Xi []int
type Xs []string

func (p Xi) Len() int {
	return len(p)
}

func (p Xi) Less(i, j int) bool {
	return p[j] < p[i]
}

func (p Xi) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Xs) Len() int {
	return len(p)
}

func (p Xs) Less(i, j int) bool {
	return p[j] < p[i]
}

func (p Xs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func Sort(x Sorter) {
	for i := 0; i < x.Len()-1; i++ {
		for j := i; j < x.Len(); j++ {
			if x.Less(i, j) {
				x.Swap(i, j)
			}
		}
	}

}

type Person struct {
	Name string "namestr"
	Age  int
}

func showTag(i interface{}) {
	switch t := reflect.TypeOf(i); t.Kind() { //不懂t.Kind()的用法？？？？？？
	case reflect.Ptr:
		tag := t.Elem().Field(0).Tag // 使用Elem() 得到了指针指向的值。
		fmt.Println(tag)
	}
}

func show(i interface{}) {
	switch i.(type) {
	case *Person:
		t := reflect.TypeOf(i)
		v := reflect.ValueOf(i)
		tag := t.Elem().Field(0).Tag
		name := v.Elem().Field(0).String()
		fmt.Println(tag)
		fmt.Println("----" + name)
	}
}

func Set(i interface{}) {
	switch i.(type) {
	case *Person:
		r := reflect.ValueOf(i)
		r.Elem().Field(0).SetString("Albert Einstein")
	}
}
