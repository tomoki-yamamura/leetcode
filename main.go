package main

import (
	"fmt"

	"github.com/tomoki-yamamura/go-docker/dp"
	// "github.com/tomoki-yamamura/go-docker/array"
)

func main() {
	coins := []int{2,4,6}
	result := dp.CoinChange(coins, 8)
	fmt.Println(result)
}
