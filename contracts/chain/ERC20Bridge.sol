// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./SignerOwnable.sol";
import "./TokenManager.sol";
import "./LiquidityPools.sol";
import "./FeeManager.sol";
import "./Globals.sol";
import "../interfaces/IERC20MintableBurnable.sol";
import "../interfaces/IRelayBridge.sol";

contract ERC20Bridge is Initializable, SignerOwnable {
    mapping(bytes32 => bool) public sent;
    mapping(bytes32 => bool) public executed;
    mapping(bytes32 => bool) public reverted;

    address public bridgeAppAddress;
    uint256 public nonce;

    TokenManager public tokenManager;
    LiquidityPools public liquidityPools;
    FeeManager public feeManager;
    IRelayBridge public relayBridge;

    event RelayBridgeUpdated(address _relayBridge);
    event TokenManagerUpdated(address _tokenManager);
    event ValidatorAddressUpdated(address _validatorAddress);
    event LiquidityPoolsUpdated(address _liquidityPools);
    event FeeManagerUpdated(address _feeManager);
    event Deposited(
        uint256 nonce,
        address sender,
        address token,
        uint256 destinationChainId,
        address receiver,
        uint256 amount,
        uint256 fee
    );
    event DepositedNative(
        uint256 nonce,
        address sender,
        address token,
        uint256 destinationChainId,
        address receiver,
        uint256 amount,
        uint256 fee
    );
    event Transferred(
        uint256 nonce,
        address sender,
        address token,
        uint256 sourceChainId,
        address receiver,
        uint256 amount,
        uint256 fee
    );

    event Reverted(
        uint256 nonce,
        address sender,
        address token,
        uint256 destinationChainId,
        address receiver,
        uint256 amount,
        uint256 fee
    );

    modifier onlyRelayBridge() {
        require(msg.sender == address(relayBridge), "ERC20Bridge: only RelayBridge");
        _;
    }

    function initialize(
        address _signerStorage,
        address _tokenManager,
        address payable _liquidityPools,
        address payable _feeManager,
        address _bridgeAppAddress,
        address _relayBridge
    ) external initializer {
        _setSignerStorage(_signerStorage);
        _setTokenManager(_tokenManager);
        _setLiquidityPools(_liquidityPools);
        _setFeeManager(_feeManager);
        bridgeAppAddress = _bridgeAppAddress;
        relayBridge = IRelayBridge(_relayBridge);
    }

    function deposit(
        address _token,
        uint256 _destinationChainId,
        address _receiver,
        uint256 _amount
    ) external {
        require(_destinationChainId != block.chainid, "ERC20Bridge: cannot deposit on the same chain ID");
        require(_amount != 0, "ERC20Bridge: amount cannot be equal to 0");
        require(tokenManager.getType(_token) != TokenType.DISABLED, "TokenManager: token is not enabled");

        uint256 fee = feeManager.calculateFee(_token, _amount);
        require(IERC20(_token).transferFrom(msg.sender, address(feeManager), fee), "IERC20: transfer failed");

        uint256 transferAmount = _amount - fee;

        bytes32 id = this.getDataId(nonce, block.chainid, _destinationChainId);
        sent[id] = true;

        if (tokenManager.getType(_token) == TokenType.MINTED) {
            IERC20MintableBurnable(_token).burnFrom(msg.sender, transferAmount);
        } else {
            require(IERC20(_token).transferFrom(msg.sender, address(this), transferAmount), "IERC20: transfer failed");

            require(IERC20(_token).approve(address(liquidityPools), _amount), "IERC20: approve failed");
            liquidityPools.deposit(_token, transferAmount);
        }

        bytes memory data = abi.encode(nonce, msg.sender, _token, _receiver, transferAmount, fee);

        // solhint-disable-next-line check-send-result
        relayBridge.send(_destinationChainId, block.gaslimit, data);

        emit Deposited(nonce, msg.sender, _token, _destinationChainId, _receiver, _amount, fee);

        nonce++;
    }

    function execute(uint256 _sourceChainId, bytes memory _data) external onlyRelayBridge {
        (uint256 _nonce, address sender, address token, address receiver, uint256 transferAmount, uint256 fee) = abi
            .decode(_data, (uint256, address, address, address, uint256, uint256));

        _executeTransfer(_nonce, sender, token, _sourceChainId, receiver, transferAmount, fee);
    }

    function revertSend(uint256 _destinationChainId, bytes memory _data) external onlyRelayBridge {
        (uint256 _nonce, address sender, address token, address receiver, uint256 amount, uint256 fee) = abi.decode(
            _data,
            (uint256, address, address, address, uint256, uint256)
        );

        require(tokenManager.getType(token) != TokenType.DISABLED, "TokenManager: token is not enabled");

        bytes32 id = this.getDataId(_nonce, block.chainid, _destinationChainId);

        require(sent[id], "ERC20Bridge: can't revert, should be deposited");
        require(!reverted[id], "ERC20Bridge: already reverted");
        reverted[id] = true;

        if (tokenManager.getType(token) == TokenType.MINTED) {
            IERC20MintableBurnable(token).mint(sender, amount);
        } else if (token == NATIVE_TOKEN) {
            liquidityPools.transferNative(sender, amount);
        } else {
            liquidityPools.transfer(token, sender, amount);
        }

        emit Reverted(_nonce, sender, token, _destinationChainId, receiver, amount, fee);
    }

    function setRelayBridge(address _relayBridge) public onlySigner {
        _setRelayBridge(_relayBridge);
    }

    function setTokenManager(address _tokenManager) public onlySigner {
        _setTokenManager(_tokenManager);
    }

    function setLiquidityPools(address payable _liquidityPools) public onlySigner {
        _setLiquidityPools(_liquidityPools);
    }

    function setFeeManager(address payable _feeManager) public onlySigner {
        _setFeeManager(_feeManager);
    }

    function depositNative(uint256 _destinationChainId, address _receiver) public payable {
        uint256 _amount = msg.value;
        require(_amount != 0, "ERC20Bridge: amount cannot be equal to 0");
        require(tokenManager.getType(NATIVE_TOKEN) == TokenType.PROVIDED, "TokenManager: token is not enabled");

        uint256 fee = feeManager.calculateFee(NATIVE_TOKEN, _amount);

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = address(feeManager).call{value: fee, gas: 21000}("");
        require(success, "ERC20Bridge: transfer native token failed");

        uint256 transferAmount = _amount - fee;

        liquidityPools.depositNative{value: transferAmount}();

        bytes32 id = this.getDataId(nonce, block.chainid, _destinationChainId);
        sent[id] = true;

        bytes memory data = abi.encode(nonce, msg.sender, NATIVE_TOKEN, _receiver, transferAmount, fee);

        // solhint-disable-next-line check-send-result
        relayBridge.send(_destinationChainId, block.gaslimit, data);

        emit DepositedNative(nonce, msg.sender, NATIVE_TOKEN, _destinationChainId, _receiver, _amount, fee);

        nonce++;
    }

    function isExecuted(
        uint256 _nonce,
        uint256 _sourceChainId,
        uint256 _destinationChainId
    ) public view returns (bool) {
        bytes32 id = this.getDataId(_nonce, _sourceChainId, _destinationChainId);
        return executed[id];
    }

    function isReverted(
        uint256 _nonce,
        uint256 _sourceChainId,
        uint256 _destinationChainId
    ) public view returns (bool) {
        bytes32 id = this.getDataId(_nonce, _sourceChainId, _destinationChainId);
        return reverted[id];
    }

    function getDataId(
        uint256 _nonce,
        uint256 _sourceChainId,
        uint256 _destinationChainId
    ) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(_nonce, _sourceChainId, _destinationChainId));
    }

    function _setRelayBridge(address _relayBridge) private {
        relayBridge = IRelayBridge(_relayBridge);
        emit RelayBridgeUpdated(_relayBridge);
    }

    function _setTokenManager(address _tokenManager) private {
        tokenManager = TokenManager(_tokenManager);
        emit TokenManagerUpdated(_tokenManager);
    }

    function _setLiquidityPools(address payable _liquidityPools) private {
        liquidityPools = LiquidityPools(_liquidityPools);
        emit LiquidityPoolsUpdated(_liquidityPools);
    }

    function _setFeeManager(address payable _feeManager) private {
        feeManager = FeeManager(_feeManager);
        emit FeeManagerUpdated(_feeManager);
    }

    function _executeTransfer(
        uint256 _nonce,
        address _sender,
        address _token,
        uint256 _sourceChainId,
        address _receiver,
        uint256 _amount,
        uint256 _fee
    ) private {
        require(tokenManager.getType(_token) != TokenType.DISABLED, "TokenManager: token is not enabled");
        bytes32 id = this.getDataId(_nonce, _sourceChainId, block.chainid);

        require(!executed[id], "ERC20Bridge: already executed");
        executed[id] = true;

        if (tokenManager.getType(_token) == TokenType.MINTED) {
            IERC20MintableBurnable(_token).mint(_receiver, _amount);
        } else if (_token == NATIVE_TOKEN) {
            liquidityPools.transferNative(_receiver, _amount);
        } else {
            liquidityPools.transfer(_token, _receiver, _amount);
        }

        emit Transferred(_nonce, _sender, _token, _sourceChainId, _receiver, _amount, _fee);
    }
}
