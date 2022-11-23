import { BigNumber } from 'ethers';
import { ethers, config, network } from 'hardhat';
import { HttpNetworkConfig } from 'hardhat/types';
import { Token } from '../../typechain';

const defaultTokenDeploymentParameters: TokenDeploymentParameters = {
  amount: ethers.utils.parseEther('10000'),
  liquidityAmount: ethers.utils.parseEther('1000'),
  tokenFee: ethers.utils.parseEther('0.001'),
  tokenLimit: ethers.utils.parseEther('0.0001'),
  validatorReward: ethers.utils.parseEther('0.001'),
  liquidityReward: ethers.utils.parseEther('0.001'),
};

export async function deployToken(options: TokenDeploymentOptions): Promise<Token> {
  const params = resolveParameters(options);
  const chainId = network.config.chainId;

  const Token = await ethers.getContractFactory('Token');
  const token = await Token.deploy(options.name, options.symbol, params.amount);
  await token.deployed();

  console.log(`\nDeploying ${options.symbol} token with address ${token.address}`);
  console.log(`Deploying ${options.symbol} with chainId ${chainId} \n`);

  return token;
}

export async function setTokenToTokenManager(tokenAddress: string, options: TokenDeploymentOptions): Promise<void> {
  if (options.tokenManagerAddress !== undefined) {
    console.log('Setting token in TokenManager');

    const tokenManager = await ethers.getContractAt('TokenManager', options.tokenManagerAddress);
    await (await tokenManager.setToken(tokenAddress, '1')).wait();

    console.log('Successfully seted token \n');
  }
}

export async function setTokenFee(tokenAddress: string, options: TokenDeploymentOptions): Promise<void> {
  if (options.feeManagerAddress !== undefined) {
    console.log('Setting token fee in FeeManager');

    const params = resolveParameters(options);

    const feeManager = await ethers.getContractAt('FeeManager', options.feeManagerAddress);
    await (
      await feeManager.setTokenFee(tokenAddress, params.tokenFee, params.validatorReward, params.liquidityReward)
    ).wait();

    console.log('Successfully seted token fee \n');
  }
}

export async function setLimitPerToken(tokenAddress: string, options: TokenDeploymentOptions): Promise<void> {
  if (options.bridgeValidatorFeePoolAddress !== undefined) {
    console.log('Setting limit per token in BridgeValidatorFeePool');

    const params = resolveParameters(options);

    const bridgeValidatorFeePool = await ethers.getContractAt(
      'BridgeValidatorFeePool',
      options.bridgeValidatorFeePoolAddress
    );
    await (await bridgeValidatorFeePool.setLimitPerToken(tokenAddress, params.tokenLimit)).wait();

    console.log('Successfully seted limit per token \n');
  }
}

export async function addTokenToERC20BridgeMediator(
  tokenAddress: string,
  options: TokenDeploymentOptions
): Promise<void> {
  if (options.homeNetwork !== undefined && options.erc20BridgeMediatorAddress !== undefined) {
    const networkConfig = config.networks[options.homeNetwork] as HttpNetworkConfig;
    const homeProvider = new ethers.providers.JsonRpcProvider(networkConfig.url, networkConfig.chainId);
    const homeSigner = new ethers.Wallet(options.privateKey, homeProvider);

    if (networkConfig.chainId !== undefined) {
      console.log('Adding token to ERC20BridgeMediator');

      const erc20BridgeMediatorFactory = await ethers.getContractFactory('ERC20BridgeMediator', homeSigner);
      const erc20BridgeMediator = erc20BridgeMediatorFactory.attach(options.erc20BridgeMediatorAddress);

      await (await erc20BridgeMediator.addToken(options.symbol, network.config.chainId ?? 1, tokenAddress)).wait();

      console.log('Successfully added token\n');
    }
  }
}

export async function addLiquidity(token: Token, options: TokenDeploymentOptions): Promise<void> {
  if (options.liquidityPoolsAddress !== undefined) {
    console.log('Adding liquidity token');

    const params = resolveParameters(options);

    await (await token.approve(options.liquidityPoolsAddress, params.liquidityAmount)).wait();

    const liquidityPools = await ethers.getContractAt('LiquidityPools', options.liquidityPoolsAddress);
    await (await liquidityPools.addLiquidity(token.address, params.liquidityAmount)).wait();

    console.log('Successfully added liquidity token\n');
  }
}

function resolveParameters(options?: TokenDeploymentOptions): TokenDeploymentParameters {
  let parameters = defaultTokenDeploymentParameters;

  if (options === undefined) {
    return parameters;
  }

  if (options.amount !== undefined) {
    parameters.amount = options.amount;
  }

  if (options.tokenFee !== undefined) {
    parameters.tokenFee = options.tokenFee;
  }

  if (options.tokenLimit !== undefined) {
    parameters.tokenLimit = options.tokenLimit;
  }

  if (options.validatorReward !== undefined) {
    parameters.validatorReward = options.validatorReward;
  }

  if (options.liquidityReward !== undefined) {
    parameters.liquidityReward = options.liquidityReward;
  }

  if (options.liquidityAmount !== undefined) {
    parameters.liquidityAmount = options.liquidityAmount;
  }

  return parameters;
}

export interface TokenDeploymentParameters {
  amount: BigNumber;
  liquidityAmount: BigNumber;
  tokenFee: BigNumber;
  tokenLimit: BigNumber;
  validatorReward: BigNumber;
  liquidityReward: BigNumber;
}

export interface TokenDeploymentOptions {
  privateKey: string;
  symbol: string;
  name: string;
  homeNetwork?: string;
  tokenManagerAddress?: string;
  feeManagerAddress?: string;
  bridgeValidatorFeePoolAddress?: string;
  erc20BridgeMediatorAddress?: string;
  liquidityPoolsAddress?: string;
  amount?: BigNumber;
  liquidityAmount?: BigNumber;
  tokenFee?: BigNumber;
  tokenLimit?: BigNumber;
  validatorReward?: BigNumber;
  liquidityReward?: BigNumber;
}
