// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./MockBridgeApp.sol";

contract MockBridgeAppFactory {
    address[] public apps;

    event BridgeAppCreated(address contractAddress, address owner);

    function createApp() public returns (MockBridgeApp) {
        MockBridgeApp bridgeApp = new MockBridgeApp(msg.sender);
        apps.push(address(bridgeApp));

        emit BridgeAppCreated(address(bridgeApp), msg.sender);

        return bridgeApp;
    }

    function getAppAddress(uint256 appId) external view returns (address) {
        require(apps.length != 0, "BridgeAppFactory: apps cannot be empty");

        return apps[appId];
    }
}
