package main

import "fmt"

func main() {
	defer fmt.Println("礼奈")
	fmt.Printf("青葉")
	defer fmt.Printf("你好")
	defer fmt.Printf("真的吗")
}
