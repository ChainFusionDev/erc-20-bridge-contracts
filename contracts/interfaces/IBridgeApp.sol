// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";

interface IBridgeApp {
    function setContractAddress(uint256 chainId, address contractAddress) external;

    function setMediator(address _mediatorAddress) external;
}
