// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IPriceFeedAuction {
    function initialize(
        address nft,
        uint256 tokenId,
        uint256 biddingTime,
        address seller,
        address paymentToken,
        address priceFeed
    ) external;
}
