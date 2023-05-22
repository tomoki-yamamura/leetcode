package main

import (
	"fmt"

	"github.com/tomoki-yamamura/go-docker/dp"
	// "github.com/tomoki-yamamura/go-docker/array"
)

func main() {
	coins := []int{4,10,4,3,8,9}
	result := dp.LengthOfLIS(coins)
	fmt.Println(result)
}
