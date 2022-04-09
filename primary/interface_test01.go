package main

import "fmt"

type Pythone struct {

}

func (p Pythone) GoCoder() string {
	panic("implement me")
}

type Goer struct {

}

type GoProgramer interface {
	GoCoder() string
}
type Pythoner interface {
	Python() string
}

type Programmer interface {
	GoProgramer
	Pythoner
}

func (g * Goer) GoCoder() string  {
	fmt.Println("我是go 开发者")
	return "go开发者"
}

func (p * Pythone) Python() string{
	fmt.Println("我是Python 开发者")
	return "python开发者"
}


func main() {
	//var p Programmer = &Pythoner{}
	//p.Python()
	//p := &Pythoner{}
	//p.Python()
	//p := Goer{}
	//p.GoCoder()
	var g GoProgramer
	g = &Goer{}

	g.GoCoder()

	var p Programmer
	p = &Pythone{}
	p.Python()


}