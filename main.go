package main

import (
	"fmt"

	"github.com/tomoki-yamamura/go-docker/array"
	// "github.com/tomoki-yamamura/go-docker/array"
)

func main() {
	arg := []int{1,8,6,2,5,4,8,3,7}
	// arg := []int{4,5,6,7,8,9,10,11,12,13,0,1,2}
	result := array.MaxArea(arg)
	fmt.Println(result)
}
