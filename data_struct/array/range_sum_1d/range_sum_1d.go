// 前缀和
package range_sum_1d

import "fmt"

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)

	for k := 1; k < len(sums); k++ {
		sums[k] = nums[k-1] + sums[k-1]
	}
	fmt.Println(sums)
	return NumArray{sums: sums}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sums[right+1] - this.sums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
