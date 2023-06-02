// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";

contract ERC20BridgeMediator is Ownable {
    mapping(uint256 => mapping(address => string)) public tokenToSymbol;
    mapping(uint256 => mapping(string => address)) public symbolToToken;

    event AddedToken(string symbol, uint256 chainId, address token);
    event RemovedToken(string symbol, uint256 chainId, address token);

    function mediate(
        uint256 _sourceChain,
        uint256 _destinationChain,
        bytes memory _sourceData
    ) external view returns (bytes memory) {
        (
            uint256 nonce,
            address sender,
            address sourceToken,
            address receiver,
            uint256 transferAmount,
            uint256 fee
        ) = abi.decode(_sourceData, (uint256, address, address, address, uint256, uint256));

        string memory symbol = tokenToSymbol[_sourceChain][sourceToken];
        require(bytes(symbol).length > 0, "ERC20BridgeMediator: can't find token symbol");

        address destinationToken = symbolToToken[_destinationChain][symbol];
        require(destinationToken != address(0), "ERC20BridgeMediator: can't find token by chain and symbol");

        bytes memory destinationData = abi.encode(nonce, sender, destinationToken, receiver, transferAmount, fee);

        return destinationData;
    }

    function addToken(
        string memory symbol,
        uint256 chainId,
        address token
    ) public onlyOwner {
        require(bytes(tokenToSymbol[chainId][token]).length == 0, "ERC20BridgeMediator: symbol already added");

        tokenToSymbol[chainId][token] = symbol;
        symbolToToken[chainId][symbol] = token;

        emit AddedToken(symbol, chainId, token);
    }

    function removeToken(string memory symbol, uint256 chainId) public onlyOwner {
        require(symbolToToken[chainId][symbol] != address(0), "ERC20BridgeMediator: token does not exist");

        address token = symbolToToken[chainId][symbol];
        delete tokenToSymbol[chainId][token];
        delete symbolToToken[chainId][symbol];

        emit RemovedToken(symbol, chainId, token);
    }
}
