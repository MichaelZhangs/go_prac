package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	// ”用户名:密码@("主机名:端口号")/数据库名“
	db, _ := sql.Open("mysql", "root:123456@(127.0.0.1:3306)/db1?charset=utf8")
	defer db.Close()
	err := db.Ping()
	if err != nil{
		fmt.Println("数据库连接失败")
		return
	}
	exec := "select * from student"
	result,_ := db.Query(exec)
	fmt.Println(result)
	for result.Next(){
		var id, name, sex , age string
		//var id int

		if err := result.Scan(&id,&name, &age, &sex); err != nil {
			fmt.Println(err)
		}
		fmt.Println(name, age, sex)
	}

}
