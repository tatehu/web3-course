// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@chainlink/contracts-ccip/contracts/interfaces/IRouterClient.sol";
import "../interfaces/LinkTokenInterface.sol";
import "../interfaces/ICrossChainAuction.sol";

contract CrossChainAuction is ICrossChainAuction {
    struct Auction {
        address seller;
        IERC721 nft;
        IERC20 paymentToken;
        uint256 tokenId;
        uint256 endTime;
        address highestBidder;
        uint256 highestBid;
        bool ended;
    }

    mapping(uint256 => Auction) public auctions;
    uint32 public nextAuctionId;
    mapping(address => uint256) public pendingReturns;

    IRouterClient public router;
    LinkTokenInterface public linkToken;
    uint64 public destinationChainSelector;
    address public destinationContract;
    address public owner;

    event AuctionCreated(uint32 auctionId, address seller, address nft, uint256 tokenId);
    event Bid(uint32 auctionId, address bidder, uint256 amount);
    event AuctionEnded(uint32 auctionId, address winner, uint256 amount);
    event CrossChainMessageSent(bytes32 messageId, uint32 auctionId);

    modifier onlyOwner() {
        require(msg.sender == owner, "Not owner");
        _;
    }

    function initialize(
        address _nft,
        uint256 _tokenId,
        uint256 _biddingTime,
        address _seller,
        address _paymentToken,
        address _router,
        address _linkToken,
        uint64 _destinationChainSelector,
        address _destinationContract
    ) public {
        require(owner == address(0), "Already initialized");
        owner = _seller;
        router = IRouterClient(_router);
        linkToken = LinkTokenInterface(_linkToken);
        destinationChainSelector = _destinationChainSelector;
        destinationContract = _destinationContract;
    }

    function createAuction(
        address _nft,
        uint256 _tokenId,
        uint256 _biddingTime,
        address _paymentToken
    ) public {
        _createAuction(_nft, _tokenId, _biddingTime, msg.sender, _paymentToken);
    }

    function _createAuction(
        address _nft,
        uint256 _tokenId,
        uint256 _biddingTime,
        address _seller,
        address _paymentToken
    ) internal {
        uint32 auctionId = nextAuctionId++;
        auctions[auctionId] = Auction({
            seller: _seller,
            nft: IERC721(_nft),
            paymentToken: IERC20(_paymentToken),
            tokenId: _tokenId,
            endTime: block.timestamp + _biddingTime,
            highestBidder: address(0),
            highestBid: 0,
            ended: false
        });

        IERC721(_nft).transferFrom(_seller, address(this), _tokenId);

        emit AuctionCreated(auctionId, _seller, _nft, _tokenId);

        bytes memory message = abi.encode(_seller, _nft, _tokenId, _biddingTime, _paymentToken);
        bytes32 messageId = router.ccipSend(
            destinationChainSelector,
            Client.EVM2AnyMessage({
                receiver: abi.encode(destinationContract),
                data: message,
                tokenAmounts: new Client.EVMTokenAmount[](0),
                extraArgs: "",
                feeToken: address(linkToken)
            })
        );
        emit CrossChainMessageSent(messageId, auctionId);
    }

    function bid(uint32 auctionId, uint256 amount) external {
        Auction storage auction = auctions[auctionId];
        require(block.timestamp < auction.endTime, "Auction ended");
        require(amount > auction.highestBid, "Bid too low");

        if (auction.highestBidder != address(0)) {
            pendingReturns[auction.highestBidder] += auction.highestBid;
        }

        auction.paymentToken.transferFrom(msg.sender, address(this), amount);
        auction.highestBidder = msg.sender;
        auction.highestBid = amount;
        emit Bid(auctionId, msg.sender, amount);
    }

    function withdraw() external {
        uint256 amount = pendingReturns[msg.sender];
        require(amount > 0, "No funds");
        pendingReturns[msg.sender] = 0;
        IERC20 token = auctions[0].paymentToken;
        token.transfer(msg.sender, amount);
    }

    function endAuction(uint32 auctionId) external {
        Auction storage auction = auctions[auctionId];
        require(block.timestamp >= auction.endTime, "Auction not yet ended");
        require(!auction.ended, "Already ended");
        auction.ended = true;

        if (auction.highestBidder != address(0)) {
            auction.nft.transferFrom(address(this), auction.highestBidder, auction.tokenId);
            auction.paymentToken.transfer(auction.seller, auction.highestBid);
        } else {
            auction.nft.transferFrom(address(this), auction.seller, auction.tokenId);
        }
        emit AuctionEnded(auctionId, auction.highestBidder, auction.highestBid);
    }
} 