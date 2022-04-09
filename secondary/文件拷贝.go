package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	f , err := os.OpenFile("./src/a.txt", os.O_RDWR, 6)
	if err != nil{
		fmt.Println("打开失败")
		return
	}
	dstf,_ := os.Create("./abak.txt")
	defer f.Close()
	defer dstf.Close()
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		fmt.Println("n = ", n)
		if err != nil && err != io.EOF{
			fmt.Println(err)
			return
		}
		if n == 0{
			fmt.Println("文件处理完毕")
			break
		}
		tmp := buf[:n]
		dstf.Write(tmp)

	}

}
