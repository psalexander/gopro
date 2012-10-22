// testeven project main.go
package main

import (
	"even"
	"fmt"
)

func main() {
	if even.Even(3) {
		fmt.Println("3 is Even!")
	} else {
		fmt.Println("3 is not Even!")
	}
	fmt.Println("Hello World!")
	str := "this is a string"

	fmt.Printf("%+v", str) //默认格式的值

	fmt.Printf("%#v", str) //go样式的值表达（这里为str加上了双引号）

	fmt.Printf("%T", str) //输出表达式的类型
}
