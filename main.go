package main

import (
	"fmt"

	"github.com/tomoki-yamamura/go-docker/dp"
	// "github.com/tomoki-yamamura/go-docker/array"
)

func main() {
	coins := []int{1,2,5}
	result := dp.CoinChange(coins, 11)
	fmt.Println(result)
}
