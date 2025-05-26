const { expect } = require("chai");
const { ethers, upgrades } = require("hardhat");

describe("Price Feed Auction", function () {
  let myNFT, auction, paymentToken, priceFeed, owner, bidder1, bidder2;

  beforeEach(async function () {
    [owner, bidder1, bidder2] = await ethers.getSigners();

    // 部署 Mock Price Feed
    const MockPriceFeed = await ethers.getContractFactory("MockV3Aggregator");
    priceFeed = await MockPriceFeed.deploy(8, 200000000); // 2000 USD with 8 decimals
    await priceFeed.waitForDeployment();

    // 部署 NFT 合约
    const MyNFT = await ethers.getContractFactory("MyNFT");
    myNFT = await MyNFT.deploy();
    await myNFT.waitForDeployment();

    // 部署支付代币
    const PaymentToken = await ethers.getContractFactory("MockERC20");
    paymentToken = await PaymentToken.deploy("Payment Token", "PAY");
    await paymentToken.waitForDeployment();

    // 用 deployProxy 部署并初始化
    const PriceFeedAuction = await ethers.getContractFactory("PriceFeedAuction");
    auction = await upgrades.deployProxy(
      PriceFeedAuction,
      [
        await myNFT.getAddress(),
        0,
        3600,
        owner.address,
        await paymentToken.getAddress(),
        await priceFeed.getAddress()
      ],
      { initializer: "initialize" }
    );
    await auction.waitForDeployment();

    // 铸造 NFT 并转入拍卖合约
    await myNFT.mint(owner.address);
    await myNFT.connect(owner).transferFrom(owner.address, await auction.getAddress(), 0);
    await paymentToken.mint(bidder1.address, ethers.parseEther("1000"));
    await paymentToken.mint(bidder2.address, ethers.parseEther("1000"));

    // 授权
    await myNFT.connect(owner).setApprovalForAll(await auction.getAddress(), true);
    await paymentToken.connect(bidder1).approve(await auction.getAddress(), ethers.parseEther("1000"));
    await paymentToken.connect(bidder2).approve(await auction.getAddress(), ethers.parseEther("1000"));
  });

  it("should get correct price from Chainlink", async function () {
    const price = await auction.getLatestPrice();
    expect(price).to.equal(200000000); // 2000 USD with 8 decimals
  });

  it("should calculate correct USD value of bids", async function () {
    const tokenAmount = ethers.parseEther("1"); // 1 token
    const usdValue = await auction.getBidInUSD(tokenAmount);
    expect(usdValue).to.equal(2000 * 1e8); // 200000000
  });

  it("should reject bids below minimum USD value", async function () {
    // 尝试出价低于最小USD价值
    const lowBid = ethers.parseEther("0.01"); // 0.01 token = 20 USD
    await expect(
      auction.connect(bidder1).bid(lowBid)
    ).to.be.revertedWith("Bid too low in USD");
  });

  it("should accept bids above minimum USD value", async function () {
    const highBid = ethers.parseEther("1.1"); // 1.1 token = 2200 USD
    await auction.connect(bidder1).bid(highBid);
    expect(await auction.highestBidder()).to.equal(bidder1.address);
  });

  it("should end auction correctly with valid bid", async function () {
    await auction.connect(bidder1).bid(ethers.parseEther("1.1"));

    // 快进时间
    await ethers.provider.send("evm_increaseTime", [3600]);
    await ethers.provider.send("evm_mine");

    await auction.connect(owner).endAuction();
    expect(await myNFT.ownerOf(0)).to.equal(bidder1.address);
    expect(await paymentToken.balanceOf(owner.address)).to.equal(ethers.parseEther("1.1"));
  });
}); 