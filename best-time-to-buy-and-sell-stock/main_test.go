package best_time_to_buy_and_sell_stock

import "testing"

func TestMaxProfit(t *testing.T) {
	fixtures := []struct {
		name     string
		prices   []int
		expected int
	}{
		{name: "test #1", prices: []int{7, 1, 5, 3, 6, 4}, expected: 5},
		{name: "test #2", prices: []int{7, 6, 5, 4, 3, 1}, expected: 0},
	}

	for _, tt := range fixtures {
		got := MaxProfit(tt.prices)
		assert(t, tt.name, tt.expected, got)
	}
}

func assert(t testing.TB, name string, expected, got int) {
	t.Helper()
	if expected != got {
		t.Errorf("%s: expected %d, got %d", name, expected, got)
	}
}
