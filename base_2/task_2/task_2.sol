// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
/*
作业2：在测试网上发行一个图文并茂的 NFT,
任务目标
1，使用 Solidity 编写一个符合 ERC721 标准的 NFT 合约。,
2，将图文数据上传到 IPFS，生成元数据链接。,
3，将合约部署到以太坊测试网（如 Goerli 或 Sepolia）。,
4，铸造 NFT 并在测试网环境中查看。,
任务步骤
1，编写 NFT 合约
	1，使用 OpenZeppelin 的 ERC721 库编写一个 NFT 合约。,
	2，合约应包含以下功能：
		1，构造函数：设置 NFT 的名称和符号。,
		2，mintNFT 函数：允许用户铸造 NFT，并关联元数据链接（tokenURI）。,
		3，在 Remix IDE 中编译合约。
2，准备图文数据
	1，准备一张图片，并将其上传到 IPFS（可以使用 Pinata 或其他工具）。,
	2，创建一个 JSON 文件，描述 NFT 的属性（如名称、描述、图片链接等）。,
	3，将 JSON 文件上传到 IPFS，获取元数据链接。,
	4，JSON文件参考 https://docs.opensea.io/docs/metadata-standards,
,
3，部署合约到测试网
	1，在 Remix IDE 中连接 MetaMask，并确保 MetaMask 连接到 Goerli 或 Sepolia 测试网。,
	2，部署 NFT 合约到测试网，并记录合约地址。,
4，铸造 NFT
	1，使用 mintNFT 函数铸造 NFT：,
	2，在 recipient 字段中输入你的钱包地址。,
	3，在 tokenURI 字段中输入元数据的 IPFS 链接。,
	4，在 MetaMask 中确认交易。,
5，查看 NFT
	1，打开 OpenSea 测试网 或 Etherscan 测试网。,
	2，连接你的钱包，查看你铸造的 NFT。
*/

contract MyNFT is ERC721URIStorage, Ownable {
	uint256 public tokenCounter;

	constructor(string memory name, string memory symbol)
	ERC721(name, symbol)
	Ownable(msg.sender)   // 传递msg.sender给Ownable
	{
		tokenCounter = 0;
	}

	function mintNFT(address recipient, string memory tokenURI) public returns (uint256) {
		uint256 newTokenId = tokenCounter;
		_safeMint(recipient, newTokenId);
		_setTokenURI(newTokenId, tokenURI);
		tokenCounter += 1;
		return newTokenId;
	}
}
