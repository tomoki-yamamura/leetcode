package main

import (
	"fmt"

	"github.com/tomoki-yamamura/go-docker/dp"
	// "github.com/tomoki-yamamura/go-docker/array"
)

func main() {
	nums := []int{1,2,3}
	result := dp.Rob2(nums)
	fmt.Println(result)
}
