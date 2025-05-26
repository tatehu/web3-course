const hre = require("hardhat");

async function main() {
  const [deployer] = await hre.ethers.getSigners();
  console.log("Deploying Price Feed Auction with the account:", deployer.address);

  // 部署价格预言机
  const MockV3Aggregator = await hre.ethers.getContractFactory("MockV3Aggregator");
  const priceFeed = await MockV3Aggregator.deploy(8, 200000000); // 2000 USD with 8 decimals
  await priceFeed.waitForDeployment();
  console.log("Price Feed deployed to:", await priceFeed.getAddress());

  // 部署 NFT 合约
  const MyNFT = await hre.ethers.getContractFactory("MyNFT");
  const myNFT = await MyNFT.deploy();
  await myNFT.waitForDeployment();
  console.log("MyNFT deployed to:", await myNFT.getAddress());

  // 部署支付代币
  const MockERC20 = await hre.ethers.getContractFactory("MockERC20");
  const paymentToken = await MockERC20.deploy("Payment Token", "PAY");
  await paymentToken.waitForDeployment();
  console.log("Payment Token deployed to:", await paymentToken.getAddress());

  // 部署工厂合约
  const AuctionFactory = await hre.ethers.getContractFactory("AuctionFactory");
  const factory = await AuctionFactory.deploy();
  await factory.waitForDeployment();
  console.log("AuctionFactory deployed to:", await factory.getAddress());

  // 创建价格预言机拍卖
  const tx = await factory.createPriceFeedAuction(
    await myNFT.getAddress(),
    0,
    3600,
    await paymentToken.getAddress(),
    await priceFeed.getAddress()
  );
  const receipt = await tx.wait();
  const event = receipt.logs.find(
    log => log.fragment && log.fragment.name === "AuctionCreated"
  );
  console.log("Price Feed Auction created at:", event.args.auction);

  // 铸造 NFT 和代币
  await myNFT.mint(deployer.address);
  await paymentToken.mint(deployer.address, hre.ethers.parseEther("1000"));
  console.log("Minted NFT and tokens to deployer");

  // 验证合约
  if (hre.network.name !== "hardhat" && hre.network.name !== "localhost") {
    console.log("Waiting for block confirmations...");
    await priceFeed.deploymentTransaction().wait(6);
    await myNFT.deploymentTransaction().wait(6);
    await paymentToken.deploymentTransaction().wait(6);
    await factory.deploymentTransaction().wait(6);

    // await hre.run("verify:verify", {
    //   address: await priceFeed.getAddress(),
    //   constructorArguments: [8, 200000000],
    // });

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
    //   constructorArguments: [],
    // });
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  }); 