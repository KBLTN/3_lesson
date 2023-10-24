package main

import "fmt"

type Point struct {
	X int
	Y int
	S string
}

func (p Point) method() {
	fmt.Println(p.X)
	fmt.Println(p.Y)
	fmt.Println(p.S)
}

func main() {
	p1 := Point{
		X: 1,
		Y: 2,
		S: "hi",
	}
	p2 := &p1
	p2.method()
}

func structs() {
	p1 := Point{
		X: 1,
		Y: 2,
		S: "A",
	}
	fmt.Println(p1)
	fmt.Println(p1.X)
	p1.X = 123
	fmt.Println(p1)

	p2 := Point{
		X: 123,
	}
	fmt.Println(p2)

	g := &p1
	fmt.Println(g)
	fmt.Println(*g)
	fmt.Println(&g)
}
