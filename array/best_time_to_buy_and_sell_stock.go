package array

func MaxProfit(prices []int) int {
	var tmp int
	var result int
	left := 0
	right := 0
	for i := 0; i < len(prices)-1; i++ {
		right++
		if prices[right] < prices[left] {
			left = right
			continue
		}
		tmp = prices[right] - prices[left]
		result = Max(result, tmp)
	}
	return result
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}