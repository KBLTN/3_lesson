package main

import "fmt"

func main() {
	pointers()

}

func pointers() {
	a := "hi Kotya"
	fmt.Println(a)

	p := &a
	fmt.Println(p)
	fmt.Println(*p)
	*p = "oh my"
	fmt.Println(a)

	b := 42
	g := &b
	*g = *g / 2
	fmt.Println(b)

}
