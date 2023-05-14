package main

import (
	"fmt"

	"github.com/tomoki-yamamura/go-docker/array"
	// "github.com/tomoki-yamamura/go-docker/array"
)

func main() {
	arg := []int{3,4,5,1,2}
	result := array.FindMin(arg)
	fmt.Println(result)
}
