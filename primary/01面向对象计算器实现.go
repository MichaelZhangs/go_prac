package main

import "fmt"

type Opt struct {
	num1 int
	num2 int
}

type AddOpt struct {
	Opt
}
type SubOpt struct {
	Opt
}


func (add * AddOpt)Operate()int{
	return add.num1 + add.num2
}
func (add * SubOpt)Operate()int{
	return add.num1 - add.num2
}



type Opter interface {
	Operate()int
}

func main(){
	//var add AddOpt
	//add.num2 = 10
	//add.num1 = 20
	//s := add.Operate()
	//var sub SubOpt
	//sub.num2 = 20
	//sub.num1 = 10
	//value := sub.Operate()
	//fmt.Println(s)
	//fmt.Println(value)
	var o Opter
	o = &AddOpt{Opt{10,20}}
	fmt.Println(o.Operate())
	o.Operate()
	o = &SubOpt{Opt{10,20}}
	fmt.Println(o.Operate())
}
