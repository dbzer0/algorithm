/*
	Дан массив целых чисел, нужно вернуть индексы двух чисел сумма которых даст
    запрашиваемый (target) элемент.
    Нельзя использовать один элемент в сложении дважды. Нужно выдать только одно решение.

 	Например:

	Дан:	    nums = [2, 7, 11, 15], target = 9,

	Потому что: nums[0] + nums[1] = 2 + 7 = 9,
	Возврат:    [0, 1].
*/

package main

import (
	"fmt"
)

// Тупой алгоритм сложности ~ O(n^2).
func twoSum(nums []int, target int) []int {
	for i, n := range nums {
		for j := i + 1; j < len(nums); j++ {
			if target-n == nums[j] {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{1, 2, 4, 5}, 6))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
