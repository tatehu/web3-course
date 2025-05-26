// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IUpgradeableAuction {
    function initialize(
        address nft,
        uint256 tokenId,
        uint256 biddingTime,
        address seller,
        address paymentToken
    ) external;
}