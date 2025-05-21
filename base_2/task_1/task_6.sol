// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/*
 solidity：二分查找 (Binary Search),
题目描述：在一个有序数组中查找目标值。
*/

contract BinarySearch {
	// 在有序数组arr中查找target，返回索引（找不到返回-1）
	function binarySearch(uint256[] memory arr, uint256 target) public pure returns (int256) {
		int256 left = 0;
		int256 right = int256(arr.length) - 1;

		while (left <= right) {
			int256 mid = left + (right - left) / 2;
			if (arr[uint256(mid)] == target) {
				return mid;
			} else if (arr[uint256(mid)] < target) {
				left = mid + 1;
			} else {
				right = mid - 1;
			}
		}
		return -1;
	}
}
