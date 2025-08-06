// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

//   二分查找 (Binary Search)
// 题目描述：在一个有序数组中查找目标值。

contract BinarySearch {


    function binarySearch(int256[] memory arr, int256 target) public pure returns (bool found, uint256 index) {
        uint256 left = 0;
        uint256 right = arr.length;

        arr = bubbleSort(arr) ;
        while (left < right) {
            // 防止整数溢出的中点计算
            uint256 mid = left + (right - left) / 2;

            if (arr[mid] == target) {
                return (true, mid);
            } else if (arr[mid] < target) {
                left = mid + 1;
            } else {
                right = mid;
            }
        }

        // 未找到目标值
        return (false, 0);
    }

      function bubbleSort(int256[] memory arr) public pure returns (int256[] memory) {
        uint256 n = arr.length;
        for (uint256 i = 0; i < n - 1; i++) {
            for (uint256 j = 0; j < n - i - 1; j++) {
                if (arr[j] > arr[j + 1]) {
                    // 交换元素
                    (arr[j], arr[j + 1]) = (arr[j + 1], arr[j]);
                }
            }
        }
        return arr;
    }

}