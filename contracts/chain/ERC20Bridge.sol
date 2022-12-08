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

    TokenManager public tokenManager;
    LiquidityPools public liquidityPools;
    FeeManager public feeManager;
    IRelayBridge public relayBridge;

    event Deposited(
        address sender,
        address token,
        uint256 destinationChainId,
        address receiver,
        uint256 fee,
        uint256 transferAmount
    );
    event DepositedNative(
        address sender,
        address token,
        uint256 destinationChainId,
        address receiver,
        uint256 fee,
        uint256 transferAmount
    );
    event Transferred(address sender, address token, uint256 sourceChainId, address receiver, uint256 amount);
    event RelayBridgeUpdated(address _relayBridge);
    event TokenManagerUpdated(address _tokenManager);
    event ValidatorAddressUpdated(address _validatorAddress);
    event LiquidityPoolsUpdated(address _liquidityPools);
    event FeeManagerUpdated(address _feeManager);
    event Reverted(address sender, address token, uint256 sourceChainId, address receiver, uint256 amount);

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
        uint256 _sourceChainId,
        address _receiver,
        uint256 _amount
    ) external {
        require(_sourceChainId != block.chainid, "ERC20Bridge: cannot deposit on the same chain ID");
        require(_amount != 0, "ERC20Bridge: amount cannot be equal to 0.");
        require(tokenManager.getType(_token) != TokenType.DISABLED, "TokenManager: token is not enabled");

        uint256 fee = feeManager.calculateFee(_token, _amount);
        require(IERC20(_token).transferFrom(msg.sender, address(feeManager), fee), "IERC20: transfer failed");

        uint256 transferAmount = _amount - fee;

        bytes32 id = this.getDataId(msg.sender, _token, _sourceChainId, _receiver, transferAmount, fee);
        sent[id] = true;

        if (tokenManager.getType(_token) == TokenType.MINTED) {
            IERC20MintableBurnable(_token).burnFrom(msg.sender, transferAmount);
        } else {
            require(
                IERC20(_token).transferFrom(msg.sender, address(liquidityPools), transferAmount),
                "IERC20: transfer failed"
            );
        }

        emit Deposited(msg.sender, _token, _sourceChainId, _receiver, fee, transferAmount);

        bytes memory data = abi.encode(msg.sender, _token, _sourceChainId, _receiver, transferAmount, fee);

        // solhint-disable-next-line check-send-result
        relayBridge.send(_sourceChainId, block.gaslimit, data);
    }

    function execute(uint256, bytes memory data) external onlyRelayBridge {
        (
            address _sender,
            address _token,
            uint256 _sourceChainId,
            address _receiver,
            uint256 transferAmount,
            uint256 _fee
        ) = abi.decode(data, (address, address, uint256, address, uint256, uint256));

        _executeTransfer(_sender, _token, _sourceChainId, _receiver, transferAmount, _fee);
    }

    function revertSend(uint256, bytes memory data) external onlyRelayBridge {
        (
            address _sender,
            address _token,
            uint256 _sourceChainId,
            address _receiver,
            uint256 _amount,
            uint256 _fee
        ) = abi.decode(data, (address, address, uint256, address, uint256, uint256));

        require(tokenManager.getType(_token) != TokenType.DISABLED, "TokenManager: token is not enabled");
        bytes32 id = this.getDataId(_sender, _token, _sourceChainId, _receiver, _amount, _fee);

        require(sent[id], "ERC20Bridge: can't revert, should be deposited");
        require(!reverted[id], "ERC20Bridge: already reverted");
        reverted[id] = true;

        if (tokenManager.getType(_token) == TokenType.MINTED) {
            IERC20MintableBurnable(_token).mint(_sender, _amount);
        } else if (_token == NATIVE_TOKEN) {
            liquidityPools.transferNative(_sender, _amount);
        } else {
            liquidityPools.transfer(_token, _sender, _amount);
        }

        emit Reverted(_sender, _token, _sourceChainId, _receiver, _amount);
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

    function depositNative(uint256 _sourceChainId, address _receiver) public payable {
        uint256 _amount = msg.value;
        require(_amount != 0, "ERC20Bridge: amount cannot be equal to 0.");
        require(tokenManager.getType(NATIVE_TOKEN) == TokenType.PROVIDED, "TokenManager: token is not enabled");

        uint256 fee = feeManager.calculateFee(NATIVE_TOKEN, _amount);

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = address(feeManager).call{value: fee, gas: 21000}("");
        require(success, "ERC20Bridge: transfer native token failed");

        uint256 transferAmount = _amount - fee;

        // solhint-disable-next-line avoid-low-level-calls
        (success, ) = address(liquidityPools).call{value: transferAmount, gas: 21000}("");
        require(success, "ERC20Bridge: transfer native token failed");

        bytes32 id = this.getDataId(msg.sender, NATIVE_TOKEN, _sourceChainId, _receiver, transferAmount, fee);
        sent[id] = true;

        emit DepositedNative(msg.sender, NATIVE_TOKEN, _sourceChainId, _receiver, fee, transferAmount);

        bytes memory data = abi.encode(msg.sender, NATIVE_TOKEN, _sourceChainId, _receiver, transferAmount, fee);

        // solhint-disable-next-line check-send-result
        relayBridge.send(_sourceChainId, block.gaslimit, data);
    }

    function isExecuted(
        address _sender,
        address _token,
        uint256 _sourceChainId,
        address _receiver,
        uint256 _amount,
        uint256 _fee
    ) public view returns (bool) {
        bytes32 id = this.getDataId(_sender, _token, _sourceChainId, _receiver, _amount, _fee);
        return executed[id];
    }

    function isReverted(
        address _sender,
        address _token,
        uint256 _sourceChainId,
        address _receiver,
        uint256 _amount,
        uint256 _fee
    ) public view returns (bool) {
        bytes32 id = this.getDataId(_sender, _token, _sourceChainId, _receiver, _amount, _fee);
        return reverted[id];
    }

    function getDataId(
        address _sender,
        address _token,
        uint256 _sourceChainId,
        address _receiver,
        uint256 _amount,
        uint256 _fee
    ) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(_sender, _token, _sourceChainId, _receiver, _amount, _fee));
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
        address _sender,
        address _token,
        uint256 _sourceChainId,
        address _receiver,
        uint256 _amount,
        uint256 _fee
    ) private {
        require(tokenManager.getType(_token) != TokenType.DISABLED, "TokenManager: token is not enabled");
        bytes32 id = this.getDataId(_sender, _token, _sourceChainId, _receiver, _amount, _fee);

        require(!executed[id], "ERC20Bridge: already executed");
        executed[id] = true;

        if (tokenManager.getType(_token) == TokenType.MINTED) {
            IERC20MintableBurnable(_token).mint(_receiver, _amount);
        } else if (_token == NATIVE_TOKEN) {
            liquidityPools.transferNative(_receiver, _amount);
        } else {
            liquidityPools.transfer(_token, _receiver, _amount);
        }

        emit Transferred(_sender, _token, _sourceChainId, _receiver, _amount);
    }
}
