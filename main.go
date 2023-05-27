package main

import (
	"fmt"

	"github.com/tomoki-yamamura/go-docker/dp"
	// "github.com/tomoki-yamamura/go-docker/array"
)

func main() {
	nums := []int{2,1,1,2}
	result := dp.Rob(nums)
	fmt.Println(result)
}
