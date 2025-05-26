// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./UpgradeableAuction.sol";

contract UpgradeableAuctionV2 is UpgradeableAuction {
    bool private _paused;

    function pause() public {
        _paused = true;
    }

    function resume() public {
        _paused = false;
    }

    function paused() public view returns (bool) {
        return _paused;
    }

    // 重写 bid 方法，增加暂停判断
    function bid(uint256 amount) public override {
        require(!_paused, "Auction is paused");
        super.bid(amount);
    }
}