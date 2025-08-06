// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

//  3.用 solidity 实现整数转罗马数字

contract RomanToInteger {


    function romanToInt(string memory roman) public pure returns (uint256) {
        bytes memory romanBytes = bytes(roman);
        uint256 total = 0;
        uint256 prevValue = 0;

        // Process the string from right to left
        for (int256 i = int256(romanBytes.length) - 1; i >= 0; i--) {

            uint256 currentValue = _charToInt(romanBytes[uint256(i)]);

            // If the current value is less than the previous value, subtract it; otherwise, add it.
            if (currentValue >= prevValue) {
                total += currentValue;
            } else {
                total -= currentValue;
            }

            prevValue = currentValue;
        }

        return total;
    }

    //
    function _charToInt(bytes1 char) internal pure returns (uint256) {
        if (char == 'I') {
            return 1;
        } else if (char == 'V') {
            return 5;
        } else if (char == 'X') {
            return 10;
        } else if (char == 'L') {
            return 50;
        } else if (char == 'C') {
            return 100;
        } else if (char == 'D') {
            return 500;
        } else if (char == 'M') {
            return 1000;
        } else {
            // Invalid character, return 0 or handle error as needed.
            return 0;
        }
    }
}