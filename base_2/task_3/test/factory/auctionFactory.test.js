const { expect } = require("chai");
const { ethers } = require("hardhat");
const { upgrades } = require("hardhat");

describe("Auction Factory", function () {
  let factory, myNFT, paymentToken, priceFeed, router, linkToken, owner, bidder1, bidder2;

  beforeEach(async function () {
    [owner, bidder1, bidder2] = await ethers.getSigners();

    // 部署 Mock 合约
    const MockPriceFeed = await ethers.getContractFactory("MockV3Aggregator");
    priceFeed = await MockPriceFeed.deploy(8, 200000000); // 2000 USD with 8 decimals
    await priceFeed.waitForDeployment();

    const MockRouter = await ethers.getContractFactory("MockRouter");
    router = await MockRouter.deploy();
    await router.waitForDeployment();

    const MockLinkToken = await ethers.getContractFactory("MockERC20");
    linkToken = await MockLinkToken.deploy("Chainlink", "LINK");
    await linkToken.waitForDeployment();

    // 部署 NFT 合约
    const MyNFT = await ethers.getContractFactory("MyNFT");
    myNFT = await MyNFT.deploy();
    await myNFT.waitForDeployment();

    // 部署支付代币
    const PaymentToken = await ethers.getContractFactory("MockERC20");
    paymentToken = await PaymentToken.deploy("Payment Token", "PAY");
    await paymentToken.waitForDeployment();

    // 部署实现合约
    const BaseAuction = await ethers.getContractFactory("BaseAuction");
    const baseAuctionImpl = await BaseAuction.deploy();
    await baseAuctionImpl.waitForDeployment();

    const PriceFeedAuction = await ethers.getContractFactory("PriceFeedAuction");
    const priceFeedAuctionImpl = await PriceFeedAuction.deploy();
    await priceFeedAuctionImpl.waitForDeployment();

    const UpgradeableAuction = await ethers.getContractFactory("UpgradeableAuction");
    const upgradeableAuctionImpl = await UpgradeableAuction.deploy();
    await upgradeableAuctionImpl.waitForDeployment();

    const CrossChainAuction = await ethers.getContractFactory("CrossChainAuction");
    const crossChainAuctionImpl = await CrossChainAuction.deploy();
    await crossChainAuctionImpl.waitForDeployment();

    // 部署工厂合约，传入4个实现合约地址
    const AuctionFactory = await ethers.getContractFactory("AuctionFactory");
    factory = await AuctionFactory.deploy(
      await baseAuctionImpl.getAddress(),
      await priceFeedAuctionImpl.getAddress(),
      await upgradeableAuctionImpl.getAddress(),
      await crossChainAuctionImpl.getAddress()
    );
    await factory.waitForDeployment();

    // 铸造 NFT 和代币
    await myNFT.mint(owner.address);
    await paymentToken.mint(bidder1.address, ethers.parseEther("1000"));
    await paymentToken.mint(bidder2.address, ethers.parseEther("1000"));
  });

  describe("Base Auction Creation", function () {
    it("should create base auction correctly", async function () {
      const tx = await factory.createBaseAuction(
        await myNFT.getAddress(),
        0,
        3600,
        await paymentToken.getAddress()
      );
      const receipt = await tx.wait();

      const event = receipt.logs.find(
        log => log.fragment && log.fragment.name === "AuctionCreated"
      );
      expect(event).to.not.be.undefined;
      expect(event.args.auctionType).to.equal("Base");
      expect(event.args.seller).to.equal(owner.address);

      const auction = await ethers.getContractAt("BaseAuction", event.args.auction);
      expect(await auction.seller()).to.equal(owner.address);
      expect(await auction.nft()).to.equal(await myNFT.getAddress());
      expect(await auction.paymentToken()).to.equal(await paymentToken.getAddress());
    });
  });

  describe("Price Feed Auction Creation", function () {
    it("should create price feed auction correctly", async function () {
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
      expect(event).to.not.be.undefined;
      expect(event.args.auctionType).to.equal("PriceFeed");
      expect(event.args.seller).to.equal(owner.address);

      const auction = await ethers.getContractAt("PriceFeedAuction", event.args.auction);
      expect(await auction.seller()).to.equal(owner.address);
      expect(await auction.nft()).to.equal(await myNFT.getAddress());
      expect(await auction.paymentToken()).to.equal(await paymentToken.getAddress());
      expect(await auction.priceFeed()).to.equal(await priceFeed.getAddress());
    });
  });

  describe("Upgradeable Auction Creation", function () {
    it("should create upgradeable auction correctly", async function () {
      const tx = await factory.createUpgradeableAuction(
        await myNFT.getAddress(),
        0,
        3600,
        await paymentToken.getAddress()
      );
      const receipt = await tx.wait();

      const event = receipt.logs.find(
        log => log.fragment && log.fragment.name === "AuctionCreated"
      );
      expect(event).to.not.be.undefined;
      expect(event.args.auctionType).to.equal("Upgradeable");
      expect(event.args.seller).to.equal(owner.address);

      const auction = await ethers.getContractAt("UpgradeableAuction", event.args.auction);
      expect(await auction.seller()).to.equal(owner.address);
      expect(await auction.nft()).to.equal(await myNFT.getAddress());
      expect(await auction.paymentToken()).to.equal(await paymentToken.getAddress());
    });
  });

  describe("Cross Chain Auction Creation", function () {
    it("should create cross chain auction correctly", async function () {
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
      expect(event).to.not.be.undefined;
      expect(event.args.auctionType).to.equal("CrossChain");
      expect(event.args.seller).to.equal(owner.address);

      const auction = await ethers.getContractAt("CrossChainAuction", event.args.auction);
      const seller = await auction.seller();
      expect(seller).to.equal(owner.address);
      expect(await auction.nft()).to.equal(await myNFT.getAddress());
      expect(await auction.paymentToken()).to.equal(await paymentToken.getAddress());
    });
  });

  describe("Multiple Auction Creation", function () {
    it("should create multiple auctions of different types", async function () {
      // 创建基础拍卖
      const tx1 = await factory.createBaseAuction(
        await myNFT.getAddress(),
        0,
        3600,
        await paymentToken.getAddress()
      );
      const receipt1 = await tx1.wait();
      const event1 = receipt1.logs.find(
        log => log.fragment && log.fragment.name === "AuctionCreated"
      );

      // 创建价格预言机拍卖
      const tx2 = await factory.createPriceFeedAuction(
        await myNFT.getAddress(),
        1,
        3600,
        await paymentToken.getAddress(),
        await priceFeed.getAddress()
      );
      const receipt2 = await tx2.wait();
      const event2 = receipt2.logs.find(
        log => log.fragment && log.fragment.name === "AuctionCreated"
      );

      // 验证不同拍卖地址
      expect(event1.args.auction).to.not.equal(event2.args.auction);
      expect(event1.args.auctionType).to.equal("Base");
      expect(event2.args.auctionType).to.equal("PriceFeed");
    });
  });
});



