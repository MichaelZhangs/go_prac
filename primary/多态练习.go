package main

import (
	"fmt"
)

//接口实现
type USBer interface {
	Read()
	Write()
}

// 创建对象
type USBDev struct {
	id int
	name string
	rspeed int
	wspeed int
}

type Mobile struct {
	USBDev
	call string
}
type Upan struct {
	USBDev
}

func (m * Mobile) Read(){
	fmt.Printf("%s 手机正在读取数据速度为 %d\n", m.name, m.rspeed)
}

func (m *Mobile) Write(){
	fmt.Printf("%s 手机正在写入数据速度为 %d\n", m.name, m.wspeed)
}

func (u *Upan) Read(){
	fmt.Printf("%s u盘正在读取数据速度为 %d\n", u.name, u.rspeed)
}
func (u *Upan) Write(){
	fmt.Printf("%s u盘正在写入数据速度为 %d\n", u.name, u.wspeed)
}

// 定义多态
func a(u USBer){
	u.Write()
	u.Read()
}

func main(){
	// 接口类型
	var u USBer
	u  = &Mobile{USBDev{1,"小米", 128, 64},"你好"}
	a(u)
	u = &Upan{USBDev{2,"dz",256,256}}
	a(u)
}