const { expect } = require("chai");

const { ethers } = require("hardhat");



describe("Cross Chain Auction", function () {

  let myNFT, auction, paymentToken, router, linkToken, owner, bidder1, bidder2;

  let sourceChainSelector = 1;

  let destinationChainSelector = 2;



  beforeEach(async function () {

    [owner, bidder1, bidder2] = await ethers.getSigners();



    // 部署 Mock 合约

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



    // 铸造 NFT

    await myNFT.mint(owner.address); // tokenId 0

    await myNFT.mint(owner.address); // tokenId 1



    // 部署支付代币

    const PaymentToken = await ethers.getContractFactory("MockERC20");

    paymentToken = await PaymentToken.deploy("Payment Token", "PAY");

    await paymentToken.waitForDeployment();



    // 部署拍卖合约

    const CrossChainAuction = await ethers.getContractFactory("CrossChainAuction");

    auction = await CrossChainAuction.deploy();

    await auction.waitForDeployment();



    // **授权 NFT 给拍卖合约（必须在 initialize 之前）**

    await myNFT.connect(owner).setApprovalForAll(await auction.getAddress(), true);



    // 初始化拍卖合约

    await auction.initialize(

      await myNFT.getAddress(),

      0,

      3600,

      owner.address,

      await paymentToken.getAddress(),

      await router.getAddress(),

      await linkToken.getAddress(),

      destinationChainSelector,

      await auction.getAddress() // 在测试中，目标合约就是自己

    );



    // 铸造代币

    await paymentToken.mint(bidder1.address, ethers.parseEther("1000"));

    await paymentToken.mint(bidder2.address, ethers.parseEther("1000"));

    await linkToken.mint(await auction.getAddress(), ethers.parseEther("1000"));



    // 授权支付代币

    await paymentToken.connect(bidder1).approve(await auction.getAddress(), ethers.parseEther("1000"));

    await paymentToken.connect(bidder2).approve(await auction.getAddress(), ethers.parseEther("1000"));

  });



  it("should create auction and emit events", async function () {

    const tx = await auction.connect(owner).createAuction(

      await myNFT.getAddress(),

      0,

      3600,

      await paymentToken.getAddress()

    );

    const receipt = await tx.wait();



    // 验证事件

    const auctionCreatedEvent = receipt.logs.find(

      log => log.fragment && log.fragment.name === "AuctionCreated"

    );

    expect(auctionCreatedEvent).to.not.be.undefined;

    expect(auctionCreatedEvent.args.nft).to.equal(await myNFT.getAddress());



    const crossChainMessageSentEvent = receipt.logs.find(

      log => log.fragment && log.fragment.name === "CrossChainMessageSent"

    );

    expect(crossChainMessageSentEvent).to.not.be.undefined;

  });



  it("should receive cross-chain message and create auction", async function () {

    await auction.connect(owner).createAuction(

      await myNFT.getAddress(),

      0,

      3600,

      await paymentToken.getAddress()

    );

    // 验证跨链消息接收

    const crossChainMessageReceivedEvent = await router.getLastMessageReceived();

    expect(crossChainMessageReceivedEvent).to.not.be.undefined;

  });



  it("should allow bidding and ending auction", async function () {

    const tx = await auction.connect(owner).createAuction(

      await myNFT.getAddress(),

      0,

      3600,

      await paymentToken.getAddress()

    );

    const receipt = await tx.wait();

    const auctionCreatedEvent = receipt.logs.find(

      log => log.fragment && log.fragment.name === "AuctionCreated"

    );

    const auctionId = auctionCreatedEvent.args.auctionId;



    // 出价

    await auction.connect(bidder1).bid(auctionId, ethers.parseEther("1"));

    const auctionInfo = await auction.auctions(auctionId);

    expect(auctionInfo.highestBidder).to.equal(bidder1.address);



    // 快进时间

    await ethers.provider.send("evm_increaseTime", [3600]);

    await ethers.provider.send("evm_mine");



    // 结束拍卖

    await auction.connect(owner).endAuction(auctionId);

    expect(await myNFT.ownerOf(0)).to.equal(bidder1.address);

  });



  it("should handle multiple auctions", async function () {

    // 创建多个拍卖

    const tx1 = await auction.connect(owner).createAuction(

      await myNFT.getAddress(),

      0,

      3600,

      await paymentToken.getAddress()

    );

    const receipt1 = await tx1.wait();

    const auctionId1 = receipt1.logs.find(

      log => log.fragment && log.fragment.name === "AuctionCreated"

    ).args.auctionId;



    const tx2 = await auction.connect(owner).createAuction(

      await myNFT.getAddress(),

      1,

      3600,

      await paymentToken.getAddress()

    );

    const receipt2 = await tx2.wait();

    const auctionId2 = receipt2.logs.find(

      log => log.fragment && log.fragment.name === "AuctionCreated"

    ).args.auctionId;



    expect(auctionId1).to.not.equal(auctionId2);



    // 分别出价

    await auction.connect(bidder1).bid(auctionId1, ethers.parseEther("1"));

    await auction.connect(bidder2).bid(auctionId2, ethers.parseEther("2"));



    // 验证出价正确记录

    const info1 = await auction.auctions(auctionId1);

    const info2 = await auction.auctions(auctionId2);

    expect(info1.highestBidder).to.equal(bidder1.address);

    expect(info2.highestBidder).to.equal(bidder2.address);

  });

}); 
