package range_sum_1d

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixSum(t *testing.T) {
	a := Constructor([]int{-2, 0, 3, -5, 2, -1})
	assert.Equal(t, a.SumRange(0, 2), 1)
	assert.Equal(t, a.SumRange(2, 5), -1)
}
