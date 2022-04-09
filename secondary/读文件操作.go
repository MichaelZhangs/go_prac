package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unsafe"
)

func main()  {
	f , err := os.OpenFile("./a.txt", os.O_RDWR, 6)
	if err != nil{
		fmt.Println("打开失败")
		return
	}
	defer f.Close()
	// 创建一个缓冲区的reader
	r := bufio.NewReader(f)
	fmt.Println(unsafe.Sizeof(r))
	if err != nil{
		fmt.Println("读取失败")
		return
	}
	for {
		buf, err := r.ReadBytes('\n')
		if err != nil && err == io.EOF{
			fmt.Println("文件读取完毕")
			return
		}else if err != nil {
			fmt.Println("readBytes err : ", err)
		}
		fmt.Print(string(buf))
	}

}

