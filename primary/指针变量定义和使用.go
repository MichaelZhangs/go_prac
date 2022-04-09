package main

import "fmt"

func  main()  {
	var a int = 10
	var p *int
	p = &a
	*p = 100
	fmt.Println(*p, a)
	fmt.Printf("%p %p\n", p, &a)
	var b * int
	b = new(int)
	fmt.Println(b)
	fmt.Println(*b)
}
