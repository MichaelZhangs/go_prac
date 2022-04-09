package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

func main(){
	d_path := path.Dir("./")
	f_path := path.Join(d_path, "/file02")
	fp, err := os.Open(f_path)
	if err != nil{
		fmt.Println("文件打开失败")
	}

	// 关闭文件
	defer fp.Close()
	// 创建文件缓冲区
	r := bufio.NewReader(fp)
	// 读取一行内容
	b, t, _ := r.ReadLine()
	fmt.Println(string(b), t)
}
