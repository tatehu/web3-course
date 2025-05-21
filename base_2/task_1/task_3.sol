// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/*
 solidity：用 solidity 实现整数转罗马数字
*/

contract IntegerToRoman {
	// 支持1-3999之间的整数
	function toRoman(uint256 num) public pure returns (string memory) {
		require(num >= 1 && num <= 3999, "only transform 1~3999");

		string[13] memory romans = [
					"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"
			];
		uint16[13] memory values = [
					1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1
			];

		bytes memory strBytes;

		for (uint8 i = 0; i < 13; i++) {
			while (num >= values[i]) {
				strBytes = abi.encodePacked(strBytes, romans[i]);
				num -= values[i];
			}
		}
		return string(strBytes);
	}
}