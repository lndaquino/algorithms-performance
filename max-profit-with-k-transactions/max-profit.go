package main

import "fmt"

type transaction struct {
	buyDay  int
	sellDay int
}

func main() {
	k := 2
	prices := []int{5, 11, 3, 50, 60, 90}
	// response := 93

	maxProfit := maxProfit(prices, k)
	fmt.Println(maxProfit)

}

// O(nk) time complexity | O(nk) space complexity
// we cut time complexity by storing the biggest tempProfit when iteracting over a line, without this optimization it would be O(n^2 * k)
// func maxProfitWithKTransactions(prices []int, k int) int {
// 	n := len(prices)
// 	if n == 0 {
// 		return 0
// 	}

// 	// var limit int
// 	// if n > k {
// 	// 	limit = n + 1
// 	// } else {
// 	// 	limit = k + 1
// 	// }

// 	// profit := make([][]int, limit)
// 	// for i := 0; i < n; i++ {
// 	// 	array := make([]int, limit)
// 	// 	for j := 0; j < n; j++ {
// 	// 		array[j] = 0
// 	// 	}
// 	// 	profit[i] = array
// 	// }
// 	profit := make([][]int, k+1)
// 	for i := range profit {
// 		profit[i] = make([]int, n)
// 	}

// 	for i := 1; i < n; i++ {
// 		for j := 1; j < k+1; j++ {
// 			maxProfit := 0 // optimization on time complexity

// 			for l := 0; l < i; l++ {
// 				tempProfit := prices[i] - prices[l] + profit[l][j-1]
// 				if tempProfit > maxProfit { // optimization on time complexity
// 					maxProfit = tempProfit
// 				}
// 			}
// 			if profit[i-1][j] > maxProfit {
// 				profit[i][j] = profit[i-1][j]
// 			} else {
// 				profit[i][j] = maxProfit
// 			}

// 		}
// 	}

// 	fmt.Println(profit)

// 	return profit[n-1][k]
// }

// O(nk) time complexity | O(n) space complexity
// we cut space complexity using only 2 rows that alternate values
func maxProfit(prices []int, k int) int {
	const MaxInt = int(^uint(0) >> 1)
	const MinInt = -MaxInt - 1
	n := len(prices)
	if n == 0 {
		return 0
	}

	evenProfits := make([]int, n)
	oddProfits := make([]int, n)
	for i := 0; i < n; i++ {
		evenProfits[i] = 0
		oddProfits[i] = 0
	}

	for t := 1; t <= k; t++ {
		maxProfit := MinInt
		currentProfits := make([]int, n)
		previousProfits := make([]int, n)
		if t%2 == 1 {
			currentProfits = oddProfits
			previousProfits = evenProfits
		} else {
			currentProfits = evenProfits
			previousProfits = oddProfits
		}
		for d := 1; d < n; d++ {
			maxProfit = max(maxProfit, previousProfits[d-1]-prices[d-1])
			currentProfits[d] = max(currentProfits[d-1], maxProfit+prices[d])
		}
	}

	if k%2 == 0 {
		return evenProfits[n-1]
	}

	return oddProfits[n-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
