package main

import (
	"fmt"
	"sort"
)

// 第七题
// 56. 合并区间：以数组 intervals 表示若干个区间的集合，
// 其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，
// 并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
// 可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，
// 遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；
// 如果没有重叠，则将当前区间添加到切片中。

func Question_7() {

	//实例1 返回
	// digits := [][]int{{0, 1}, {0, 3}, {0, 2}, {1, 2}, {1, 4}, {2, 5}, {11, 12}, {8, 9}, {22, 26}}
	digits := [][]int{{1, 4}, {22, 26}, {0, 3}, {8, 9}, {0, 2}, {1, 2}, {0, 1}, {11, 12}, {2, 5}}

	//先根据 每个数组的第一个元素进行排序
	sort.Slice(digits, func(i, j int) bool {
		return digits[i][0] < digits[j][0]
	})
	fmt.Println(digits)

	start, end := digits[0][0], digits[0][1]
	reuslt := [][]int{}

	reuslt = append(reuslt, []int{start, end})
	fmt.Println(reuslt)

	for i := 0; i < len(digits)-1; i++ {

		//因为是根据每一个子数组的第一个元素升序的，所以如果第一个元素相等情况下，则将end进行更新
		if digits[i][0] == digits[i+1][0] {
			start = digits[i][0]
			end = max(end, digits[i+1][1])
			reuslt[len(reuslt)-1][1] = end

		} else {

			//digits[i][0] < digits[i+1][0] ,且 digits[i+1][1] > end ,这说明不是一个区间，需要添加新的区间
			if digits[i+1][1] > end {
				start = digits[i+1][0]
				end = digits[i+1][1]
				reuslt = append(reuslt, []int{start, end})
			}

		}
	}

	fmt.Println(reuslt)

}
