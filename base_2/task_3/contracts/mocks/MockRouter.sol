// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@chainlink/contracts-ccip/contracts/interfaces/IRouterClient.sol";
import "../interfaces/LinkTokenInterface.sol";

contract MockRouter is IRouterClient {
    Client.EVM2AnyMessage public lastMessage;
    bytes32 public lastMessageId;
    uint256 public lastFee;

    function ccipSend(
        uint64,
        Client.EVM2AnyMessage calldata
    ) external payable override returns (bytes32) {
        return bytes32(0);
    }

    function isChainSupported(uint64) external pure override returns (bool) {
        return true;
    }

   // function getSupportedTokens(uint64) external pure override returns (address[] memory) {
   //      address[] memory tokens;
   //      return tokens;
   //  }

    function getFee(
        uint64 destinationChainSelector,
        Client.EVM2AnyMessage memory message
    ) external pure override returns (uint256) {
        return 0;
    }

    function getLastMessageReceived() external view returns (bytes32) {
        return lastMessageId;
    }
} 