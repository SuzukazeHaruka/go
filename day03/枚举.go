package main

import "fmt"

func main() {
	const  a int  =10
	const b =11.2
	const str="你好世界"
	fmt.Println(a,b,str)
	const(
		c=iota
		d=iota
		f=iota
	)

}
