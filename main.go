package main

import (
	"fmt"

	"github.com/tomoki-yamamura/go-docker/array"
)

func main() {
	result := array.MaxProfit([]int{7,6,4, 3, 1})
	fmt.Println(result)
}
