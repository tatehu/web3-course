// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/*
solidity：ERC20 代币,
任务：参考 openzeppelin-contracts/contracts/token/ERC20/IERC20.sol实现一个简单的 ERC20 代币合约。要求：
1，合约包含以下标准 ERC20 功能：,
2，balanceOf：查询账户余额。,
3，transfer：转账。,
4，approve 和 transferFrom：授权和代扣转账。,
5，使用 event 记录转账和授权操作。,
6，提供 mint 函数，允许合约所有者增发代币。,
提示：
	1，使用 mapping 存储账户余额和授权信息。,
	2，使用 event 定义 Transfer 和 Approval 事件。,
	3，部署到sepolia 测试网，导入到自己的钱包
*/

contract SimpleERC20 {
	string public name = "MyToken";
	string public symbol = "MTK";
	uint8 public decimals = 18;
	uint256 public totalSupply;

	address public owner;

	mapping(address => uint256) private balances;
	mapping(address => mapping(address => uint256)) private allowances;

	event Transfer(address indexed from, address indexed to, uint256 value);
	event Approval(address indexed owner, address indexed spender, uint256 value);

	modifier onlyOwner() {
		require(msg.sender == owner, "Only owner can mint");
		_;
	}

	constructor() {
		owner = msg.sender;
	}

	// 查询余额
	function balanceOf(address account) external view returns (uint256) {
		return balances[account];
	}

	// 转账
	function transfer(address to, uint256 amount) external returns (bool) {
		require(balances[msg.sender] >= amount, "insufficient balance");
		require(to != address(0), "cannot transfer address zero");

		balances[msg.sender] -= amount;
		balances[to] += amount;
		emit Transfer(msg.sender, to, amount);

		return true;
	}

	// 授权
	function approve(address spender, uint256 amount) external returns (bool) {
		allowances[msg.sender][spender] = amount;
		emit Approval(msg.sender, spender, amount);

		return true;
	}

	// 允许被授权人转账
	function transferFrom(address from, address to, uint256 amount) external returns (bool) {
		require(balances[from] >= amount, "insufficient balance");
		require(allowances[from][msg.sender] >= amount, "The authorized amount is insufficient");
		require(to != address(0), "cannot transfer address zero");

		balances[from] -= amount;
		balances[to] += amount;
		allowances[from][msg.sender] -= amount;
		emit Transfer(from, to, amount);

		return true;
	}

	// 授权额度查询
	function allowance(address _owner, address spender) external view returns (uint256) {
		return allowances[_owner][spender];
	}

	// 铸币（仅限合约所有者）
	function mint(address to, uint256 amount) external onlyOwner {
		require(to != address(0), "cannot mint address zero");
		totalSupply += amount;
		balances[to] += amount;
		emit Transfer(address(0), to, amount);
	}
}

