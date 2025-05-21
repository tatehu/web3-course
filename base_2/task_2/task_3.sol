// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";


/*
作业3：编写一个讨饭合约,
任务目标
1，使用 Solidity 编写一个合约，允许用户向合约地址发送以太币。,
2，记录每个捐赠者的地址和捐赠金额。,
3，允许合约所有者提取所有捐赠的资金。,

任务步骤
1，编写合约
创建一个名为 BeggingContract 的合约。,
合约应包含以下功能：,
一个 mapping 来记录每个捐赠者的捐赠金额。,
一个 donate 函数，允许用户向合约发送以太币，并记录捐赠信息。,
一个 withdraw 函数，允许合约所有者提取所有资金。,
一个 getDonation 函数，允许查询某个地址的捐赠金额。,
使用 payable 修饰符和 address.transfer 实现支付和提款。,
,
2，部署合约
在 Remix IDE 中编译合约。,
部署合约到 Goerli 或 Sepolia 测试网。,
,
3，测试合约
使用 MetaMask 向合约发送以太币，测试 donate 功能。,
调用 withdraw 函数，测试合约所有者是否可以提取资金。,
调用 getDonation 函数，查询某个地址的捐赠金额。,
,

任务要求
1，合约代码：
使用 mapping 记录捐赠者的地址和金额。,
使用 payable 修饰符实现 donate 和 withdraw 函数。,
使用 onlyOwner 修饰符限制 withdraw 函数只能由合约所有者调用。,
,
2，测试网部署：
合约必须部署到 Goerli 或 Sepolia 测试网。,
,
3，功能测试：
确保 donate、withdraw 和 getDonation 函数正常工作。,
,

提交内容
1，合约代码：提交 Solidity 合约文件（如 BeggingContract.sol）。,
2，合约地址：提交部署到测试网的合约地址。,
3，测试截图：提交在 Remix 或 Etherscan 上测试合约的截图。,

额外挑战（可选）
1，捐赠事件：添加 Donation 事件，记录每次捐赠的地址和金额。,
2，捐赠排行榜：实现一个功能，显示捐赠金额最多的前 3 个地址。,
3，时间限制：添加一个时间限制，只有在特定时间段内才能捐赠。
*/

contract BeggingContract is Ownable {
	// 记录每位捐赠者累计捐赠金额
	mapping(address => uint256) public donations;

	// “挑战1”：捐赠事件
	event Donation(address indexed donor, uint256 amount);

	// “挑战2”：捐赠者地址数组（用于排行）
	address[] private donorList;

	// “挑战3”：时间段限制
	uint256 public donateStartTime;
	uint256 public donateEndTime;

	// 构造函数可设置捐赠起止时间（单位：秒，0 为无限制）
	constructor(uint256 _start, uint256 _end) Ownable(msg.sender) {
		donateStartTime = _start;
		donateEndTime = _end;
	}

	// 捐赠函数
	function donate() external payable {
		require(msg.value > 0, "donate must > 0");

		// 挑战3：如启用时间限制，检查当前时间
		if (donateStartTime != 0 && block.timestamp < donateStartTime) {
			revert("donate time is not start");
		}
		if (donateEndTime != 0 && block.timestamp > donateEndTime) {
			revert("donate time is over");
		}

		// 首次捐赠的加入捐赠者列表（便于做排行榜）
		if (donations[msg.sender] == 0) {
			donorList.push(msg.sender);
		}
		donations[msg.sender] += msg.value;

		emit Donation(msg.sender, msg.value);
	}

	// 合约拥有者提取所有捐赠
	function withdraw() external onlyOwner {
		require(address(this).balance > 0, "balance is zero");
		payable(owner()).transfer(address(this).balance);
	}

	// 查询某个地址已捐金额
	function getDonation(address donor) external view returns (uint256) {
		return donations[donor];
	}

	// 查看合约当前余额
	function getBalance() external view returns (uint256) {
		return address(this).balance;
	}

	// “挑战2”：获取捐赠排行榜（前3名）
	function getTopDonors() external view returns (address[3] memory, uint256[3] memory) {
		address[3] memory topAddrs;
		uint256[3] memory topAmounts;

		for (uint256 i = 0; i < donorList.length; i++) {
			uint256 amount = donations[donorList[i]];
			// 简单三名插入排序
			for (uint256 j = 0; j < 3; j++) {
				if (amount > topAmounts[j]) {
					// 右移
					for (uint256 k = 2; k > j; k--) {
						topAmounts[k] = topAmounts[k - 1];
						topAddrs[k] = topAddrs[k - 1];
					}
					topAmounts[j] = amount;
					topAddrs[j] = donorList[i];
					break;
				}
			}
		}
		return (topAddrs, topAmounts);
	}
}
