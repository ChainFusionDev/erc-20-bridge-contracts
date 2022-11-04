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
}
