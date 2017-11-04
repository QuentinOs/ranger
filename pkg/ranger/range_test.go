package ranger

import (
	"testing"
)

func TestNewRangeRepo(t *testing.T) {
	bounds := []Bound{
		{
			Upper: false,
			Value: 0.0,
		},
		{
			Upper: true,
			Value: 2.0,
		},
		{
			Upper: false,
			Value: 1.0,
		},
		{
			Upper: true,
			Value: 3.0,
		},
		{
			Upper: false,
			Value: 4.0,
		},
		{
			Upper: true,
			Value: 5.0,
		},
	}

	repo := New(bounds)
	// Case : in 1 range
	testCountRangeOnValue(repo, t, 0.5, 1)

	// Case : in 2 ranges
	testCountRangeOnValue(repo, t, 1.5, 2)

	// Case : in between two ranges
	testCountRangeOnValue(repo, t, 3.5, 0)

	// Case : smaller that lower bound
	testCountRangeOnValue(repo, t, -0.5, 0)

	// Case : bigger that higher bound
	testCountRangeOnValue(repo, t, 6.0, 0)

	// Case : On range limit openning
	testCountRangeOnValue(repo, t, 1.0, 2)

	// Case : On range limit closing
	testCountRangeOnValue(repo, t, 2.0, 1)
}

func testCountRangeOnValue(repo *RangeRepo, t *testing.T, input float64, expected int) {
	if repo.CountRangesForValue(input) != expected {
		t.Errorf("expecting %d, got %d", expected, repo.CountRangesForValue(input))
	}
}
