// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./SignerOwnable.sol";
import "./ERC20Bridge.sol";
import "./Globals.sol";

contract BridgeValidatorFeePool is Initializable, SignerOwnable {
    ERC20Bridge public erc20Bridge;

    mapping(address => uint256) public limitPerToken;

    address public validatorFeeReceiver;
    uint256 public homeChainId;

    event ValidatorFeeReceiverUpdated(address validatorFeeReceiver);
    event LimitPerTokenUpdated(address token, uint256 limit);
    event ERC20BridgeUpdated(address erc20Bridge);
    event Collected(address token, uint256 amount);

    // solhint-disable-next-line no-empty-blocks
    receive() external payable {}

    function initialize(
        address _signerStorage,
        address _erc20Bridge,
        address _validatorFeeReceiver,
        uint256 _homeChainId
    ) external initializer {
        _setSignerStorage(_signerStorage);
        _setERC20Bridge(_erc20Bridge);
        _setValidatorFeeReceiver(_validatorFeeReceiver);
        homeChainId = _homeChainId;
    }

    function setERC20Bridge(address _erc20Bridge) public onlySigner {
        _setERC20Bridge(_erc20Bridge);
    }

    function setValidatorFeeReceiver(address _validatorFeeReceiver) public onlySigner {
        _setValidatorFeeReceiver(_validatorFeeReceiver);
    }

    function setLimitPerToken(address _token, uint256 _limit) public onlySigner {
        limitPerToken[_token] = _limit;
        emit LimitPerTokenUpdated(_token, _limit);
    }

    function collect(address _token) public {
        require(limitPerToken[_token] > 0, "BridgeValidatorFeePool: no limit for this token");

        uint256 balanceAmount;

        if (_token == NATIVE_TOKEN) {
            balanceAmount = address(this).balance;

            require(limitPerToken[_token] < balanceAmount, "BridgeValidatorFeePool: insufficient funds");
            erc20Bridge.depositNative{value: balanceAmount}(homeChainId, validatorFeeReceiver);
        } else {
            balanceAmount = IERC20(_token).balanceOf(address(this));

            require(limitPerToken[_token] < balanceAmount, "BridgeValidatorFeePool: insufficient funds");

            IERC20(_token).approve(address(erc20Bridge), balanceAmount);

            erc20Bridge.deposit(_token, homeChainId, validatorFeeReceiver, balanceAmount);
        }

        emit Collected(_token, balanceAmount);
    }

    function _setERC20Bridge(address _erc20Bridge) private {
        erc20Bridge = ERC20Bridge(_erc20Bridge);
        emit ERC20BridgeUpdated(_erc20Bridge);
    }

    function _setValidatorFeeReceiver(address _validatorFeeReceiver) private {
        validatorFeeReceiver = _validatorFeeReceiver;
        emit ValidatorFeeReceiverUpdated(_validatorFeeReceiver);
    }
}
