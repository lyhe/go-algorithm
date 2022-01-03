package range_sum_2d

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeSum2D(t *testing.T) {
	a := Constructor([][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}})
	assert.Equal(t, a.SumRegion(2, 1, 4, 3), 8)
}
