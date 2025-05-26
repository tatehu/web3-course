project-root/
│
├── contracts/                  # 所有智能合约源码
│   ├── base/                   # 基础拍卖合约
│   │   └── BaseAuction.sol
│   ├── extensions/             # 扩展功能合约（如价格预言机）
│   │   └── PriceFeedAuction.sol
│   ├── factory/                # 工厂合约
│   │   └── AuctionFactory.sol
│   ├── crosschain/             # 跨链拍卖合约
│   │   └── CrossChainAuction.sol
│   ├── upgradeable/            # 可升级拍卖合约
│   │   ├── UpgradeableAuction.sol
│   │   └── UpgradeableAuctionV2.sol
│   ├── mocks/                  # 测试用 Mock 合约
│   │   ├── MyNFT.sol
│   │   ├── MockERC20.sol
│   │   ├── MockV3Aggregator.sol
│   │   └── MockRouter.sol
│   └── interfaces/             # 合约接口定义
│       ├── IAuction.sol
│       ├── IPriceFeedAuction.sol
│       ├── IUpgradeableAuction.sol
│       ├── ICrossChainAuction.sol
│       ├── AggregatorV3Interface.sol
│       └── LinkTokenInterface.sol
│
├── scripts/
│   └── deploy/                 # 部署脚本
│       ├── base.js             # 部署基础及工厂合约
│       ├── price-feed.js       # 部署价格预言机相关合约
│       └── cross-chain.js      # 部署跨链相关合约
│
├── test/                       # 测试用例
│   ├── base/                   # 基础拍卖测试
│   │   └── baseAuction.test.js
│   ├── extensions/             # 扩展功能测试
│   │   └── priceFeedAuction.test.js
│   ├── upgradeable/            # 可升级拍卖测试
│   │   └── upgradeableAuction.test.js
│   ├── crosschain/             # 跨链拍卖测试
│   │   └── crossChainAuction.test.js
│   └── factory/                # 工厂合约测试
│       └── auctionFactory.test.js
│
├── hardhat.config.js           # Hardhat 配置文件
├── package.json                # 项目依赖与脚本
├── package-lock.json           # 依赖锁定
├── .gitignore                  # Git 忽略文件
└── README.md                   # 项目说明文档

主要合约说明
BaseAuction.sol：基础拍卖逻辑，支持 NFT 拍卖、出价、撤回、结束拍卖等基本功能。
PriceFeedAuction.sol：基于基础拍卖，增加价格预言机（如 Chainlink）支持，按 USD 价值判断出价门槛。
UpgradeableAuction.sol / UpgradeableAuctionV2.sol：可升级拍卖合约，支持合约升级，V2 增加了暂停/恢复功能。
CrossChainAuction.sol：支持跨链拍卖，集成 Chainlink CCIP，允许在多链间同步拍卖信息。
AuctionFactory.sol：工厂合约，统一部署和管理不同类型的拍卖合约实例。
Mocks 目录：包含 MyNFT（测试 NFT）、MockERC20（测试代币）、MockV3Aggregator（价格预言机 mock）、MockRouter（跨链路由 mock）等测试辅助合约。
interfaces 目录：定义各类合约的接口。

脚本说明
base.js：部署基础拍卖、价格预言机拍卖、可升级拍卖、跨链拍卖及工厂合约，并初始化测试 NFT 和代币。
price-feed.js：部署价格预言机相关合约及工厂，并创建价格预言机拍卖实例。
cross-chain.js：部署跨链相关合约及工厂，并创建跨链拍卖实例。

测试说明
本项目测试覆盖所有核心功能，测试文件分为以下几类：
基础拍卖测试（test/base/baseAuction.test.js）：覆盖拍卖创建、出价、撤回、结束拍卖、无人出价时 NFT 归还等场景。
扩展功能测试（test/extensions/priceFeedAuction.test.js）：覆盖价格预言机获取、USD 价值判断、最低出价限制、正常出价与结算等。
可升级拍卖测试（test/upgradeable/upgradeableAuction.test.js）：覆盖合约初始化、升级、升级后状态保持、暂停/恢复功能、升级后新功能可用性等。
跨链拍卖测试（test/crosschain/crossChainAuction.test.js）：覆盖跨链拍卖创建、事件验证、跨链消息、出价、结束拍卖、多拍卖实例等。
工厂合约测试（test/factory/auctionFactory.test.js）：覆盖工厂合约对不同类型拍卖的创建、参数校验、事件验证等。

一、安装依赖 
npm install
二、编译合约
npx hardhat compile
三、部署脚本
npx hardhat run scripts/deploy/base.js --network <network>
> <network> 可为 localhost、hardhat 或你配置的其他网络（如 goerli、sepolia 等）。
2. 部署价格预言机相关合约
npx hardhat run scripts/deploy/price-feed.js --network <network>
3. 部署跨链相关合约
npx hardhat run scripts/deploy/cross-chain.js --network <network>
四、测试
npx hardhat test test/...
