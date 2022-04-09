package main

import (
	"fmt"
	"net/http"
)

func myhandler(w http.ResponseWriter, r *http.Request)  {

	// 写给客户端
	w.Write([]byte("welcome login......"))
	fmt.Println("Header: ",r.Header)
	fmt.Println("URL: ", r.URL)
	fmt.Println("Method: ", r.Method)
	fmt.Println("Host: ", r.Host)
	fmt.Println("body: ", r.Body)
	fmt.Println("RemoteAddr: ", r.RemoteAddr)
	fmt.Println("=====\n")
	for k,v := range r.Header{
		fmt.Println(k,": ", v[0])
	}
}

func main()  {
	// 注册回调 函数
	http.HandleFunc("/login.html", myhandler)
	// 绑定地址
	http.ListenAndServe("127.0.0.1:8000", nil)
}
