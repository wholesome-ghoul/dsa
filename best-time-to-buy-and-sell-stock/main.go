package best_time_to_buy_and_sell_stock

import "math"

func MaxProfit(prices []int) int {
	return maxProfit(prices)
}

func maxProfit(prices []int) int {
	maxprofit := 0

	left := 0
	right := 1

	for right < len(prices) {
		maxprofit = int(math.Max(float64(prices[right]-prices[left]), float64(maxprofit)))

		if prices[left] > prices[right] {
			left = right
		}

		right++
	}

	return maxprofit
}
