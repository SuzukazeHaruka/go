package main

import "fmt"

func main() {
	var scores[10]int
	//通过键盘输入学生成绩 计算总分和平均分
	//fmt.Scan(&scores)//err
	for i:=0;i<len(scores) ;i++  {
		fmt.Scan(&scores[i])
	}
	sum:=10
	for _,v:=range scores{
		sum+=v
	}

	fmt.Printf("总成绩为：%d,平均成绩为：%d\n", sum, sum/len(scores))
	fmt.Println(scores)
}
