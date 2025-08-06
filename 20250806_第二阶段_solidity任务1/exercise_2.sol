// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

// 2 反转字符串 (Reverse String)
//   题目描述：反转一个字符串。输入 "abcde"，输出 "edcba"

contract ReverseString {
    //回文案例
    string[]  public vst = ["a", "b", "c", "b", "a"];

    // string[]  public vst = ["a", "b", "c", "b", "a", "a1"];

    bool public flag = true; // 定义一个布尔值，用于判断是否为回文

    function isPalindrome () public view returns (bool) {
        for (uint i = 0; i < vst.length/2; i++) {
            if ( keccak256(abi.encodePacked(vst[i])) != keccak256(abi.encodePacked(vst[vst.length - i - 1]))) {
                return false;
            }
        }
        return true;
    }
    
}