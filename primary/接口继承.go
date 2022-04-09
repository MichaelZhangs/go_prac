package main

import "fmt"

type Humaner interface {
	SayHi()
}

type Personer interface {
	Humaner
	Sing()
}
type student3 struct {

	name string
	sex string
	age int
}



func (s * student3)SayHi(){
	fmt.Printf("你好。我是%s, 今年%d\n", s.name, s.age)
}
func (s * student3)Sing(){
	fmt.Printf("你好。我是%s, 今年%d, 爱唱歌\n", s.name, s.age)
}

func a(p Personer){
	p.SayHi()
	p.Sing()

}
func b(h Humaner){
	h.SayHi()
}
func main(){
	var h Personer
	h = &student3{"张三","男", 18}
	a(h)
	b(h)
}