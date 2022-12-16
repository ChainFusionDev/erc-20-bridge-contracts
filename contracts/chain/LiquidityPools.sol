// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "./SignerOwnable.sol";
import "./TokenManager.sol";
import "./ERC20Bridge.sol";
import "./FeeManager.sol";
import "./Globals.sol";
import "hardhat/console.sol";

contract LiquidityPools is Initializable, Ownable, SignerOwnable {
    struct LiquidityPosition {
        uint256 balance;
        uint256 lastRewardPoints;
    }

    TokenManager public tokenManager;
    ERC20Bridge public erc20Bridge;
    FeeManager public feeManager;

    uint256 public feePercentage;

    mapping(address => uint256) public providedLiquidity;
    mapping(address => uint256) public availableLiquidity;
    mapping(address => mapping(address => LiquidityPosition)) public liquidityPositions;
    mapping(address => uint256) public collectedFees;
    mapping(address => uint256) public totalRewardPoints;

    event TokenManagerUpdated(address tokenManager);
    event ERC20BridgeUpdated(address erc20Bridge);
    event FeeManagerUpdated(address feeManager);
    event FeePercentageUpdated(uint256 feePercentage);

    event LiquidityAdded(address token, address account, uint256 amount);
    event LiquidityRemoved(address token, address account, uint256 amount);

    modifier onlyERC20Bridge() {
        require(msg.sender == address(erc20Bridge), "LiquidityPools: only erc20Bridge");
        _;
    }

    modifier onlyFeeManager() {
        require(msg.sender == address(feeManager), "LiquidityPools: only feeManager");
        _;
    }

    // solhint-disable-next-line no-empty-blocks
    receive() external payable {}

    function initialize(
        address _signerStorage,
        address _tokenManager,
        address _erc20Bridge,
        address payable _feeManager,
        uint256 _feePercentage
    ) external initializer {
        _setSignerStorage(_signerStorage);
        _setTokenManager(_tokenManager);
        _setERC20Bridge(_erc20Bridge);
        _setFeeManager(_feeManager);
        _setFeePercentage(_feePercentage);
    }

    function transfer(
        address _token,
        address _receiver,
        uint256 _transferAmount
    ) external onlyERC20Bridge {
        require(
            IERC20(_token).balanceOf(address(this)) >= _transferAmount,
            "IERC20: amount more than contract balance"
        );

        availableLiquidity[_token] -= _transferAmount;

        require(ERC20(_token).transfer(_receiver, _transferAmount), "ERC20: transfer failed");
    }

    function transferNative(address _receiver, uint256 _amount) external onlyERC20Bridge {
        availableLiquidity[NATIVE_TOKEN] -= _amount;

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = _receiver.call{value: _amount, gas: 21000}("");
        require(success, "LiquidityPools: transfer native token failed");
    }

    function deposit(address _token, uint256 _amount) external onlyERC20Bridge {
        require(IERC20(_token).transferFrom(msg.sender, address(this), _amount), "IERC20: transfer failed");
        availableLiquidity[_token] += _amount;
    }

    function depositNative(address _token, uint256 _amount) external payable onlyERC20Bridge {
        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = address(this).call{value: _amount, gas: 21000}("");
        require(success, "LiquidityPools: transfer native token failed");

        availableLiquidity[_token] += _amount;
    }

    function distributeFee(address _token, uint256 _amount) external onlyFeeManager {
        require(_amount > 0, "LiquidityPools: amount must be greater than zero");
        totalRewardPoints[_token] += (_amount * BASE_DIVISOR) / providedLiquidity[_token];
        providedLiquidity[_token] += _amount;
        collectedFees[_token] += _amount;
        availableLiquidity[_token] += _amount;
    }

    function setTokenManager(address _tokenManager) public onlySigner {
        _setTokenManager(_tokenManager);
    }

    function setERC20Bridge(address _erc20Bridge) public onlySigner {
        _setERC20Bridge(_erc20Bridge);
    }

    function setFeeManager(address payable _feeManager) public onlySigner {
        _setFeeManager(_feeManager);
    }

    function setFeePercentage(uint256 _feePercentage) public onlyOwner {
        _setFeePercentage(_feePercentage);
    }

    function addLiquidity(address _token, uint256 _amount) public {
        claimRewards(_token);

        require(tokenManager.getType(_token) == TokenType.PROVIDED, "TokenManager: token is not supported");
        require(IERC20(_token).transferFrom(msg.sender, address(this), _amount), "IERC20: transfer failed");

        _addLiquidity(_token, _amount);
    }

    function removeLiquidity(address _token, uint256 _amount) public payable {
        claimRewards(_token);

        require(tokenManager.getType(_token) == TokenType.PROVIDED, "TokenManager: token is not supported");
        require(liquidityPositions[_token][msg.sender].balance >= _amount, "LiquidityPools: too much amount");

        _removeLiquidity(_token, _amount);

        if (_token == NATIVE_TOKEN) {
            // solhint-disable-next-line avoid-low-level-calls
            (bool success, ) = msg.sender.call{value: _amount, gas: 21000}("");
            require(success, "LiquidityPools: transfer native token failed");
        } else {
            require(IERC20(_token).balanceOf(address(this)) >= _amount, "IERC20: amount more than contract balance");
            require(IERC20(_token).transfer(msg.sender, _amount), "IERC20: transfer failed");
        }
    }

    function claimRewards(address _token) public {
        uint256 rewardsOwingAmount = rewardsOwing(_token);
        if (rewardsOwingAmount > 0) {
            collectedFees[_token] -= rewardsOwingAmount;
            liquidityPositions[_token][msg.sender].balance += rewardsOwingAmount;
            liquidityPositions[_token][msg.sender].lastRewardPoints = totalRewardPoints[_token];
        }
    }

    function addNativeLiquidity() public payable {
        claimRewards(NATIVE_TOKEN);

        _addLiquidity(NATIVE_TOKEN, msg.value);
    }

    function rewardsOwing(address _token) public view returns (uint256) {
        uint256 newRewardPoints = totalRewardPoints[_token] - liquidityPositions[_token][msg.sender].lastRewardPoints;
        return (liquidityPositions[_token][msg.sender].balance * newRewardPoints) / BASE_DIVISOR;
    }

    function _setTokenManager(address _tokenManager) private {
        tokenManager = TokenManager(_tokenManager);
        emit TokenManagerUpdated(_tokenManager);
    }

    function _setERC20Bridge(address _erc20Bridge) private {
        erc20Bridge = ERC20Bridge(_erc20Bridge);
        emit ERC20BridgeUpdated(_erc20Bridge);
    }

    function _setFeeManager(address payable _feeManager) private {
        feeManager = FeeManager(_feeManager);
        emit FeeManagerUpdated(_feeManager);
    }

    function _setFeePercentage(uint256 _feePercentage) private {
        feePercentage = _feePercentage;
        emit FeePercentageUpdated(_feePercentage);
    }

    function _addLiquidity(address _token, uint256 _amount) private {
        providedLiquidity[_token] += _amount;
        availableLiquidity[_token] += _amount;
        liquidityPositions[_token][msg.sender].balance += _amount;
        liquidityPositions[_token][msg.sender].lastRewardPoints = totalRewardPoints[_token];

        emit LiquidityAdded(_token, msg.sender, _amount);
    }

    function _removeLiquidity(address _token, uint256 _amount) private {
        providedLiquidity[_token] -= _amount;
        availableLiquidity[_token] -= _amount;
        liquidityPositions[_token][msg.sender].balance -= _amount;

        emit LiquidityRemoved(_token, msg.sender, _amount);
    }
}
