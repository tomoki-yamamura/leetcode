package main

import (
	"fmt"

	"github.com/tomoki-yamamura/go-docker/array"
)

func main() {
	result := array.TwoSum([]int{3,2,3}, 6)
	fmt.Println(result)
}
