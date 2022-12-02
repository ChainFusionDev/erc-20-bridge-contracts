// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../interfaces/ISignerStorage.sol";

abstract contract SignerOwnable {
    ISignerStorage public signerStorage;

    modifier onlySigner() {
        require(signerStorage.getAddress() == msg.sender, "SignerOwnable: only signer");
        _;
    }

    function _setSignerStorage(address _signerStorage) internal virtual {
        signerStorage = ISignerStorage(_signerStorage);
    }
}
