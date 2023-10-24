package main

import "fmt"

func main() {
	//slices

	createSlices := make([]string, 3)
	fmt.Println(fmt.Sprintf("len: %d", len(createSlices)))
	fmt.Println(fmt.Sprintf("cap: %d", cap(createSlices)))
	createSlices[0] = "1"
	createSlices[1] = "2"
	createSlices[2] = "3"
	createSlices = append(createSlices, "4")
	fmt.Println(createSlices)
	fmt.Println(fmt.Sprintf("len: %d", len(createSlices)))
	fmt.Println(fmt.Sprintf("cap: %d", cap(createSlices)))
	createSlices = append(createSlices, "4", "5", "7")
	fmt.Println(createSlices)
	fmt.Println(fmt.Sprintf("len: %d", len(createSlices)))
	fmt.Println(fmt.Sprintf("cap: %d", cap(createSlices)))
}
