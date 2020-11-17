package main

// run with : go test -v
import "testing"

func TestMaxProfitWithKTransactions(t *testing.T) {
	t.Run("Teste 1", func(t *testing.T) {
		k := 2
		prices := []int{5, 11, 3, 50, 60, 90}
		response := 93

		result := maxProfit(prices, k)

		if result != response {
			t.Errorf("Expected response: %d - your response: %d", response, result)
		}
	})

	t.Run("Teste 2", func(t *testing.T) {
		k := 1
		var prices []int
		response := 0

		result := maxProfit(prices, k)

		if result != response {
			t.Errorf("Expected response: %d - your response: %d", response, result)
		}
	})

	t.Run("Teste 3", func(t *testing.T) {
		k := 1
		prices := []int{1}
		response := 0

		result := maxProfit(prices, k)

		if result != response {
			t.Errorf("Expected response: %d - your response: %d", response, result)
		}
	})

	t.Run("Teste 4", func(t *testing.T) {
		k := 1
		prices := []int{1, 10}
		response := 9

		result := maxProfit(prices, k)

		if result != response {
			t.Errorf("Expected response: %d - your response: %d", response, result)
		}
	})

	t.Run("Teste 5", func(t *testing.T) {
		k := 3
		prices := []int{1, 10}
		response := 9

		result := maxProfit(prices, k)

		if result != response {
			t.Errorf("Expected response: %d - your response: %d", response, result)
		}
	})

	t.Run("Teste 6", func(t *testing.T) {
		k := 1
		prices := []int{3, 2, 5, 7, 1, 3, 7}
		response := 6

		result := maxProfit(prices, k)

		if result != response {
			t.Errorf("Expected response: %d - your response: %d", response, result)
		}
	})

	t.Run("Teste 7", func(t *testing.T) {
		k := 3
		prices := []int{5, 11, 3, 50, 60, 90}
		response := 93

		result := maxProfit(prices, k)

		if result != response {
			t.Errorf("Expected response: %d - your response: %d", response, result)
		}
	})

	t.Run("Teste 8", func(t *testing.T) {
		k := 2
		prices := []int{5, 11, 3, 50, 40, 90}
		response := 97

		result := maxProfit(prices, k)

		if result != response {
			t.Errorf("Expected response: %d - your response: %d", response, result)
		}
	})

	t.Run("Teste 9", func(t *testing.T) {
		k := 3
		prices := []int{5, 11, 3, 50, 40, 90}
		response := 103

		result := maxProfit(prices, k)

		if result != response {
			t.Errorf("Expected response: %d - your response: %d", response, result)
		}
	})

	t.Run("Teste 10", func(t *testing.T) {
		k := 2
		prices := []int{50, 25, 12, 4, 3, 10, 1, 100}
		response := 106

		result := maxProfit(prices, k)

		if result != response {
			t.Errorf("Expected response: %d - your response: %d", response, result)
		}
	})

	t.Run("Teste 11", func(t *testing.T) {
		k := 5
		prices := []int{100, 99, 98, 97, 1}
		response := 0

		result := maxProfit(prices, k)

		if result != response {
			t.Errorf("Expected response: %d - your response: %d", response, result)
		}
	})

	t.Run("Teste 12", func(t *testing.T) {
		k := 5
		prices := []int{1, 100, 2, 200, 3, 300, 4, 400, 5, 500}
		response := 1485

		result := maxProfit(prices, k)

		if result != response {
			t.Errorf("Expected response: %d - your response: %d", response, result)
		}
	})

	t.Run("Teste 13", func(t *testing.T) {
		k := 5
		prices := []int{1, 100, 101, 200, 201, 300, 301, 400, 401, 500}
		response := 499

		result := maxProfit(prices, k)

		if result != response {
			t.Errorf("Expected response: %d - your response: %d", response, result)
		}
	})

	t.Run("Teste 14", func(t *testing.T) {
		k := 4
		prices := []int{1, 25, 24, 23, 12, 36, 14, 40, 31, 41, 5}
		response := 84

		result := maxProfit(prices, k)

		if result != response {
			t.Errorf("Expected response: %d - your response: %d", response, result)
		}
	})

	t.Run("Teste 15", func(t *testing.T) {
		k := 2
		prices := []int{1, 25, 24, 23, 12, 36, 14, 40, 31, 41, 5}
		response := 62

		result := maxProfit(prices, k)

		if result != response {
			t.Errorf("Expected response: %d - your response: %d", response, result)
		}
	})
}
