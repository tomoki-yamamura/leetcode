package main

import (
	"fmt"

	"github.com/tomoki-yamamura/go-docker/dp"
	// "github.com/tomoki-yamamura/go-docker/array"
)

func main() {
	nums := []int{1,2,3}
	target := 4
	result := dp.CombinationSum4(nums, target)
	fmt.Println(result)
}
