package main

import (
	"fmt"

	"github.com/tomoki-yamamura/go-docker/array"
	// "github.com/tomoki-yamamura/go-docker/array"
)

func main() {
	arg := []int{7,0,1,2,3,4,5,6}
	// arg := []int{4,5,6,7,8,9,10,11,12,13,0,1,2}
	result := array.Search(arg, 8)
	fmt.Println(result)
}
