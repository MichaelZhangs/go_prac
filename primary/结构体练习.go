package main

import "fmt"

type Student struct {
	id    int
	name  string
	score [3]int
}
func count(arr []Student){
	// 算总成绩
	for i:=0; i< len(arr); i++{
		//fmt.Println(arr[i].score)
		sum := 0
		for j:=0; j<len(arr[i].score); j++{
			sum += arr[i].score[j]
		}
		fmt.Printf("%d 学号 %s 总成绩 %d 平均成绩 %d\n", arr[i].id, arr[i].name, sum, sum / 3)
	}

}

func main() {
	arr:=[]Student{
		{101, "小明",[3]int{68,72,88}},
		{102,"小红",[3]int{72,88,69}},
		{103,"小白", [3]int{62,58,74}},
	}
	count(arr)
	fmt.Println(arr)
}
