package main

import "fmt"

func main0703(){
	//var 变量名 数据类型
	//数组定义和使用
	//var 数组名[元素个数]数据类型
	//var score[820000]int
	var arr[10] int //数组声明后默认的值为0
	//可以通过 数组名[下标] 访问具体的元素
	//数组下标是从0开始的到数组元素最大个数-1为最大小标
	//为数组赋值
	arr[0]=10
	arr[1]=20
	arr[2]=30
	//打印数组元素
	fmt.Println(arr[0])


}

func main0709() {
	var arr[10] int=[10]int {1:123,3:456,6:789}
	fmt.Println(len(arr))
	for i,v:=range  arr{
		fmt.Printf("下标%d,值为%d\n",i,v)
	}
}

func main0708() {
	var arr[5]int =[5]int {1,2,3,4}
	var arr1[5]int
	//数组下标越界
	//arr[5]=5
	//arr[-1]
	//数组名 表示数组的首地址 不能将值赋值给数组名
	//arr-1
	//数组如果类型相同可以赋值
	arr1=arr
	fmt.Println(arr1)
}

func main(){
	var arr [5]int =[5]int{1,2,3,4,5}
	//数组元素在内存单元中连续存储的
	for i:=0;i<len(arr) ;i++  {
		fmt.Printf("%p\n",&arr[i])
	}
	//数组名代表数组元素的首地址
	fmt.Printf("%p\n",&arr)
}
