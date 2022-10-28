// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

interface IERC20Bridge {
    function execute(uint256 sourceChain, bytes memory data) external;

    function revertSend(uint256 destinationChain, bytes memory data) external;

    function bridgeAppAddress() external view returns (address);
}
