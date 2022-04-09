package main

import (
	"encoding/json"
	"fmt"
)

type Info struct {
	Name string `json:"name"`
	Age int `json:"age,omitempty"`
	Gender string `json:"gender"`
}

func main() {
	info := Info{Name:"bobby", Gender: "男"}
	res, _ := json.Marshal(info)
	fmt.Println(string(res))

	//t := reflect.TypeOf(info)
	//fmt.Println(t.Name())
	//for i:= 0; i< t.NumField(); i++{
	//	field := t.Field(i)  // 获取结构体的每个字段
 	//	fmt.Println(field)
	//	name := field.Tag.Get("json")
	//	fmt.Println(name)
	//}
}