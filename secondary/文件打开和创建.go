package main

import (
	"fmt"
	"os"
)

func main01()  {
	f, err := os.Create("./a.txt")
	if err != nil{
		fmt.Println(" 创建文件失败")
	}
	defer f.Close()
	fmt.Println(err)
	f.WriteString("hello \n how are u")
}
func main02(){
	f, err := os.Open("./a.txt")
	if err != nil{
		fmt.Println(" 打开文件失败")
	}
	defer f.Close()
	 _, err = f.WriteString("123233")
	if err != nil{
		fmt.Println("写入失败")
	}
	fmt.Println(err)
}
func main(){
	f, err := os.OpenFile("./a.txt",os.O_RDWR ,6)
	if err != nil{
		fmt.Println(" 打开文件失败")
	}
	defer f.Close()
	 _, err = f.WriteString("121sadfs22342\n")
	 fmt.Println(err)
}