// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/*
 solidity： 合并两个有序数组 (Merge Sorted Array),
题目描述：将两个有序数组合并为一个有序数组。
*/

contract MergeSortedArray {
	// 合并两个有序的uint数组
	function merge(uint256[] memory arr1, uint256[] memory arr2) public pure returns (uint256[] memory) {
		uint256 m = arr1.length;
		uint256 n = arr2.length;
		uint256[] memory result = new uint256[](m + n);

		uint256 i = 0; // index for arr1
		uint256 j = 0; // index for arr2
		uint256 k = 0; // index for result

		while (i < m && j < n) {
			if (arr1[i] <= arr2[j]) {
				result[k] = arr1[i];
				i++;
			} else {
				result[k] = arr2[j];
				j++;
			}
			k++;
		}

		// 处理剩余元素
		while (i < m) {
			result[k++] = arr1[i++];
		}

		while (j < n) {
			result[k++] = arr2[j++];
		}

		return result;
	}
}

