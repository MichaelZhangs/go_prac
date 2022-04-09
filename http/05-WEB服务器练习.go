package main

import (
	"fmt"
	"net/http"
	"os"
)

func OpenSendFile(fname string, w http.ResponseWriter )  {
	pathName := "D:/爬虫代码" + fname
	f, err := os.Open(pathName)
	if err != nil{
		fmt.Println("err : ", err)
		w.Write([]byte("no such file or dictory ...."))
		return
	}
	buf := make([]byte, 4096)
	for {
		n, _ := f.Read(buf)
		if n == 0{
			return
		}
		w.Write(buf[:n]) // 写给客户端
	}

}

func han(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("客户端请求: ", r.URL)
	OpenSendFile(r.URL.String(), w)

}

func main()  {
	http.HandleFunc("/", han)
	http.ListenAndServe("127.0.0.1:8000", nil)
}
