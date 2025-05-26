// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/proxy/Clones.sol";
import "../interfaces/IAuction.sol";
import "../interfaces/IPriceFeedAuction.sol";
import "../interfaces/IUpgradeableAuction.sol";
import "../interfaces/ICrossChainAuction.sol";

contract AuctionFactory {
    address public baseAuctionImpl;
    address public priceFeedAuctionImpl;
    address public upgradeableAuctionImpl;
    address public crossChainAuctionImpl;

    event AuctionCreated(address indexed auction, address indexed seller, string auctionType);

    constructor(
        address _baseAuctionImpl,
        address _priceFeedAuctionImpl,
        address _upgradeableAuctionImpl,
        address _crossChainAuctionImpl
    ) {
        baseAuctionImpl = _baseAuctionImpl;
        priceFeedAuctionImpl = _priceFeedAuctionImpl;
        upgradeableAuctionImpl = _upgradeableAuctionImpl;
        crossChainAuctionImpl = _crossChainAuctionImpl;
    }

    function createBaseAuction(
        address nft,
        uint256 tokenId,
        uint256 biddingTime,
        address paymentToken
    ) external returns (address) {
        address clone = Clones.clone(baseAuctionImpl);
        IAuction(clone).initialize(
            nft,
            tokenId,
            biddingTime,
            msg.sender,
            paymentToken
        );
        emit AuctionCreated(clone, msg.sender, "Base");
        return clone;
    }

    function createPriceFeedAuction(
        address nft,
        uint256 tokenId,
        uint256 biddingTime,
        address paymentToken,
        address priceFeed
    ) external returns (address) {
        address clone = Clones.clone(priceFeedAuctionImpl);
        IPriceFeedAuction(clone).initialize(
            nft,
            tokenId,
            biddingTime,
            msg.sender,
            paymentToken,
            priceFeed
        );
        emit AuctionCreated(clone, msg.sender, "PriceFeed");
        return clone;
    }

    function createUpgradeableAuction(
        address nft,
        uint256 tokenId,
        uint256 biddingTime,
        address paymentToken
    ) external returns (address) {
        address clone = Clones.clone(upgradeableAuctionImpl);
        IUpgradeableAuction(clone).initialize(
            nft,
            tokenId,
            biddingTime,
            msg.sender,
            paymentToken
        );
        emit AuctionCreated(clone, msg.sender, "Upgradeable");
        return clone;
    }

    function createCrossChainAuction(
        address nft,
        uint256 tokenId,
        uint256 biddingTime,
        address paymentToken,
        address router,
        address linkToken,
        uint64 destinationChainSelector,
        address destinationContract
    ) external returns (address) {
        address clone = Clones.clone(crossChainAuctionImpl);
        ICrossChainAuction(clone).initialize(
            nft,
            tokenId,
            biddingTime,
            msg.sender,
            paymentToken,
            router,
            linkToken,
            destinationChainSelector,
            destinationContract
        );
        emit AuctionCreated(clone, msg.sender, "CrossChain");
        return clone;
    }
}