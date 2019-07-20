package main

import "fmt"


type FUNCTEST func() int

func test6() func() int{
	var i int
	return func() int{
		i++
		return i
	}
}

func main() {
	for i:=0;i<10 ;i++  {
		//	函数结束会从栈区2内存销毁
		//test4()
		//
		//v是一个函数类型的变量
		v:=test6()
		//函数调用
		//
		for i:=0;i<10 ;i++  {
			fmt.Println(v())
		}
	}
}