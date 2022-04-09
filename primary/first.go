package main 
//fmt format 包含格式化输入输出
import (
	"fmt"
	"time"
)

//注释  注释不参与程序编译 可以帮助可以程序
//行注释 只能注释一行
//快速注释的快捷键 ctrl+/
/*
块注释
可以注释多行
*/

//main 叫主函数 是程序的主入口 程序有且只有一个主函数

func main01() {
	//在屏幕打印hello world
	fmt.Println("hello world")
	fmt.Println("传智播客")

}
func main(){
	fmt.Println("这是第一个程序")
	main01()
	for i:=0; i < 10; i++{
		go func( n int ) {
				fmt.Println(n)
				//time.Sleep(time.Millisecond)
		}(i)
	}
	time.Sleep(time.Second)
}