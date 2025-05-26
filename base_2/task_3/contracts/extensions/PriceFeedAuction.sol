// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "../base/BaseAuction.sol";
import "../interfaces/AggregatorV3Interface.sol";
import "../interfaces/IPriceFeedAuction.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract PriceFeedAuction is BaseAuction, IPriceFeedAuction {
    AggregatorV3Interface public priceFeed;

    function initialize(
        address _nft,
        uint256 _tokenId,
        uint256 _biddingTime,
        address _seller,
        address _paymentToken,
        address _priceFeed
    ) public initializer {
        __BaseAuction_init(_nft, _tokenId, _biddingTime, _seller, _paymentToken);
        priceFeed = AggregatorV3Interface(_priceFeed);
    }

    function getLatestPrice() public view returns (uint256) {
        (
            uint80 roundID, 
            int price,
            uint startedAt,
            uint timeStamp,
            uint80 answeredInRound
        ) = priceFeed.latestRoundData();
        require(price > 0, "Invalid price");
        return uint256(price);
    }

    function getBidInUSD(uint256 tokenAmount) public view returns (uint256) {
        uint256 price = getLatestPrice();
        return (tokenAmount * price) / 1e18;
    }

    function bid(uint256 amount) public override {
        uint256 bidInUSD = getBidInUSD(amount);
        require(bidInUSD >= 100 * 1e8, "Bid too low in USD");
        super.bid(amount); // 用父合约名调用
    }
} 