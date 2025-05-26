// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface ICrossChainAuction {
    function initialize(
        address nft,
        uint256 tokenId,
        uint256 biddingTime,
        address seller,
        address paymentToken,
        address router,
        address linkToken,
        uint64 destinationChainSelector,
        address destinationContract
    ) external;
}
