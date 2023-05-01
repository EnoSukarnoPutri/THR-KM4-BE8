package main

import "fmt"

func AddElement(data []int, element int, position string) []int {
	if position == "up" {
		return append([]int{element}, data...)
	} else if position == "down" {
		return append(data, element)
	}
	return data
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	element := 6
	position := "up"

	newData := AddElement(data, element, position)
	fmt.Println(newData)

	position = "down"
	newData = AddElement(data, element, position)
	fmt.Println(newData)
}
