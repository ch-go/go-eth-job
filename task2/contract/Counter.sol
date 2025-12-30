// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;
contract Counter {
    uint256 private _count;

    event CountChanged(uint256 newCount);

    function count() external view returns (uint256) {
        return _count;
    }

    function increment() external {
        _count += 1;
        emit CountChanged(_count);
    }

    function decrement() external {
        require(_count > 0, "Counter: below zero");
        _count -= 1;
        emit CountChanged(_count);
    }

    function reset() external {
        _count = 0;
        emit CountChanged(_count);
    }
}
