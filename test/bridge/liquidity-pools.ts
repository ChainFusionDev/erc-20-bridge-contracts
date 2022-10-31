import hre from 'hardhat';
import { expect } from 'chai';
import { utils } from 'ethers';
import { ethers } from 'hardhat';
import { deployBridge, deployBridgeWithMocks } from '../utils/deploy';

describe('LiquidityPools', function () {
  it('should change token manager', async function () {
    const [, v1] = await ethers.getSigners();

    const { tokenManager, liquidityPools } = await deployBridge();
    const newTokenManager = await ethers.getContractAt('TokenManager', tokenManager.address, v1);

    await expect(liquidityPools.setTokenManager(newTokenManager.address))
      .to.emit(liquidityPools, 'TokenManagerUpdated')
      .withArgs(newTokenManager.address);

    expect(await liquidityPools.tokenManager()).to.equal(newTokenManager.address);
  });

  it('should change fee percentage', async function () {
    const newFeePercentage = utils.parseEther('1');

    const { liquidityPools } = await deployBridge();

    await expect(liquidityPools.setFeePercentage(newFeePercentage))
      .to.emit(liquidityPools, 'FeePercentageUpdated')
      .withArgs(newFeePercentage);

    expect(await liquidityPools.feePercentage()).to.equal(newFeePercentage);
  });

  it('should add liquidity for supported token', async function () {
    const [validator] = await ethers.getSigners();
    const amount = utils.parseEther('0.1');
    const untrueAmount = utils.parseEther('1');

    const { mockToken, liquidityPools, tokenManager } = await deployBridgeWithMocks();

    expect(await tokenManager.getType(mockToken.address)).to.equal(1);

    await mockToken.approve(liquidityPools.address, amount);
    await expect(liquidityPools.addLiquidity(mockToken.address, untrueAmount)).to.be.revertedWith(
      'ERC20: insufficient allowance'
    );
    await expect(liquidityPools.addLiquidity(mockToken.address, amount))
      .to.emit(liquidityPools, 'LiquidityAdded')
      .withArgs(mockToken.address, validator.address, amount);

    expect(await liquidityPools.providedLiquidity(mockToken.address)).to.equal(amount);
    expect(await liquidityPools.availableLiquidity(mockToken.address)).to.equal(amount);
    expect((await liquidityPools.liquidityPositions(mockToken.address, validator.address)).balance).to.equal(amount);
  });

  it('should not add liquidity for unsupported token', async function () {
    const amount = utils.parseEther('1');
    const mintAmount = utils.parseEther('1');

    const { mockToken, liquidityPools } = await deployBridgeWithMocks();

    const MockToken = await ethers.getContractFactory('MockToken');
    const mockToken2 = await MockToken.deploy('Token2', 'TKN2', mintAmount);
    await mockToken2.deployed();

    await mockToken.approve(liquidityPools.address, amount);
    await mockToken2.approve(liquidityPools.address, amount);

    await expect(liquidityPools.addLiquidity(mockToken2.address, amount)).to.be.revertedWith(
      'TokenManager: token is not supported'
    );
  });

  it('should remove liquidity', async function () {
    const [validator] = await ethers.getSigners();
    const amount = utils.parseEther('1');

    const { mockToken, liquidityPools, tokenManager } = await deployBridgeWithMocks();

    expect(await tokenManager.getType(mockToken.address)).to.equal(1);

    await mockToken.approve(liquidityPools.address, amount);
    await expect(liquidityPools.addLiquidity(mockToken.address, amount))
      .to.emit(liquidityPools, 'LiquidityAdded')
      .withArgs(mockToken.address, validator.address, amount);

    await expect(liquidityPools.removeLiquidity(mockToken.address, amount))
      .to.emit(liquidityPools, 'LiquidityRemoved')
      .withArgs(mockToken.address, validator.address, amount);

    expect(await liquidityPools.providedLiquidity(mockToken.address)).to.equal(0);
    expect(await liquidityPools.availableLiquidity(mockToken.address)).to.equal(0);
    expect((await liquidityPools.liquidityPositions(mockToken.address, validator.address)).balance).to.equal(0);
  });

  it('should not remove liquidity more than provided', async function () {
    const [validator] = await ethers.getSigners();

    const amount = utils.parseEther('0.1');
    const amountRemove = utils.parseEther('1');

    const { mockToken, liquidityPools, tokenManager } = await deployBridgeWithMocks();

    expect(await tokenManager.getType(mockToken.address)).to.equal(1);

    await mockToken.approve(liquidityPools.address, amount);
    await expect(liquidityPools.addLiquidity(mockToken.address, amount))
      .to.emit(liquidityPools, 'LiquidityAdded')
      .withArgs(mockToken.address, validator.address, amount);

    await expect(liquidityPools.removeLiquidity(mockToken.address, amountRemove)).to.be.revertedWith(
      'LiquidityPools: too much amount'
    );
  });

  it('should transfer token', async function () {
    const [, receiver] = await ethers.getSigners();
    const amount = utils.parseEther('1');

    const { mockChainId, mockToken, liquidityPools, tokenManager, erc20Bridge } = await deployBridgeWithMocks();

    expect(await tokenManager.getType(mockToken.address)).to.equal(1);

    await mockToken.approve(erc20Bridge.address, amount);
    await mockToken.approve(liquidityPools.address, amount);
    await liquidityPools.addLiquidity(mockToken.address, amount);

    await erc20Bridge.deposit(mockToken.address, mockChainId, receiver.address, amount);
  });

  it('should collect fees', async function () {
    const [, liquidityProvider, receiver] = await ethers.getSigners();
    const depositAmount = utils.parseEther('1');
    const tokenFee = utils.parseEther('0.1');
    const validatorReward = utils.parseEther('0.3');
    const liquidityReward = utils.parseEther('0.3');

    const { mockChainId, mockToken, liquidityPools, tokenManager, erc20Bridge, feeManager } =
      await deployBridgeWithMocks();

    const liquidityPoolsProvider = await ethers.getContractAt(
      'LiquidityPools',
      liquidityPools.address,
      liquidityProvider
    );
    const mockTokenProvider = await ethers.getContractAt('MockToken', mockToken.address, liquidityProvider);

    expect(await tokenManager.getType(mockToken.address)).to.equal(1);

    await mockToken.approve(erc20Bridge.address, depositAmount);
    await mockToken.transfer(liquidityProvider.address, depositAmount);

    await mockTokenProvider.approve(liquidityPools.address, depositAmount);
    await liquidityPoolsProvider.addLiquidity(mockToken.address, depositAmount);

    await feeManager.setTokenFee(mockToken.address, tokenFee, validatorReward, liquidityReward);

    await erc20Bridge.deposit(mockToken.address, mockChainId, receiver.address, depositAmount);
    await feeManager.distributeRewards(mockToken.address);

    await expect(liquidityPools.distributeFee(mockToken.address, depositAmount)).to.be.revertedWith(
      'LiquidityPools: only feeManager'
    );
    expect(await liquidityPoolsProvider.collectedFees(mockToken.address)).to.equal(utils.parseEther('0.033'));
  });

  it('should claim rewards', async function () {
    const [sender, receiver] = await ethers.getSigners();
    const depositAmount = utils.parseEther('1');
    const tokenFee = utils.parseEther('0.1');
    const validatorReward = utils.parseEther('0.3');
    const liquidityReward = utils.parseEther('0.3');
    const sourceChainId = 5;
    const gasLimit = (await ethers.provider.getBlock(0)).gasLimit;
    const transferAmount = utils.parseEther('0.99');
    const nullAddress = '0x0000000000000000000000000000000000000000';
    const leader = '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266';
    const nonce = 1;

    const { mockChainId, mockToken, liquidityPools, erc20Bridge, feeManager, mockRelayBridge } =
      await deployBridgeWithMocks();

    const abiCoder = ethers.utils.defaultAbiCoder;
    const data = abiCoder.encode(
      ['address', 'address', 'uint256', 'address', 'uint256'],
      [sender.address, mockToken.address, mockChainId, receiver.address, transferAmount]
    );

    const dataNullAddress = abiCoder.encode(
      ['address', 'address', 'uint256', 'address', 'uint256'],
      [sender.address, mockToken.address, mockChainId, nullAddress, transferAmount]
    );

    await hre.network.provider.request({
      method: "hardhat_impersonateAccount",
      params: [mockRelayBridge.address],
    });
    const signer = await ethers.getSigner(mockRelayBridge.address);
    await (await sender.sendTransaction({
      to: signer.address,
      value: ethers.utils.parseEther('1'),
    })).wait();

    await expect(
      erc20Bridge.connect(signer).execute(mockChainId, data)
    ).to.be.revertedWith('IERC20: amount more than contract balance');

    await mockToken.approve(erc20Bridge.address, depositAmount);
    await mockToken.approve(liquidityPools.address, depositAmount);
    await liquidityPools.addLiquidity(mockToken.address, depositAmount);

    await expect(
      erc20Bridge.connect(signer).execute(mockChainId, dataNullAddress)
    ).to.be.revertedWith('ERC20: transfer to the zero address');
    await expect(feeManager.distributeRewards(mockToken.address)).to.be.revertedWith(
      'LiquidityPools: amount must be greater than zero'
    );
    await expect(liquidityPools.transfer(mockToken.address, receiver.address, depositAmount)).to.be.revertedWith(
      'LiquidityPools: only erc20Bridge'
    );
    await expect(liquidityPools.transferNative(receiver.address, depositAmount)).to.be.revertedWith(
      'LiquidityPools: only erc20Bridge'
    );

    await erc20Bridge.deposit(mockToken.address, sourceChainId, receiver.address, depositAmount);

    await erc20Bridge.connect(signer).execute(mockChainId, data);

    await feeManager.setTokenFee(mockToken.address, tokenFee, validatorReward, liquidityReward);
    await feeManager.distributeRewards(mockToken.address);

    expect(await liquidityPools.collectedFees(mockToken.address)).to.equal(utils.parseEther('0.003'));
    expect(await liquidityPools.rewardsOwing(mockToken.address)).to.equal(utils.parseEther('0.003'));
    await liquidityPools.claimRewards(mockToken.address);
    expect(await liquidityPools.rewardsOwing(mockToken.address)).to.equal(0);
  });

  it('should add and remove liquidity several times', async function () {
    const amount = utils.parseEther('1');
    const amountLiquidity = utils.parseEther('0.001');

    const { mockToken, liquidityPools, erc20Bridge } = await deployBridgeWithMocks();

    await mockToken.approve(erc20Bridge.address, amount);
    await mockToken.approve(liquidityPools.address, amount);
    await liquidityPools.addLiquidity(mockToken.address, amountLiquidity);
    await liquidityPools.addLiquidity(mockToken.address, amountLiquidity);
    await liquidityPools.addLiquidity(mockToken.address, amountLiquidity);

    expect(await liquidityPools.providedLiquidity(mockToken.address)).to.equal(amountLiquidity.mul(3));

    const MockToken = await ethers.getContractFactory('MockToken');
    const mockToken2 = await MockToken.deploy('Token2', 'TKN2', amountLiquidity);
    await expect(liquidityPools.removeLiquidity(mockToken2.address, amountLiquidity)).to.be.revertedWith(
      'TokenManager: token is not supported'
    );

    await liquidityPools.removeLiquidity(mockToken.address, amountLiquidity);
    await liquidityPools.removeLiquidity(mockToken.address, amountLiquidity);
    await liquidityPools.removeLiquidity(mockToken.address, amountLiquidity);
    expect(await liquidityPools.providedLiquidity(mockToken.address)).to.equal(0);
  });

  it('should add and remove liquidity using native currency', async function () {
    const [validator] = await ethers.getSigners();

    const amount = utils.parseEther('1');

    const { liquidityPools, tokenManager } = await deployBridge();
    const NATIVE_TOKEN = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    await tokenManager.setToken(NATIVE_TOKEN, 1);
    expect(await tokenManager.getType(NATIVE_TOKEN)).to.equal(1);

    await expect(liquidityPools.addNativeLiquidity({ value: amount }))
      .to.emit(liquidityPools, 'LiquidityAdded')
      .withArgs(NATIVE_TOKEN, validator.address, amount);

    expect(await liquidityPools.providedLiquidity(NATIVE_TOKEN)).to.equal(amount);
    expect(await liquidityPools.availableLiquidity(NATIVE_TOKEN)).to.equal(amount);
    expect((await liquidityPools.liquidityPositions(NATIVE_TOKEN, validator.address)).balance).to.equal(amount);

    await expect(liquidityPools.removeLiquidity(NATIVE_TOKEN, amount, { value: amount }))
      .to.emit(liquidityPools, 'LiquidityRemoved')
      .withArgs(NATIVE_TOKEN, validator.address, amount);

    expect(await liquidityPools.providedLiquidity(NATIVE_TOKEN)).to.equal(0);
    expect(await liquidityPools.availableLiquidity(NATIVE_TOKEN)).to.equal(0);
    expect((await liquidityPools.liquidityPositions(NATIVE_TOKEN, validator.address)).balance).to.equal(0);
  });

  it('should distribute the reward equally', async function () {
    const [sender, v1, v2, receiver] = await ethers.getSigners();
    const amountTransfer = utils.parseEther('1');
    const amount = utils.parseEther('0.1');
    const tokenFee = utils.parseEther('0.001');
    const validatorReward = utils.parseEther('0.1');
    const liquidityReward = utils.parseEther('0.1');
    const fee = utils.parseEther('0.001');
    const sourceChainId = 5;
    const gasLimit = (await ethers.provider.getBlock(0)).gasLimit;
    const transferAmount = utils.parseEther('0.99');
    const nonce = 1;
    const leader = '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266';

    const { mockChainId, mockToken, liquidityPools, erc20Bridge, feeManager, mockRelayBridge } =
      await deployBridgeWithMocks();

    const abiCoder = ethers.utils.defaultAbiCoder;
    const data = abiCoder.encode(
      ['address', 'address', 'uint256', 'address', 'uint256'],
      [sender.address, mockToken.address, mockChainId, receiver.address, transferAmount]
    );

    await mockToken.transfer(v1.address, amountTransfer);
    await mockToken.transfer(v2.address, amountTransfer);

    const mockToken1 = await ethers.getContractAt('MockToken', mockToken.address, v1);
    const mockToken2 = await ethers.getContractAt('MockToken', mockToken.address, v2);

    const bridge1 = await ethers.getContractAt('ERC20Bridge', erc20Bridge.address, v1);
    const bridge2 = await ethers.getContractAt('ERC20Bridge', erc20Bridge.address, v2);

    const liquidityPools1 = await ethers.getContractAt('LiquidityPools', liquidityPools.address, v1);
    const liquidityPools2 = await ethers.getContractAt('LiquidityPools', liquidityPools.address, v2);

    await mockToken1.approve(liquidityPools1.address, amount);
    await mockToken2.approve(liquidityPools2.address, amount);

    await mockToken1.approve(bridge1.address, amount);
    await mockToken2.approve(bridge2.address, amount);

    await liquidityPools1.addLiquidity(mockToken1.address, amount);
    await liquidityPools2.addLiquidity(mockToken2.address, amount);

    await bridge1.deposit(mockToken1.address, sourceChainId, receiver.address, amount);

    expect(mockRelayBridge.execute(erc20Bridge.address, mockChainId, gasLimit, data, nonce, leader));

    await feeManager.setTokenFee(mockToken.address, tokenFee, validatorReward, liquidityReward);

    const before = await liquidityPools.availableLiquidity(mockToken.address);
    await feeManager.distributeRewards(mockToken.address);
    const after = await liquidityPools.availableLiquidity(mockToken.address);

    expect(after).to.equal(before.add(utils.parseEther('0.001')));

    expect(await liquidityPools.collectedFees(mockToken.address)).to.equal(fee);

    await liquidityPools1.claimRewards(mockToken1.address);

    expect(await liquidityPools.collectedFees(mockToken.address)).to.equal(fee.div(2));

    await liquidityPools2.claimRewards(mockToken2.address);
    expect(await liquidityPools.collectedFees(mockToken.address)).to.equal(0);
  });

  it('should remove liquidity using native token', async function () {
    const [validator] = await ethers.getSigners();

    const amount = utils.parseEther('1');

    const { liquidityPools, tokenManager } = await deployBridge();
    const NATIVE_TOKEN = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';
    await tokenManager.setToken(NATIVE_TOKEN, 1);
    expect(await tokenManager.getType(NATIVE_TOKEN)).to.equal(1);

    await expect(liquidityPools.addNativeLiquidity({ value: amount }))
      .to.emit(liquidityPools, 'LiquidityAdded')
      .withArgs(NATIVE_TOKEN, validator.address, amount);

    expect(await liquidityPools.providedLiquidity(NATIVE_TOKEN)).to.equal(amount);
    expect(await liquidityPools.availableLiquidity(NATIVE_TOKEN)).to.equal(amount);
    expect((await liquidityPools.liquidityPositions(NATIVE_TOKEN, validator.address)).balance).to.equal(amount);

    await expect(liquidityPools.removeLiquidity(NATIVE_TOKEN, amount, { value: amount }))
      .to.emit(liquidityPools, 'LiquidityRemoved')
      .withArgs(NATIVE_TOKEN, validator.address, amount);

    expect(await liquidityPools.providedLiquidity(NATIVE_TOKEN)).to.equal(0);
    expect(await liquidityPools.availableLiquidity(NATIVE_TOKEN)).to.equal(0);
    expect((await liquidityPools.liquidityPositions(NATIVE_TOKEN, validator.address)).balance).to.equal(0);
  });
});
