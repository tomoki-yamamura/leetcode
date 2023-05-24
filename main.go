package main

import (
	"fmt"

	"github.com/tomoki-yamamura/go-docker/dp"
	// "github.com/tomoki-yamamura/go-docker/array"
)

func main() {
	result := dp.LongestCommonSubsequence("abcde", "ace")
	fmt.Println(result)
}
