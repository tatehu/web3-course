const hre = require("hardhat");

async function main() {
  const [deployer] = await hre.ethers.getSigners();
  console.log("Deploying contracts with the account:", deployer.address);

  // 部署四个实现合约
  const BaseAuction = await hre.ethers.getContractFactory("BaseAuction");
  const baseAuctionImpl = await BaseAuction.deploy();
  await baseAuctionImpl.waitForDeployment();
  console.log("BaseAuction implementation deployed to:", await baseAuctionImpl.getAddress());

  const PriceFeedAuction = await hre.ethers.getContractFactory("PriceFeedAuction");
  const priceFeedAuctionImpl = await PriceFeedAuction.deploy();
  await priceFeedAuctionImpl.waitForDeployment();
  console.log("PriceFeedAuction implementation deployed to:", await priceFeedAuctionImpl.getAddress());

  const UpgradeableAuction = await hre.ethers.getContractFactory("UpgradeableAuction");
  const upgradeableAuctionImpl = await UpgradeableAuction.deploy();
  await upgradeableAuctionImpl.waitForDeployment();
  console.log("UpgradeableAuction implementation deployed to:", await upgradeableAuctionImpl.getAddress());

  const CrossChainAuction = await hre.ethers.getContractFactory("CrossChainAuction");
  const crossChainAuctionImpl = await CrossChainAuction.deploy();
  await crossChainAuctionImpl.waitForDeployment();
  console.log("CrossChainAuction implementation deployed to:", await crossChainAuctionImpl.getAddress());

  // 部署工厂合约，传入四个实现合约地址
  const AuctionFactory = await hre.ethers.getContractFactory("AuctionFactory");
  const factory = await AuctionFactory.deploy(
    await baseAuctionImpl.getAddress(),
    await priceFeedAuctionImpl.getAddress(),
    await upgradeableAuctionImpl.getAddress(),
    await crossChainAuctionImpl.getAddress()
  );
  await factory.waitForDeployment();
  console.log("AuctionFactory deployed to:", await factory.getAddress());

  // 铸造一些 NFT 和代币用于测试
  const MyNFT = await hre.ethers.getContractFactory("MyNFT");
  const myNFT = await MyNFT.deploy();
  await myNFT.waitForDeployment();
  console.log("MyNFT deployed to:", await myNFT.getAddress());

  const MockERC20 = await hre.ethers.getContractFactory("MockERC20");
  const paymentToken = await MockERC20.deploy("Payment Token", "PAY");
  await paymentToken.waitForDeployment();
  console.log("Payment Token deployed to:", await paymentToken.getAddress());

  await myNFT.mint(deployer.address);
  await paymentToken.mint(deployer.address, hre.ethers.parseEther("1000"));
  console.log("Minted NFT and tokens to deployer");

  // 验证合约
  if (hre.network.name !== "hardhat" && hre.network.name !== "localhost") {
    console.log("Waiting for block confirmations...");
    await myNFT.deploymentTransaction().wait(6);
    await paymentToken.deploymentTransaction().wait(6);
    await factory.deploymentTransaction().wait(6);

    // await hre.run("verify:verify", {
    //   address: await myNFT.getAddress(),
    //   constructorArguments: [],
    // });

    // await hre.run("verify:verify", {
    //   address: await paymentToken.getAddress(),
    //   constructorArguments: ["Payment Token", "PAY"],
    // });

    // await hre.run("verify:verify", {
    //   address: await factory.getAddress(),
    //   constructorArguments: [
    //     await baseAuctionImpl.getAddress(),
    //     await priceFeedAuctionImpl.getAddress(),
    //     await upgradeableAuctionImpl.getAddress(),
    //     await crossChainAuctionImpl.getAddress()
    //   ],
    // });
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  }); 