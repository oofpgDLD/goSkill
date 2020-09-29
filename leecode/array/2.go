package array

func maxProfit(prices []int) int {
	maxprofit := 0
	valley := prices[0]
	peak := prices[0]
	i := 0
	for i < len(prices)-1{
		for i < len(prices)-1 && prices[i] >= prices[i + 1] {
			i++
		}
		valley = prices[i]
		for i < len(prices)-1 && prices[i] <= prices[i+1] {
			i++
		}
		peak = prices[i]
		maxprofit += peak - valley
	}
	return maxprofit
}