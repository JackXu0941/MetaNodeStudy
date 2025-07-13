package main

import "fmt"

// 第四题
// 最长公共前缀
// 考察：字符串处理、循环嵌套
// 题目：查找字符串数组中的最长公共前缀

func Question_4() string {
	// strs := []string{"abcower", "abcdow", "abight", "abcdefg", "abcdefghijklmnopqrstuvwxyz"}
	strs := []string{"abcower", "abcdow", "abight", "abcdefg", "wqcdefghijklmnopqrstuvwxyz"}

	result := LongestCommonPrefix(strs)
	fmt.Println("最长前缀是:" + result)
	return result
}

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	count := len(strs)
	for i := 1; i < count; i++ {
		prefix = findCommonPrefix(prefix, strs[i])
		if len(prefix) == 0 {
			return ""
		}
	}
	return prefix

}

func findCommonPrefix(str1, str2 string) string {
	lenth := min(len(str1), len(str2))
	index := 0
	for index < lenth && str1[index] == str2[index] {
		index++
	}
	return str1[:index]

}
