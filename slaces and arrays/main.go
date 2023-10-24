package main

import "fmt"

func main() {
	//arrays and slices
	animalsArr := [4]string{
		"dog",
		"cat",
		"giraffe",
		"lephant",
	}

	var a []string = animalsArr[0:2]
	fmt.Println(a)
	var b []string = animalsArr[1:]
	fmt.Println(b)

	b[0] = "123"
	fmt.Println(a)
	fmt.Println(animalsArr)

}
