package main

import (
	"fmt"
	"sort"
)


type IntNum []int
func (i IntNum) Len() int{
	fmt.Println(i)
	return len(i)
}

func (d IntNum) Less(i, j int ) bool  {
	return d[i] > d[j]
}
func (d IntNum) Swap(i, j int){

	d[i] , d[j] = d[j], d[i]
}

type Courses []CourseInfo

type CourseInfo struct {
	Name string
	Price int
	Url string
}

func (d Courses) Len() int {
	return len(d)
}
func (d Courses) Less(i, j int) bool {
	return d[i].Price < d[j].Price
}
func (d Courses) Swap(i, j int)  {
	d[i], d[j] = d[j], d[i]
}

func main() {

	data := IntNum{1,4,7,2,5}
	sort.Sort(data)
	for _, v := range data{
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	courses := Courses{
		CourseInfo{"python",100,""},
		CourseInfo{"go",300,""},
		CourseInfo{"sanic",200,""},
		CourseInfo{"java",400,""},
	}
	sort.Sort(courses)
	for _, v := range courses{
		fmt.Println(v)
	}

}
