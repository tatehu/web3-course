const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("Base Auction", function () {
  let myNFT, auction, paymentToken, owner, bidder1, bidder2;

  beforeEach(async function () {
    [owner, bidder1, bidder2] = await ethers.getSigners();

    // 部署 NFT 合约
    const MyNFT = await ethers.getContractFactory("MyNFT");
    myNFT = await MyNFT.deploy();
    await myNFT.waitForDeployment();

    // 部署支付代币
    const PaymentToken = await ethers.getContractFactory("MockERC20");
    paymentToken = await PaymentToken.deploy("Payment Token", "PAY");
    await paymentToken.waitForDeployment();

    // 部署拍卖合约实现
    const BaseAuction = await ethers.getContractFactory("BaseAuction");
    auction = await BaseAuction.deploy();
    await auction.waitForDeployment();
    await auction.initialize(
      await myNFT.getAddress(),
      0,
      3600,
      owner.address,
      await paymentToken.getAddress()
    );

    // 铸造 NFT 和代币
    await myNFT.mint(owner.address);
    await myNFT.connect(owner).transferFrom(owner.address, await auction.getAddress(), 0);
    await paymentToken.mint(bidder1.address, ethers.parseEther("1000"));
    await paymentToken.mint(bidder2.address, ethers.parseEther("1000"));

    // 授权
    await myNFT.connect(owner).setApprovalForAll(await auction.getAddress(), true);
    await paymentToken.connect(bidder1).approve(await auction.getAddress(), ethers.parseEther("1000"));
    await paymentToken.connect(bidder2).approve(await auction.getAddress(), ethers.parseEther("1000"));
  });

  it("should create auction correctly", async function () {
    expect(await auction.seller()).to.equal(owner.address);
    expect(await auction.nft()).to.equal(await myNFT.getAddress());
    expect(await auction.paymentToken()).to.equal(await paymentToken.getAddress());
    expect(await auction.tokenId()).to.equal(0);
  });

  it("should accept bids", async function () {
    await auction.connect(bidder1).bid(ethers.parseEther("1"));
    expect(await auction.highestBidder()).to.equal(bidder1.address);
    expect(await auction.highestBid()).to.equal(ethers.parseEther("1"));

    await auction.connect(bidder2).bid(ethers.parseEther("2"));
    expect(await auction.highestBidder()).to.equal(bidder2.address);
    expect(await auction.highestBid()).to.equal(ethers.parseEther("2"));
  });

  it("should allow withdrawal of outbid amount", async function () {
    await auction.connect(bidder1).bid(ethers.parseEther("1"));
    await auction.connect(bidder2).bid(ethers.parseEther("2"));

    const balanceBefore = await paymentToken.balanceOf(bidder1.address);
    await auction.connect(bidder1).withdraw();
    const balanceAfter = await paymentToken.balanceOf(bidder1.address);

    expect(balanceAfter - balanceBefore).to.equal(ethers.parseEther("1"));
  });

  it("should end auction correctly", async function () {
    await auction.connect(bidder1).bid(ethers.parseEther("1"));

    // 快进时间
    await ethers.provider.send("evm_increaseTime", [3600]);
    await ethers.provider.send("evm_mine");

    await auction.connect(owner).endAuction();
    expect(await myNFT.ownerOf(0)).to.equal(bidder1.address);
    expect(await paymentToken.balanceOf(owner.address)).to.equal(ethers.parseEther("1"));
  });

  it("should return NFT to seller if no bids", async function () {
    // 快进时间
    await ethers.provider.send("evm_increaseTime", [3600]);
    await ethers.provider.send("evm_mine");

    await auction.connect(owner).endAuction();
    expect(await myNFT.ownerOf(0)).to.equal(owner.address);
  });
}); 