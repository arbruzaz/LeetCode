package main

import (
	"fmt"
)

/*
121 Best Time to buy and sell stock
https://leetcode.com/problems/best-time-to-buy-and-sell-stock
*/
func maxProfit(prices []int) int {
	lens := len(prices)
	minPrice := prices[0]
	result := 0
	for i := 1; i < lens; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > result {
			result = prices[i] - minPrice
		}
	}

	return result
}

func main() {
	var prices = []int{7, 1, 5, 3, 6, 4}
	k := maxProfit(prices)
	fmt.Print(k)
}
