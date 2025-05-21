// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/*
 solidity：用 solidity 实现罗马数字转数整数
*/
contract RomanToInteger {
	function romanToInt(string memory s) public pure returns (uint256) {
		bytes memory str = bytes(s);
		uint256 len = str.length;
		uint256 total = 0;
		uint256 i = 0;

		while (i < len) {
			uint256 curr = _charValue(str[i]);
			uint256 next = 0;
			if (i + 1 < len) {
				next = _charValue(str[i + 1]);
			}

			if (next > curr) {
				total += (next - curr);
				i += 2;
			} else {
				total += curr;
				i += 1;
			}
		}
		return total;
	}

	// 将单个罗马字符转为整数
	function _charValue(bytes1 ch) internal pure returns (uint256) {
		if (ch == "I") return 1;
		if (ch == "V") return 5;
		if (ch == "X") return 10;
		if (ch == "L") return 50;
		if (ch == "C") return 100;
		if (ch == "D") return 500;
		if (ch == "M") return 1000;
		revert("undefined char");
	}
}
