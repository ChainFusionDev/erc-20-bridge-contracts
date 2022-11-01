// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

interface IBridgeAppFactory {
    function createApp() external returns (address);

    function getAppAddress(uint256 appId) external view returns (address);
}
