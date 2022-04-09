package main

import (
	"fmt"
	"os"

	//"os"
	"path"
)

func main01(){
	d_path := path.Dir("./")
	f_path := path.Join(d_path, "/file")
	fmt.Println(d_path)
	fmt.Println(f_path)

//	f, err := os.Create(f_path)
//	if err!= nil{
//		fmt.Println()
//	}
//	f.WriteString("11111")
//	f.Close()
//
//
}
func main(){
	d_path := path.Dir("./")
	f_path := path.Join(d_path, "/file02")
	fmt.Println(d_path)
	fmt.Println(f_path)

	f, err := os.Create(f_path)
	if err!= nil{
		fmt.Println()
	}
	// 关闭文件
	defer f.Close()
	s := "1234455, 你是\r\n好啊\r\n哈哈哈哈\r\n是的"
	b := []byte(s)
	n , _ := f.Write(b)
	fmt.Println(n)
}