package main

import "fmt"

func main() {
	a:=10
	b:=20
	defer func(a int,b int) {
		fmt.Println("匿名函数中a",a)
		fmt.Println("匿名函数中b",b)
	}(a,b)
	a=100
	b=200
	fmt.Println("main函数中a",a)
	fmt.Println("main函数中b",b)
}


