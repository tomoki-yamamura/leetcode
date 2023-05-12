package main

import (
	"fmt"

	"github.com/tomoki-yamamura/go-docker/array"
)

func main() {
	arg := []int{7,6,4, 3, 1}
	result := array.Reverse(arg)
	fmt.Println(result)
}
