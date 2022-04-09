package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request)  {
	// w 写回给客户端
	//buf := make([]byte, 4096)
	w.Write([]byte("Hello world, zzzz"))
}

func main()  {
	// 注册回调函数，该回调函数会在服务器被访问时，自动被调用
	http.HandleFunc("/itcast", handler)
	// 绑定服务器监听地址
	http.ListenAndServe("127.0.0.1:8000", nil)



}
