package main

import "fmt"

func main() {
	lists := []int{1,2,3,4}
	length := len(lists)
	fmt.Println(lists[0:length/3])
	fmt.Println(length/3)
}