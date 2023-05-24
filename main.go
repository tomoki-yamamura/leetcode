package main

import (
	"fmt"

	"github.com/tomoki-yamamura/go-docker/dp"
	// "github.com/tomoki-yamamura/go-docker/array"
)

func main() {
	coins := []int{7,7,7,7,7,7,7}
	result := dp.LengthOfLIS(coins)
	fmt.Println(result)
}
