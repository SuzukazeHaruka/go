package main

import "fmt"

func test7(a int){
	if(a==0){
		return
	}
	fmt.Println(a)
	test7(a-1)
}



var sum=1
func mlt(num int){
	if(num==0){
		return
	}
	mlt(num-1)
	sum*=num
}
func main() {
	//test7(18)
	mlt(8)
	fmt.Println(sum)
}


