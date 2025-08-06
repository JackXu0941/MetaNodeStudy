
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract IntegerToRoman {
    /**
     * @dev Converts an integer to a Roman numeral.
     * This function supports numbers typically representable by standard Roman numerals (up to ~3999).
     */
    function intToRoman(uint256 num) public pure returns (string memory) {
        // Arrays holding the integer values and their corresponding Roman numeral symbols.
        // Ordered from largest to smallest, including the special subtractive cases (IV, IX, XL, XC, CD, CM).
        // uint256[13] memory values = uint256[]([1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];
        string[13] memory numerals = ["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"];

        require(num > 0 && num < 4000, "Input number must be between 1 and 3999");

        string memory result = "";

        for (uint256 i = 0; i < values.length; i++) {
            while (num >= values[i]) {
                result = string(abi.encodePacked(result, numerals[i]));
                num -= values[i];
            }
        }

        return result;
    }
}