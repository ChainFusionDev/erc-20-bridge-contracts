// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

interface IBridgeAppFactory {
    function createApp() external returns (address);

    function apps(uint256 appId) external view returns (address);
}
