package main

import (
	"fmt"

	"github.com/tomoki-yamamura/go-docker/dp"
	// "github.com/tomoki-yamamura/go-docker/array"
)

func main() {
	result := dp.CombinationSum4([]int{1,2,3}, 4)
	fmt.Println(result)
}
