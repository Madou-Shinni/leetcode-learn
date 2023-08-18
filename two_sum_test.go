package leetcode_learn

import (
	"fmt"
	"testing"
)

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// 在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。
// Given nums = [2, 7, 11, 15], target = 9,
//
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1]
func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	anotherNumMap := make(map[int]int)
	// 1.顺序扫描数组
	for i, item := range nums {
		// 2.在map中找到能组合成给定值的另一半数字
		anotherNum := target - item
		if _, ok := anotherNumMap[anotherNum]; ok {
			// 3.找到 返回两数下标
			return []int{anotherNumMap[anotherNum], i}
		}

		// 3.没找到，当前数字和下标存入map
		anotherNumMap[item] = i
	}

	return nil
}
