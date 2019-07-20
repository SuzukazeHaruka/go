package main

import "fmt"

func main0401() {
	var i int
	for i=0;i<10 ;i++  {
		fmt.Println(i)
	}
}

//将函数类型作为函数参数
func test5(f func(int,int)){
	f(10,20)
}
func main0402(){
	a:=10
	b:=20
	//匿名函数 在代码去进行存储
	f:=func(a int , b int){
		fmt.Println(a+b)
	}
	f(a,b)
	//f(1,2)
	test5(f)

}

func main() {
	a:=10
	b:=20
	//匿名函数如果在{}后面有()表示函数调用
	 value:=func(a int , b int)int{
		 return a + b
	}
	v:=value(a,b)
	fmt.Printf("%T\n",value)
	fmt.Printf("%T\n",v)

}
