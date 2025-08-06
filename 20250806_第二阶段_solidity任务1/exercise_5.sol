// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// 5 .合并两个有序数组 (Merge Sorted Array)
// 题目描述：将两个有序数组合并为一个有序数组。

contract MergeSortedArrays {
    function mergeSortedArrays(int[] memory arr1, int[] memory arr2) public pure returns (int[] memory) {
        uint len1 = arr1.length;
        uint len2 = arr2.length;
        int[] memory merged = new int[](len1 + len2);

        uint i = 0;
        uint j = 0;
        uint k = 0;

        while (i < len1 && j < len2) {
            if (arr1[i] <= arr2[j]) {
                merged[k] = arr1[i];
                i++;
            } else {
                merged[k] = arr2[j];
                j++;
            }
            k++;
        }

        while (i < len1) {
            merged[k] = arr1[i];
            i++;
            k++;
        }

        while (j < len2) {
            merged[k] = arr2[j];
            j++;
            k++;
        }

        return merged;
    }
}