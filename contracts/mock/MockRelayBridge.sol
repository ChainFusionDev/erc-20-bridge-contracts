// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../interfaces/IRelayBridge.sol";
import "../interfaces/IERC20Bridge.sol";

contract MockRelayBridge is IRelayBridge {
    mapping(bytes32 => bytes) public sentData;

    mapping(bytes32 => bool) public sent;
    mapping(bytes32 => bool) public executed;
    mapping(bytes32 => bool) public reverted;

    uint256 public nonce;

    event Sent(
        bytes32 hash,
        address appContract,
        uint256 sourceChain,
        uint256 destinationChain,
        bytes data,
        uint256 gasLimit,
        uint256 nonce,
        uint256 value
    );
    event Reverted(bytes32 hash, uint256 sourceChain, uint256 destinationChain);
    event Executed(bytes32 hash, uint256 sourceChain, uint256 destinationChain);

    // solhint-disable-next-line no-empty-blocks
    receive() external payable {}

    function send(
        uint256 destinationChain,
        uint256 gasLimit,
        bytes memory data
    ) external payable {
        bytes32 hash = dataHash(msg.sender, block.chainid, destinationChain, gasLimit, data, nonce);
        require(sentData[hash].length == 0, "RelayBridge: data already send");

        sent[hash] = true;
        sentData[hash] = data;

        emit Sent(hash, msg.sender, block.chainid, destinationChain, data, gasLimit, nonce, msg.value);

        nonce++;
    }

    function revertSend(
        address appContract,
        uint256 destinationChain,
        uint256 gasLimit,
        bytes memory data,
        uint256 _nonce,
        address
    ) external {
        bytes32 hash = dataHash(appContract, block.chainid, destinationChain, gasLimit, data, _nonce);
        require(sent[hash], "RelayBridge: data never sent");
        require(!reverted[hash], "RelayBridge: data already reverted");

        reverted[hash] = true;

        emit Reverted(hash, block.chainid, destinationChain);
    }

    function execute(
        address appContract,
        uint256 sourceChain,
        uint256 gasLimit,
        bytes memory data,
        uint256 _nonce,
        address
    ) external {
        bytes32 hash = dataHash(appContract, sourceChain, block.chainid, gasLimit, data, _nonce);
        require(!executed[hash], "RelayBridge: data already executed");

        executed[hash] = true;

        emit Executed(hash, sourceChain, block.chainid);
    }

    function leaderHistoryLength() external pure returns (uint256) {
        return 0;
    }

    function dataHash(
        address appContract,
        uint256 sourceChain,
        uint256 destinationChain,
        uint256 gasLimit,
        bytes memory data,
        uint256 _nonce
    ) public pure returns (bytes32) {
        return keccak256(abi.encode(appContract, sourceChain, destinationChain, gasLimit, data, _nonce));
    }
}
