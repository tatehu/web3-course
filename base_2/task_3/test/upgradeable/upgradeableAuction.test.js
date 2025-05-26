const { expect } = require("chai");
const { ethers, upgrades } = require("hardhat");

describe("Upgradeable Auction", function () {
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

    // 用 deployProxy 部署并初始化 UpgradeableAuction
    const UpgradeableAuction = await ethers.getContractFactory("UpgradeableAuction");
    auction = await upgrades.deployProxy(
      UpgradeableAuction,
      [
        await myNFT.getAddress(),
        0,
        3600,
        owner.address,
        await paymentToken.getAddress()
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

  it("should initialize correctly", async function () {
    expect(await auction.seller()).to.equal(owner.address);
    expect(await auction.nft()).to.equal(await myNFT.getAddress());
    expect(await auction.paymentToken()).to.equal(await paymentToken.getAddress());
    expect(await auction.tokenId()).to.equal(0);
  });

  it("should upgrade to new implementation and use new features", async function () {
    // 升级到 V2
    const UpgradeableAuctionV2 = await ethers.getContractFactory("UpgradeableAuctionV2");
    const upgraded = await upgrades.upgradeProxy(auction, UpgradeableAuctionV2);

    // 新功能：pause/resume
    expect(await upgraded.paused()).to.be.false;
    await upgraded.pause();
    expect(await upgraded.paused()).to.be.true;

    // 验证暂停功能
    await expect(
      upgraded.connect(bidder1).bid(ethers.parseEther("1"))
    ).to.be.revertedWith("Auction is paused");

    await upgraded.resume();
    expect(await upgraded.paused()).to.be.false;

    // 验证恢复后可以正常出价
    await upgraded.connect(bidder1).bid(ethers.parseEther("1"));
    expect(await upgraded.highestBidder()).to.equal(bidder1.address);
  });

  it("should maintain state after upgrade", async function () {
    // 升级前出价
    await auction.connect(bidder1).bid(ethers.parseEther("1"));
    const highestBidBefore = await auction.highestBid();
    const highestBidderBefore = await auction.highestBidder();

    // 升级到 V2
    const UpgradeableAuctionV2 = await ethers.getContractFactory("UpgradeableAuctionV2");
    const upgraded = await upgrades.upgradeProxy(auction, UpgradeableAuctionV2);

    // 状态保持
    expect(await upgraded.highestBid()).to.equal(highestBidBefore);
    expect(await upgraded.highestBidder()).to.equal(highestBidderBefore);
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
}); 