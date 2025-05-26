const hre = require("hardhat");

async function main() {
  const [deployer] = await hre.ethers.getSigners();
  console.log("Deploying Cross Chain Auction with the account:", deployer.address);

  // 部署 Mock Router
  const MockRouter = await hre.ethers.getContractFactory("MockRouter");
  const router = await MockRouter.deploy();
  await router.waitForDeployment();
  console.log("Mock Router deployed to:", await router.getAddress());

  // 部署 Link Token
  const MockERC20 = await hre.ethers.getContractFactory("MockERC20");
  const linkToken = await MockERC20.deploy("Chainlink", "LINK");
  await linkToken.waitForDeployment();
  console.log("Link Token deployed to:", await linkToken.getAddress());

  // 部署 NFT 合约
  const MyNFT = await hre.ethers.getContractFactory("MyNFT");
  const myNFT = await MyNFT.deploy();
  await myNFT.waitForDeployment();
  console.log("MyNFT deployed to:", await myNFT.getAddress());

  // 部署支付代币
  const paymentToken = await MockERC20.deploy("Payment Token", "PAY");
  await paymentToken.waitForDeployment();
  console.log("Payment Token deployed to:", await paymentToken.getAddress());

  // 部署工厂合约
  const AuctionFactory = await hre.ethers.getContractFactory("AuctionFactory");
  const factory = await AuctionFactory.deploy();
  await factory.waitForDeployment();
  console.log("AuctionFactory deployed to:", await factory.getAddress());

  // 创建跨链拍卖
  const tx = await factory.createCrossChainAuction(
    await myNFT.getAddress(),
    0,
    3600,
    await paymentToken.getAddress(),
    await router.getAddress(),
    await linkToken.getAddress(),
    2, // destinationChainSelector
    await factory.getAddress() // destinationContract
  );
  const receipt = await tx.wait();
  const event = receipt.logs.find(
    log => log.fragment && log.fragment.name === "AuctionCreated"
  );
  console.log("Cross Chain Auction created at:", event.args.auction);

  // 铸造 NFT 和代币
  await myNFT.mint(deployer.address);
  await paymentToken.mint(deployer.address, hre.ethers.parseEther("1000"));
  await linkToken.mint(deployer.address, hre.ethers.parseEther("1000"));
  console.log("Minted NFT and tokens to deployer");

  // 验证合约
  if (hre.network.name !== "hardhat" && hre.network.name !== "localhost") {
    console.log("Waiting for block confirmations...");
    await router.deploymentTransaction().wait(6);
    await linkToken.deploymentTransaction().wait(6);
    await myNFT.deploymentTransaction().wait(6);
    await paymentToken.deploymentTransaction().wait(6);
    await factory.deploymentTransaction().wait(6);

    // await hre.run("verify:verify", {
    //   address: await router.getAddress(),
    //   constructorArguments: [],
    // });

    // await hre.run("verify:verify", {
    //   address: await linkToken.getAddress(),
    //   constructorArguments: ["Chainlink", "LINK"],
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