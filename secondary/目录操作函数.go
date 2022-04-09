package main

import (
	"fmt"
	"os"
)

func main()  {

	fmt.Println("请输入待查询的目录: ")
	var path string
	fmt.Scan(&path)
	// 打开目录
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil{
		fmt.Println("打开失败")
		return
	}
	defer f.Close()

	//  读取目录
	info , err:= f.ReadDir(-1) // -1 表示读取目录所有文件
	// 遍历返回的切片
	for _, fileInfo := range info{
		if fileInfo.IsDir(){
			fmt.Println(fileInfo.Name(), " 是目录")
		}else {
			fmt.Println(fileInfo.Name(), " 是文件")
		}

	}

}
