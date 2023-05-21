package dp

import "math"

// could not!!!
// func CoinChange(coins []int, amount int) int {
// 	dp := make(map[int]int)
// 	for _, num := range coins {
// 		dp[num] = 1
// 	}
// 	if coins[0] > amount {
// 		return -1
// 	}
// 	for i := coins[0]; i <= amount; i++ {
// 		_, ok := dp[i]
// 		if ok {
// 			continue
// 		}
// 		left := 1
// 		right := i-1
// 		dp[i] = int(math.MaxInt64)
// 		for left < right {
// 			if !(dp[left] != 0 && dp[right] != 0) {
// 				return -1
// 			}
// 			tmp := dp[left] + dp[right]
// 			dp[i] = min(tmp, dp[i])
// 			right--
// 			left++
// 		}
// 	}
// 	return dp[amount]
// }

// func min(a, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }

func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ { dp[i] = math.MaxInt32 }
	
	for _, c := range coins {
			for a := 1; a <= amount; a++ {
					if a >= c {
							dp[a] = Min(dp[a], 1+dp[a-c])
					}
			}
	}
	
	if dp[amount] == math.MaxInt32 { return -1 }
	return dp[amount]
}

func Min(a, b int) int { if a < b { return a }; return b }