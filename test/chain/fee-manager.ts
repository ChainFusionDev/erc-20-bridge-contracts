import { expect } from 'chai';
import { utils } from 'ethers';
import { ethers } from 'hardhat';
import { deployBridgeWithMocks } from '../utils/deploy';

describe('FeeManager', function () {
  it('should set fee by owner', async function () {
    const [, other] = await ethers.getSigners();

    const { mockToken, feeManager } = await deployBridgeWithMocks();
    const tokenFee = utils.parseEther('0.01');
    const validatorReward = utils.parseEther('0.3');
    const liquidityReward = utils.parseEther('0.3');

    await feeManager.setTokenFee(mockToken.address, tokenFee, validatorReward, liquidityReward);

    const otherFeeManager = feeManager.connect(other);
    await expect(
      otherFeeManager.setTokenFee(mockToken.address, tokenFee, validatorReward, liquidityReward)
    ).to.be.revertedWith('Ownable: caller is not the owner');
  });

  it('should check if fee is transferred to FeeManager contract', async function () {
    const [, receiver] = await ethers.getSigners();
    const amount = utils.parseEther('1');
    const fee = utils.parseEther('0.01');

    const { mockToken, mockChainId, liquidityPools, tokenManager, erc20Bridge, feeManager } =
      await deployBridgeWithMocks();

    expect(await tokenManager.getType(mockToken.address)).to.equal(1);

    await mockToken.approve(erc20Bridge.address, amount);
    await mockToken.approve(liquidityPools.address, amount);
    await liquidityPools.addLiquidity(mockToken.address, amount);

    await erc20Bridge.deposit(mockToken.address, mockChainId, receiver.address, amount);
    expect(await mockToken.balanceOf(feeManager.address)).to.equal(fee);
  });

  it('should collect fee in native currency', async function () {
    const [, foundation, receiver, user] = await ethers.getSigners();
    const amount = utils.parseEther('1');
    const tokenFee = utils.parseEther('0.01');
    const validatorReward = utils.parseEther('0.3');
    const liquidityReward = utils.parseEther('0.3');
    const tokenLimit = utils.parseEther('0.1');
    const NATIVE_TOKEN = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const { mockChainId, liquidityPools, tokenManager, erc20Bridge, feeManager, bridgeValidatorFeePool } =
      await deployBridgeWithMocks({
        foundationAddress: foundation.address,
      });

    await tokenManager.setToken(NATIVE_TOKEN, 1);
    expect(await tokenManager.getType(NATIVE_TOKEN)).to.equal(1);

    await feeManager.setTokenFee(NATIVE_TOKEN, tokenFee, validatorReward, liquidityReward);
    await bridgeValidatorFeePool.setLimitPerToken(NATIVE_TOKEN, tokenLimit);

    await liquidityPools.addNativeLiquidity({ value: amount });
    await erc20Bridge.depositNative(mockChainId, receiver.address, {
      value: amount,
    });

    const fee = await feeManager.calculateFee(NATIVE_TOKEN, amount);

    const balance = await ethers.provider.getBalance(feeManager.address);
    expect(balance).to.equal(fee);

    const estimateDeposit = await feeManager.estimateDeposit(NATIVE_TOKEN, amount.sub(fee));
    expect(amount).to.equal(estimateDeposit);

    const balanceValidatorFeePoolBefore = await ethers.provider.getBalance(bridgeValidatorFeePool.address);
    const balanceLiquidityPoolsBefore = await ethers.provider.getBalance(feeManager.liquidityPools());
    const balanceFoundationAddressBefore = await ethers.provider.getBalance(foundation.address);

    const userFeeManager = await ethers.getContractAt('FeeManager', feeManager.address, user);
    await userFeeManager.distributeRewards(NATIVE_TOKEN);

    const balanceValidatorFeePoolAfter = await ethers.provider.getBalance(bridgeValidatorFeePool.address);
    const balanceLiquidityPoolsAfter = await ethers.provider.getBalance(feeManager.liquidityPools());
    const balanceFoundationAddressAfter = await ethers.provider.getBalance(foundation.address);

    let validatorRewards = utils.parseEther('0.006');
    let liquidityRewards = utils.parseEther('0.006');
    let foundationRewards = utils.parseEther('0.008');

    expect(balanceValidatorFeePoolAfter).to.equal(balanceValidatorFeePoolBefore.add(validatorRewards));
    expect(balanceLiquidityPoolsAfter).to.equal(balanceLiquidityPoolsBefore.add(liquidityRewards));
    expect(balanceFoundationAddressAfter).to.equal(balanceFoundationAddressBefore.add(foundationRewards));

    expect(await liquidityPools.collectedFees(NATIVE_TOKEN)).to.equal(liquidityRewards);
    await liquidityPools.claimRewards(NATIVE_TOKEN);
    expect(await liquidityPools.collectedFees(NATIVE_TOKEN)).to.equal(0);
  });
});
