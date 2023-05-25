package main

import (
	"fmt"

	"github.com/tomoki-yamamura/go-docker/dp"
	// "github.com/tomoki-yamamura/go-docker/array"
)

func main() {
	s := "leetcode"
	wordDict := []string{"leet","code","neet"}
	// s := "catsandog"
	// wordDict := []string{"cats","dog","sand","and","cat"}
	result := dp.WordBreak(s, wordDict)
	fmt.Println(result)
}
