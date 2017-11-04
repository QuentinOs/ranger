package ranger

import (
	"sort"
)

// Bounds represent a limit of a range
type Bound struct {
	Value float64
	Upper bool
}

type boundRepo []Bound

func (b boundRepo) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b boundRepo) Less(i, j int) bool {
	return b[i].Value < b[j].Value
}

func (b boundRepo) Len() int {
	return len(b)
}

// Range defines a range between two float64
type Range struct {
	From    float64
	To      float64
	Counter int
}

// RangeRepo is a repo of ranges
type RangeRepo struct {
	ranges []Range
}

// New creates a new RangeRepo from a list of
// lower & higher bounds
func New(bounds []Bound) *RangeRepo {
	bRepo := boundRepo(bounds)
	sort.Sort(bRepo)

	ranges := make([]Range, 0)
	acc := 0
	var from float64

	for _, b := range bRepo {
		ranges = append(ranges, Range{
			From:    from,
			To:      b.Value,
			Counter: acc,
		})

		if !b.Upper {
			acc++
		} else {
			acc--
		}

		from = b.Value
	}

	return &RangeRepo{
		ranges: ranges,
	}
}

// CountRangesForValue returns the number of range the input
// belongs to.
func (repo *RangeRepo) CountRangesForValue(v float64) int {
	if len(repo.ranges) == 0 {
		return 0
	}

	lo, hi := 0, len(repo.ranges)-1
	if v < repo.ranges[lo].From || v >= repo.ranges[hi].To {
		return 0
	}

	for lo <= hi {
		mid := (hi + lo) / 2
		midRange := repo.ranges[mid]

		if v < midRange.From {
			hi = mid - 1
			continue
		}
		if v >= midRange.To {
			lo = mid + 1
			continue
		}

		return midRange.Counter
	}
	return 0
}
