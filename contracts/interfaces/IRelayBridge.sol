// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

interface IRelayBridge {
    function send(uint256 destinationChain, uint256 gasLimit, bytes memory data) external payable;

    function revertSend(
        address appContract,
        uint256 destinationChain,
        uint256 gasLimit,
        bytes memory data,
        uint256 _nonce,
        address leader
    ) external;

    function execute(
        address appContract,
        uint256 sourceChain,
        uint256 gasLimit,
        bytes memory data,
        uint256 _nonce,
        address leader
    ) external;

    function leaderHistoryLength() external view returns (uint256);

    function sentData(bytes32 data) external view returns (bytes memory);

    function dataHash(
        address appContract,
        uint256 sourceChain,
        uint256 destinationChain,
        uint256 gasLimit,
        bytes memory data,
        uint256 _nonce
    ) external pure returns (bytes32);
}
