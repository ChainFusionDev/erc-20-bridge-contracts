import { BigNumber } from 'ethers';
import { ethers, config, network } from 'hardhat';
import { HttpNetworkConfig } from 'hardhat/types';
import { Token } from '../../typechain';
import { Deployer } from './deployer';

const defaultTokenDeploymentParameters: TokenDeploymentParameters = {
  amount: ethers.utils.parseEther('10000'),
  liquidityAmount: ethers.utils.parseEther('1000'),
  tokenFee: ethers.utils.parseEther('0.001'),
  tokenLimit: ethers.utils.parseEther('0.0001'),
  validatorReward: ethers.utils.parseEther('0.001'),
  liquidityReward: ethers.utils.parseEther('0.001'),

  displayLogs: false,
  verify: false,
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
    const params = resolveParameters(options);
    const deployer = new Deployer(params.displayLogs);

    const tokenManager = await ethers.getContractAt('TokenManager', options.tokenManagerAddress);
    await deployer.sendTransaction(tokenManager.setToken(tokenAddress, '1'), 'Setting token in TokenManager');
  }
}

export async function setTokenFee(tokenAddress: string, options: TokenDeploymentOptions): Promise<void> {
  if (options.feeManagerAddress !== undefined) {
    const params = resolveParameters(options);
    const deployer = new Deployer(params.displayLogs);

    const feeManager = await ethers.getContractAt('FeeManager', options.feeManagerAddress);

    await deployer.sendTransaction(
      feeManager.setTokenFee(tokenAddress, params.tokenFee, params.validatorReward, params.liquidityReward),
      'Setting token in FeeManager'
    );
  }
}

export async function setLimitPerToken(tokenAddress: string, options: TokenDeploymentOptions): Promise<void> {
  if (options.bridgeValidatorFeePoolAddress !== undefined) {
    const params = resolveParameters(options);
    const deployer = new Deployer(params.displayLogs);

    const bridgeValidatorFeePool = await ethers.getContractAt(
      'BridgeValidatorFeePool',
      options.bridgeValidatorFeePoolAddress
    );

    await deployer.sendTransaction(
      bridgeValidatorFeePool.setLimitPerToken(tokenAddress, params.tokenLimit),
      'Setting limit per token in BridgeValidatorFeePool'
    );
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
      const params = resolveParameters(options);
      const deployer = new Deployer(params.displayLogs);

      const erc20BridgeMediatorFactory = await ethers.getContractFactory('ERC20BridgeMediator', homeSigner);
      const erc20BridgeMediator = erc20BridgeMediatorFactory.attach(options.erc20BridgeMediatorAddress);

      await deployer.sendTransaction(
        erc20BridgeMediator.addToken(options.symbol, network.config.chainId ?? 1, tokenAddress),
        'Adding token to ERC20BridgeMediator'
      );
    }
  }
}

export async function addLiquidity(token: Token, options: TokenDeploymentOptions): Promise<void> {
  if (options.liquidityPoolsAddress !== undefined) {
    const params = resolveParameters(options);
    const deployer = new Deployer(params.displayLogs);

    const liquidityPools = await ethers.getContractAt('LiquidityPools', options.liquidityPoolsAddress);

    await deployer.sendTransaction(
      token.approve(options.liquidityPoolsAddress, params.liquidityAmount),
      'Approving token for LiquidityPools'
    );

    await deployer.sendTransaction(
      liquidityPools.addLiquidity(token.address, params.liquidityAmount),
      'Adding token to LiquidityPools'
    );
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

  if (options.verify !== undefined) {
    parameters.verify = options.verify;
  }

  if (options.displayLogs !== undefined) {
    parameters.displayLogs = options.displayLogs;
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

  verify: boolean;
  displayLogs: boolean;
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

  verify?: boolean;
  displayLogs?: boolean;
}
