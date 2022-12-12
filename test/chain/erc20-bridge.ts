import hre from 'hardhat';
import { expect } from 'chai';
import { utils } from 'ethers';
import { ethers, network } from 'hardhat';
import { deployBridge, deployBridgeWithMocks } from '../utils/deploy';

describe('ERC20Bridge', function () {
  it('should change token manager', async function () {
    const { erc20Bridge, tokenManager } = await deployBridge();

    await expect(erc20Bridge.setTokenManager(tokenManager.address))
      .to.emit(erc20Bridge, 'TokenManagerUpdated')
      .withArgs(tokenManager.address);
    expect(await erc20Bridge.tokenManager()).to.equal(tokenManager.address);
  });

  it('should change liquidity pools', async function () {
    const [, v1] = await ethers.getSigners();

    const { erc20Bridge, liquidityPools } = await deployBridge();
    const newLiquidityPools = await ethers.getContractAt('LiquidityPools', liquidityPools.address, v1);

    await expect(erc20Bridge.setLiquidityPools(newLiquidityPools.address))
      .to.emit(erc20Bridge, 'LiquidityPoolsUpdated')
      .withArgs(newLiquidityPools.address);

    expect(await erc20Bridge.liquidityPools()).to.equal(newLiquidityPools.address);
  });

  it('should deposit tokens to bridge', async function () {
    const [sender, receiver] = await ethers.getSigners();

    const depositAmount = utils.parseEther('1');
    const depositAmountZero = utils.parseEther('0');
    const transferAmount = utils.parseEther('0.99');
    const { mockChainId, mockToken, erc20Bridge, liquidityPools, mockRelayBridge, feeManager } =
      await deployBridgeWithMocks();
    const sourceChain = network.config.chainId;
    const gasLimit = (await ethers.provider.getBlock(0)).gasLimit;
    const nonce = 0;

    await mockToken.approve(erc20Bridge.address, depositAmount);
    await mockToken.approve(liquidityPools.address, depositAmount);
    await expect(
      erc20Bridge.deposit(mockToken.address, Number(sourceChain), receiver.address, depositAmountZero)
    ).to.be.revertedWith('ERC20Bridge: cannot deposit on the same chain ID');
    await expect(
      erc20Bridge.deposit(mockToken.address, mockChainId, receiver.address, depositAmountZero)
    ).to.be.revertedWith('Bridge: amount cannot be equal to 0.');
    await erc20Bridge.deposit(mockToken.address, mockChainId, receiver.address, depositAmount);

    const fee = await feeManager.calculateFee(mockToken.address, transferAmount);

    const abiCoder = ethers.utils.defaultAbiCoder;
    const data = abiCoder.encode(
      ['uint256', 'address', 'address', 'address', 'uint256', 'uint256'],
      [nonce, sender.address, mockToken.address, receiver.address, transferAmount, fee]
    );

    const dataForHash = abiCoder.encode(
      ['address', 'uint256', 'uint256', 'uint256', 'bytes', 'uint256'],
      [erc20Bridge.address, sourceChain, mockChainId, gasLimit, data, nonce]
    );
    const hash = ethers.utils.keccak256(dataForHash);

    expect(await mockToken.balanceOf(liquidityPools.address)).to.equal(transferAmount);
    expect(await mockRelayBridge.sentData(hash)).to.equals(data);
  });

  it('should revert deposit', async function () {
    const [sender, receiver] = await ethers.getSigners();

    const initialSupply = utils.parseEther('1');
    const depositAmount = utils.parseEther('1');
    const transferAmount = utils.parseEther('0.99');
    const NATIVE_TOKEN = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const sourceChainId = network.config.chainId;
    const gasLimit = (await ethers.provider.getBlock(0)).gasLimit;
    const nonce = 0;
    const mintNonce = 1;
    const nativeNonce = 2;

    const {
      mockChainId,
      mockToken,
      erc20Bridge,
      liquidityPools,
      mockMintableBurnableToken,
      mockRelayBridge,
      tokenManager,
      feeManager,
    } = await deployBridgeWithMocks();

    const fee = await feeManager.calculateFee(mockToken.address, depositAmount);
    const mintFee = await feeManager.calculateFee(mockMintableBurnableToken.address, depositAmount);
    const nativeFee = await feeManager.calculateFee(NATIVE_TOKEN, depositAmount);

    const abiCoder = ethers.utils.defaultAbiCoder;
    const data = abiCoder.encode(
      ['uint256', 'address', 'address', 'address', 'uint256', 'uint256'],
      [nonce, sender.address, mockToken.address, receiver.address, transferAmount, fee]
    );

    const dataMintableToken = abiCoder.encode(
      ['uint256', 'address', 'address', 'address', 'uint256', 'uint256'],
      [mintNonce, sender.address, mockMintableBurnableToken.address, receiver.address, transferAmount, mintFee]
    );

    const dataNativeToken = abiCoder.encode(
      ['uint256', 'address', 'address', 'address', 'uint256', 'uint256'],
      [nativeNonce, sender.address, NATIVE_TOKEN, receiver.address, transferAmount, nativeFee]
    );

    const hashToken = await mockRelayBridge.dataHash(
      erc20Bridge.address,
      Number(sourceChainId),
      mockChainId,
      gasLimit,
      data,
      nonce
    );

    const hashMintToken = await mockRelayBridge.dataHash(
      erc20Bridge.address,
      Number(sourceChainId),
      mockChainId,
      gasLimit,
      dataMintableToken,
      mintNonce
    );

    const hashNativeToken = await mockRelayBridge.dataHash(
      erc20Bridge.address,
      Number(sourceChainId),
      mockChainId,
      gasLimit,
      dataNativeToken,
      nativeNonce
    );

    const mockSenderBalanceBeforeDeposit = await mockToken.balanceOf(sender.address);
    const mockLiquidityBalanceBeforeDeposit = await mockToken.balanceOf(liquidityPools.address);

    await mockToken.approve(erc20Bridge.address, depositAmount);
    await mockToken.approve(liquidityPools.address, depositAmount);
    await erc20Bridge.deposit(mockToken.address, mockChainId, receiver.address, depositAmount);

    const mockLiquidityBalanceAfterDeposit = await mockToken.balanceOf(liquidityPools.address);
    expect(mockLiquidityBalanceAfterDeposit.sub(mockLiquidityBalanceBeforeDeposit)).to.equal(transferAmount);

    await erc20Bridge.setRelayBridge(sender.address);
    await erc20Bridge.revertSend(mockChainId, data);
    await erc20Bridge.setRelayBridge(mockRelayBridge.address);

    expect(await mockRelayBridge.sent(hashToken)).to.equals(true);

    const mockExpectedBalance = mockSenderBalanceBeforeDeposit.sub(depositAmount).add(transferAmount);
    expect(await mockToken.balanceOf(sender.address)).to.equal(mockExpectedBalance);

    const mintableSenderBalanceBeforeDeposit = await mockMintableBurnableToken.balanceOf(sender.address);

    await mockMintableBurnableToken.mint(sender.address, initialSupply);
    await mockMintableBurnableToken.transferOwnership(erc20Bridge.address);
    await mockMintableBurnableToken.approve(erc20Bridge.address, depositAmount);

    await erc20Bridge.deposit(mockMintableBurnableToken.address, mockChainId, receiver.address, depositAmount);

    await erc20Bridge.setRelayBridge(sender.address);
    await expect(erc20Bridge.revertSend(mockChainId, dataMintableToken))
      .emit(erc20Bridge, 'Reverted')
      .withArgs(
        mintNonce,
        sender.address,
        mockMintableBurnableToken.address,
        mockChainId,
        receiver.address,
        transferAmount,
        mintFee
      );
    await erc20Bridge.setRelayBridge(mockRelayBridge.address);

    expect(await mockRelayBridge.sent(hashMintToken)).to.equals(true);

    const mintableExpectedBalance = mintableSenderBalanceBeforeDeposit
      .add(initialSupply)
      .sub(depositAmount)
      .add(transferAmount);
    expect(await mockMintableBurnableToken.balanceOf(sender.address)).to.equal(mintableExpectedBalance);

    await tokenManager.setToken(NATIVE_TOKEN, 1);
    expect(await tokenManager.getType(NATIVE_TOKEN)).to.equal(1);
    await liquidityPools.addNativeLiquidity({ value: depositAmount });

    const nativeSenderBalanceBeforeDeposit = await ethers.provider.getBalance(sender.address);
    const nativeLiquidityBalanceBeforeDeposit = await ethers.provider.getBalance(liquidityPools.address);

    const nativeDepositTx = await (
      await erc20Bridge.depositNative(mockChainId, receiver.address, {
        value: depositAmount,
      })
    ).wait();
    const depositTxFee = nativeDepositTx.gasUsed.mul(nativeDepositTx.effectiveGasPrice);

    const nativeSenderBalanceAfterDeposit = await ethers.provider.getBalance(sender.address);
    const nativeExpectedBalanceDeposit = nativeSenderBalanceBeforeDeposit.sub(depositAmount.add(depositTxFee));
    expect(nativeSenderBalanceAfterDeposit).to.equal(nativeExpectedBalanceDeposit);

    const nativeLiquidityBalanceAfterDeposit = await ethers.provider.getBalance(liquidityPools.address);
    expect(nativeLiquidityBalanceAfterDeposit.sub(nativeLiquidityBalanceBeforeDeposit)).to.equal(transferAmount);

    const nativeSenderBalanceBeforeRevert = await ethers.provider.getBalance(sender.address);

    const nativeSetRelayTx = await (await erc20Bridge.setRelayBridge(sender.address)).wait();
    const setRelayTxFee = nativeSetRelayTx.gasUsed.mul(nativeSetRelayTx.effectiveGasPrice);

    const nativeRevertTx = await (await erc20Bridge.revertSend(mockChainId, dataNativeToken)).wait();
    const revertTxFee = nativeRevertTx.gasUsed.mul(nativeRevertTx.effectiveGasPrice);

    expect(await mockRelayBridge.sent(hashNativeToken)).to.equals(true);

    const nativeSenderBalanceAfterRevert = await ethers.provider.getBalance(sender.address);
    const nativeExpectedBalanceRevert = nativeSenderBalanceBeforeRevert.add(
      transferAmount.sub(revertTxFee).sub(setRelayTxFee)
    );
    expect(nativeSenderBalanceAfterRevert).to.equal(nativeExpectedBalanceRevert);
  });

  it('should revert only previously made deposits', async function () {
    const [sender, receiver] = await ethers.getSigners();

    const depositAmount = utils.parseEther('1');
    const transferAmount = utils.parseEther('0.99');
    const sourceChain = network.config.chainId;
    const nonce = 0;

    const { mockChainId, erc20Bridge, mockToken, liquidityPools, feeManager, mockRelayBridge } =
      await deployBridgeWithMocks();

    const fee = await feeManager.calculateFee(mockToken.address, transferAmount);

    const abiCoder = ethers.utils.defaultAbiCoder;
    const data = abiCoder.encode(
      ['uint256', 'address', 'address', 'address', 'uint256', 'uint256'],
      [nonce, sender.address, mockToken.address, receiver.address, transferAmount, fee]
    );

    await mockToken.approve(erc20Bridge.address, depositAmount);
    await mockToken.approve(liquidityPools.address, depositAmount);

    await erc20Bridge.setRelayBridge(sender.address);
    await expect(erc20Bridge.revertSend(mockChainId, data)).to.be.revertedWith(
      "ERC20Bridge: can't revert, should be deposited"
    );
    await erc20Bridge.setRelayBridge(mockRelayBridge.address);

    await erc20Bridge.deposit(mockToken.address, mockChainId, receiver.address, depositAmount);

    await erc20Bridge.setRelayBridge(sender.address);

    expect(await erc20Bridge.isReverted(nonce, Number(sourceChain), mockChainId)).to.equal(false);
    await erc20Bridge.revertSend(mockChainId, data);
    expect(await erc20Bridge.isReverted(nonce, Number(sourceChain), mockChainId)).to.equal(true);

    await erc20Bridge.setRelayBridge(sender.address);
    await expect(erc20Bridge.revertSend(mockChainId, data)).to.be.revertedWith('ERC20Bridge: already reverted');
    await erc20Bridge.setRelayBridge(mockRelayBridge.address);
  });

  it('should emit event Reverted native token', async function () {
    const [sender, receiver] = await ethers.getSigners();

    const depositAmount = utils.parseEther('1');
    const transferAmount = utils.parseEther('0.99');
    const NATIVE_TOKEN = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const nonce = 0;

    const { mockChainId, erc20Bridge, liquidityPools, mockRelayBridge, tokenManager, feeManager } =
      await deployBridgeWithMocks();

    const fee = await feeManager.calculateFee(NATIVE_TOKEN, transferAmount);

    const abiCoder = ethers.utils.defaultAbiCoder;
    const dataNativeToken = abiCoder.encode(
      ['uint256', 'address', 'address', 'address', 'uint256', 'uint256'],
      [nonce, sender.address, NATIVE_TOKEN, receiver.address, transferAmount, fee]
    );

    await tokenManager.setToken(NATIVE_TOKEN, 1);
    expect(await tokenManager.getType(NATIVE_TOKEN)).to.equal(1);
    await liquidityPools.addNativeLiquidity({ value: depositAmount });
    await erc20Bridge.depositNative(mockChainId, receiver.address, {
      value: depositAmount,
    });

    await hre.network.provider.request({
      method: 'hardhat_impersonateAccount',
      params: [mockRelayBridge.address],
    });
    const signer = await ethers.getSigner(mockRelayBridge.address);
    await (
      await sender.sendTransaction({
        to: signer.address,
        value: ethers.utils.parseEther('1'),
      })
    ).wait();

    await expect(erc20Bridge.connect(signer).revertSend(mockChainId, dataNativeToken))
      .emit(erc20Bridge, 'Reverted')
      .withArgs(nonce, sender.address, NATIVE_TOKEN, mockChainId, receiver.address, transferAmount, fee);
  });

  it('should execute', async function () {
    const [sender, receiver, user] = await ethers.getSigners();
    const depositAmount = utils.parseEther('1');
    const transferAmount = utils.parseEther('0.99');
    const sourceChain = network.config.chainId;
    const nonce = 0;

    const { mockChainId, mockToken, erc20Bridge, liquidityPools, tokenManager, mockRelayBridge, feeManager } =
      await deployBridgeWithMocks();

    const fee = await feeManager.calculateFee(mockToken.address, transferAmount);

    const abiCoder = ethers.utils.defaultAbiCoder;
    const data = abiCoder.encode(
      ['uint256', 'address', 'address', 'address', 'uint256', 'uint256'],
      [nonce, sender.address, mockToken.address, receiver.address, transferAmount, fee]
    );

    expect(await tokenManager.getType(mockToken.address)).to.equal(1);

    await mockToken.approve(erc20Bridge.address, depositAmount);
    await mockToken.approve(liquidityPools.address, depositAmount);
    await erc20Bridge.deposit(mockToken.address, mockChainId, receiver.address, depositAmount);

    const receiverBalanceBeforeExecute = await mockToken.balanceOf(receiver.address);

    const bridgeUser = await ethers.getContractAt('ERC20Bridge', erc20Bridge.address, user);
    await expect(bridgeUser.execute(erc20Bridge.address, data)).to.be.revertedWith('ERC20Bridge: only RelayBridge');

    await hre.network.provider.request({
      method: 'hardhat_impersonateAccount',
      params: [mockRelayBridge.address],
    });
    const signer = await ethers.getSigner(mockRelayBridge.address);
    await (
      await sender.sendTransaction({
        to: signer.address,
        value: ethers.utils.parseEther('1'),
      })
    ).wait();

    await erc20Bridge.connect(signer).execute(mockChainId, data);

    const receiverBalanceAfterExecute = await mockToken.balanceOf(receiver.address);
    expect(receiverBalanceAfterExecute.sub(receiverBalanceBeforeExecute)).to.equal(transferAmount);

    expect(await erc20Bridge.isExecuted(nonce, mockChainId, Number(sourceChain))).to.equal(true);

    await expect(erc20Bridge.connect(signer).execute(mockChainId, data)).to.be.revertedWith(
      'ERC20Bridge: already executed'
    );
  });

  it('should execute transfer', async function () {
    const [sender, receiver] = await ethers.getSigners();
    const depositAmount = utils.parseEther('1');
    const transferAmount = utils.parseEther('0.99');
    const sourceChain = network.config.chainId;
    const nonce = 0;

    const { mockChainId, mockToken, erc20Bridge, liquidityPools, tokenManager, mockRelayBridge, feeManager } =
      await deployBridgeWithMocks();

    const fee = await feeManager.calculateFee(mockToken.address, transferAmount);

    const abiCoder = ethers.utils.defaultAbiCoder;
    const data = abiCoder.encode(
      ['uint256', 'address', 'address', 'address', 'uint256', 'uint256'],
      [nonce, sender.address, mockToken.address, receiver.address, transferAmount, fee]
    );

    expect(await tokenManager.getType(mockToken.address)).to.equal(1);

    await mockToken.approve(erc20Bridge.address, depositAmount);
    await mockToken.approve(liquidityPools.address, depositAmount);
    await liquidityPools.addLiquidity(mockToken.address, depositAmount);
    await erc20Bridge.deposit(mockToken.address, mockChainId, receiver.address, depositAmount);

    await hre.network.provider.request({
      method: 'hardhat_impersonateAccount',
      params: [mockRelayBridge.address],
    });
    const signer = await ethers.getSigner(mockRelayBridge.address);
    await (
      await sender.sendTransaction({
        to: signer.address,
        value: ethers.utils.parseEther('1'),
      })
    ).wait();

    await expect(erc20Bridge.connect(signer).execute(mockChainId, data))
      .to.emit(erc20Bridge, 'Transferred')
      .withArgs(nonce, sender.address, mockToken.address, mockChainId, receiver.address, transferAmount, fee);

    expect(await erc20Bridge.isExecuted(nonce, mockChainId, Number(sourceChain))).to.equal(true);
  });

  it('should deposit supported tokens to bridge', async function () {
    const [, receiver] = await ethers.getSigners();
    const depositAmount = utils.parseEther('1');
    const mintAmount = utils.parseEther('1');
    const { mockChainId, mockToken, erc20Bridge } = await deployBridgeWithMocks();

    const MockToken = await ethers.getContractFactory('MockToken');
    const mockToken2 = await MockToken.deploy('Token2', 'TKN2', mintAmount);
    await mockToken2.deployed();

    await mockToken.approve(erc20Bridge.address, depositAmount);
    await mockToken2.approve(erc20Bridge.address, depositAmount);

    await erc20Bridge.deposit(mockToken.address, mockChainId, receiver.address, depositAmount);

    await expect(
      erc20Bridge.deposit(mockToken2.address, mockChainId, receiver.address, depositAmount)
    ).to.be.revertedWith('TokenManager: token is not enabled');
  });

  it('should mint and burn tokens', async function () {
    const [sender, receiver] = await ethers.getSigners();
    const depositAmount = utils.parseEther('1');
    const initialSupply = utils.parseEther('1');
    const transferAmount = utils.parseEther('0.99');
    const nonce = 0;

    const { mockChainId, mockMintableBurnableToken, erc20Bridge, mockRelayBridge, feeManager } =
      await deployBridgeWithMocks();

    const fee = await feeManager.calculateFee(mockMintableBurnableToken.address, transferAmount);

    const abiCoder = ethers.utils.defaultAbiCoder;
    const dataMintableToken = abiCoder.encode(
      ['uint256', 'address', 'address', 'address', 'uint256', 'uint256'],
      [nonce, sender.address, mockMintableBurnableToken.address, receiver.address, transferAmount, fee]
    );

    await mockMintableBurnableToken.mint(sender.address, initialSupply);
    await mockMintableBurnableToken.transferOwnership(erc20Bridge.address);

    await hre.network.provider.request({
      method: 'hardhat_impersonateAccount',
      params: [mockRelayBridge.address],
    });
    const signer = await ethers.getSigner(mockRelayBridge.address);
    await (
      await sender.sendTransaction({
        to: signer.address,
        value: ethers.utils.parseEther('1'),
      })
    ).wait();

    await expect(erc20Bridge.connect(signer).execute(mockChainId, dataMintableToken))
      .to.emit(mockMintableBurnableToken, 'Transfer')
      .withArgs('0x0000000000000000000000000000000000000000', receiver.address, transferAmount);

    await mockMintableBurnableToken.approve(erc20Bridge.address, depositAmount);
    await expect(erc20Bridge.deposit(mockMintableBurnableToken.address, mockChainId, receiver.address, depositAmount))
      .emit(mockMintableBurnableToken, 'Transfer')
      .withArgs(sender.address, '0x0000000000000000000000000000000000000000', transferAmount);
  });

  it('should deposit and transfer using native currency', async function () {
    const [sender, receiver] = await ethers.getSigners();

    const depositAmount = utils.parseEther('1');
    const depositAmountZero = utils.parseEther('0');
    const transferAmount = utils.parseEther('0.99');
    const NATIVE_TOKEN = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    const sourceChain = network.config.chainId;
    const nonce = 0;

    const { mockChainId, erc20Bridge, liquidityPools, mockRelayBridge, tokenManager, feeManager } =
      await deployBridgeWithMocks();

    const fee = await feeManager.calculateFee(NATIVE_TOKEN, transferAmount);

    const abiCoder = ethers.utils.defaultAbiCoder;
    const data = abiCoder.encode(
      ['uint256', 'address', 'address', 'address', 'uint256', 'uint256'],
      [nonce, sender.address, NATIVE_TOKEN, receiver.address, transferAmount, fee]
    );

    await tokenManager.setToken(NATIVE_TOKEN, 1);
    expect(await tokenManager.getType(NATIVE_TOKEN)).to.equal(1);

    await liquidityPools.addNativeLiquidity({ value: depositAmount });

    await expect(
      erc20Bridge.depositNative(mockChainId, receiver.address, {
        value: depositAmountZero,
      })
    ).to.be.revertedWith('ERC20Bridge: amount cannot be equal to 0.');
    await erc20Bridge.depositNative(mockChainId, receiver.address, {
      value: depositAmount,
    });

    const balanceReceiverBefore = await ethers.provider.getBalance(receiver.address);

    await hre.network.provider.request({
      method: 'hardhat_impersonateAccount',
      params: [mockRelayBridge.address],
    });
    const signer = await ethers.getSigner(mockRelayBridge.address);
    await (
      await sender.sendTransaction({
        to: signer.address,
        value: ethers.utils.parseEther('1'),
      })
    ).wait();

    await erc20Bridge.connect(signer).execute(mockChainId, data);

    const balanceReceiverAfter = await ethers.provider.getBalance(receiver.address);
    expect(balanceReceiverAfter).to.equal(balanceReceiverBefore.add(transferAmount));

    expect(await erc20Bridge.isExecuted(nonce, mockChainId, Number(sourceChain))).to.equal(true);
  });
});
