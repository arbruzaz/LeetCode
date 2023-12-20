package main

import (
	"fmt"
)

/*
122 Best Time to buy and sell stock II
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii
*/
func maxProfit(prices []int) int {
	profit := 0

	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			profit += prices[i+1] - prices[i]
		}
	}

	return profit
}

func main() {
	var prices = []int{7, 1, 5, 6, 1, 5}
	k := maxProfit(prices)
	fmt.Print(k)
}
