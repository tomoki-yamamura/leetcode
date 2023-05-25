package dp



// couldnot!!!!!!!!
// func WordBreak(s string, wordDict []string) bool {
// 	dp := make([]bool, len(s)+1)
// 	dp[0] = true
// 	for i := range s {
// 		if !dp[i] {
// 			continue
// 		}
// 		for _, word	:= range wordDict {
// 			if i + len(word) > len(s) {
// 				continue
// 			}
// 			if s[i:i+len(word)] == word {
// 				dp[i+len(word)] = true
// 			}
// 		}
// 	}
// 	return dp[len(s)]
// }

func WordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[len(s)] = true

	for i := len(s) - 1; i >= 0; i-- {
		for _, w := range wordDict {
			if (i + len(w)) <= len(s) && s[i:i+len(w)] == w {
				dp[i] = dp[i+len(w)]
			}
			if dp[i] {
				break
			}
		}
	}

	return dp[0]
}