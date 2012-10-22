// funcdemo project main.go
package main

import (
	"fmt"
)

func main() {
	a := []byte{'1', '2', '3', '4'}
	var x int
	for i := 0; i < len(a); {
		x, i = nextInt(a, i)
		fmt.Println(x)
	}

	fmt.Println(f())

	tt := func() { //函数做值
		fmt.Println("Hello")
	}
	tt()

	var xs = map[int]func() int{ //使用函数作为map的value
		1: func() int { return 10 },
		2: func() int { return 20 },
		3: func() int { return 30 },
	}
	fmt.Println(xs[1]()) //y调用map中的value函数

	callback(10, printit)

	//ff := make([...]float64{23.32, 234.45, 352, 634, 223.45})
	//fmt.Println(practiseOne4Two(ff))
	x, y := practiseQ7(7, 2)
	fmt.Println(x, y)

	var st stack
	st.push(12)
	st.push(13)
	fmt.Printf("stack %v\n", st)
	fmt.Println(st.pop())
	fmt.Printf("stack %v\n", st)
}

func nextInt(b []byte, i int) (int, int) {
	x := 0
	for ; i < len(b); i++ {
		x = x*10 + int(b[i]) - '0'
	}
	return x, i
}

func f() (ret int) { //初始化为0
	defer func() {
		ret++ //ret增加1
	}()
	return 0 //返回的是1而不是0
}

func printit(x int) { //定义函数
	fmt.Printf("%d\n", x)
}

func callback(y int, f func(int)) { //f 将会保存函数
	f(y) //调用回调函数，f输入变量y
}

func throwsPanic(f func()) (b bool) { //定义一个新函数接受一个函数作为参数

	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f()
	return
}

//计算类型为floatt64的slice的平均值
func practiseOne4Two(f []float64) float64 {

	sum := 0.0
	switch len(f) {
	case 0:
		return 0
	default:
		for _, va := range f {
			sum += va
		}
	}
	return sum / float64(len(f))

}

func practiseQ7(x int, y int) (xx int, yy int) {

	return y, x
}

type stack struct {
	i    int
	data [10]int
}

func practiseQ91(x int, y int) (xx int, yy int) {

	return y, x
}

func (x *stack) push(k int) {

	x.data[x.i] = k
	x.i++
}

func (x *stack) pop() int {
	x.i--
	return x.data[x.i]
}
