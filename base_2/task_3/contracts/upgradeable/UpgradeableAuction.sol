// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "../interfaces/IUpgradeableAuction.sol";

contract UpgradeableAuction is Initializable, UUPSUpgradeable, OwnableUpgradeable, IUpgradeableAuction {
    address public seller;
    IERC721 public nft;
    IERC20 public paymentToken;
    uint256 public tokenId;
    uint256 public endTime;
    address public highestBidder;
    uint256 public highestBid;
    bool public ended;

    mapping(address => uint256) public pendingReturns;

    event Bid(address indexed bidder, uint256 amount);
    event AuctionEnded(address winner, uint256 amount);

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(
        address _nft,
        uint256 _tokenId,
        uint256 _biddingTime,
        address _seller,
        address _paymentToken
    ) public initializer {
        __UUPSUpgradeable_init();
        __Ownable_init(msg.sender);
        
        seller = _seller;
        nft = IERC721(_nft);
        paymentToken = IERC20(_paymentToken);
        tokenId = _tokenId;
        endTime = block.timestamp + _biddingTime;
    }

    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    function bid(uint256 amount) public virtual {
        require(block.timestamp < endTime, "Auction ended");
        require(amount > highestBid, "Bid too low");

        if (highestBid != 0) {
            pendingReturns[highestBidder] += highestBid;
        }

        paymentToken.transferFrom(msg.sender, address(this), amount);
        highestBidder = msg.sender;
        highestBid = amount;
        emit Bid(msg.sender, amount);
    }

    function withdraw() external {
        uint256 amount = pendingReturns[msg.sender];
        require(amount > 0, "No funds");
        pendingReturns[msg.sender] = 0;
        paymentToken.transfer(msg.sender, amount);
    }

    function endAuction() external {
        require(block.timestamp >= endTime, "Auction not yet ended");
        require(!ended, "Already ended");
        ended = true;

        if (highestBidder != address(0)) {
            nft.transferFrom(address(this), highestBidder, tokenId);
            paymentToken.transfer(seller, highestBid);
        } else {
            nft.transferFrom(address(this), seller, tokenId);
        }
        emit AuctionEnded(highestBidder, highestBid);
    }
} 