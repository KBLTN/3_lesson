package main

import "fmt"

func main() {
	//arrays
	var a [2]string
	a[0] = "hi"
	a[1] = "Kotya"
	// a[2] = "ok" // - error, invalid array index
	fmt.Println(a)
	fmt.Println(a[0])

	numbers := [...]int{1, 2, 3}
	numbers[2] = 4

}
