package main

import "fmt"

// 第六题
// 26. 删除有序数组中的重复项：给你一个有序数组(非严格递增排列) nums ，请你原地删除重复出现的元素，
// 使每个元素只出现一次，返回删除后数组的新长度。不要使用额外的数组空间，
// 你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。可以使用双指针法，
// 一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，
// 当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。

// leetCode 原题:
// 给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
// 元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
// 考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
// 更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
// 返回 k 。

func Question_6() int {

	// digits := []int{1, 2, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 5, 8, 9}  // 错误 , 不是非严格递排列
	//实例1 返回[1,2,3,4,5,8,9]
	digits := []int{1, 1, 1, 2, 2, 2, 3, 3, 4, 4, 5, 5, 5, 8, 9} // 正确 , 是非严格递排列

	count := len(digits)
	if count == 0 {
		fmt.Println("数组为空")
		return 0
	}
	i := 1

	for j := 1; j < count; j++ {
		if digits[j] != digits[i-1] {
			digits[i] = digits[j]
			i++
		}
	}
	digits = digits[:i]
	fmt.Println(digits)
	return i

}
