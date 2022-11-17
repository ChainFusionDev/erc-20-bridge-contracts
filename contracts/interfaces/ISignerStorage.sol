// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

interface ISignerStorage {
    function setAddress(address _newSigner) external payable;

    function getAddress() external view returns (address);
}
