// SPDX-License-Identifier: MIT
// contracts/interfaces/IAuction.sol
pragma solidity ^0.8.20;

interface IAuction {
    function initialize(
        address nft,
        uint256 tokenId,
        uint256 biddingTime,
        address seller,
        address paymentToken
    ) external;
}