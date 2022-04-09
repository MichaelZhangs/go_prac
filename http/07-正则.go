package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main()  {
	str := "abc a7c msdf cat 09sd cb%a"
	reg := regexp.MustCompile(`a[0-9]c`)
	res := reg.FindAllStringSubmatch(str, -1)
	fmt.Println(res)
	s := []string{"1232", "sfsf"}
	a := strings.Join(s, "####")
	fmt.Println(a)

}
