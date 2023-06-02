import { BigNumber } from 'ethers';
import { ethers, config, network } from 'hardhat';
import { HttpNetworkConfig } from 'hardhat/types';
import { Token } from '../../typechain';
import { Deployer } from './deployer';

const defaultTokenDeploymentParameters: TokenDeploymentParameters = {
  amount: ethers.utils.parseEther('1000000'),
  liquidityAmount: ethers.utils.parseEther('10000'),
  tokenFee: ethers.utils.parseEther('0.001'),
  tokenLimit: ethers.utils.parseEther('0.0001'),
  validatorReward: ethers.utils.parseEther('0.001'),
  liquidityReward: ethers.utils.parseEther('0.001'),

  displayLogs: false,
  parallelDeployment: false,
  verify: false,
};

export async function deployToken(options: TokenDeploymentOptions): Promise<Token> {
  const params = resolveParameters(options);

  const deployer = new Deployer(params.displayLogs, params.displayLogs);

  const Token = await ethers.getContractFactory('Token');
  const token = await Token.deploy(options.name, options.symbol, params.amount);

  deployer.log(`🧾 Deploying ${options.symbol} token with address ${token.address} (${network.name} chain)`);
  await token.deployed();
  deployer.log(`Successfully deployed ${options.symbol} token\n`);

  const [deployerSigner] = await ethers.getSigners();
  const deployerAddress = await deployerSigner.getAddress();
  deployer.setNonce(await ethers.provider.getTransactionCount(deployerAddress));
  let pendingTxs: Promise<void>[] = [];

  if (options.tokenManagerAddress !== undefined) {
    const tokenManager = await ethers.getContractAt('TokenManager', options.tokenManagerAddress);
    pendingTxs.push(
      deployer.sendTransaction(
        tokenManager.setToken(token.address, '1', deployer.getOverrides()),
        'Setting token in TokenManager'
      )
    );
  }

  if (options.feeManagerAddress !== undefined) {
    const feeManager = await ethers.getContractAt('FeeManager', options.feeManagerAddress);
    pendingTxs.push(
      deployer.sendTransaction(
        feeManager.setTokenFee(
          token.address,
          params.tokenFee,
          params.validatorReward,
          params.liquidityReward,
          deployer.getOverrides()
        ),
        'Setting token in FeeManager'
      )
    );
  }

  if (options.bridgeValidatorFeePoolAddress !== undefined) {
    const bridgeValidatorFeePool = await ethers.getContractAt(
      'BridgeValidatorFeePool',
      options.bridgeValidatorFeePoolAddress
    );

    pendingTxs.push(
      deployer.sendTransaction(
        bridgeValidatorFeePool.setLimitPerToken(token.address, params.tokenLimit, deployer.getOverrides()),
        'Setting limit per token in BridgeValidatorFeePool'
      )
    );
  }

  if (options.liquidityPoolsAddress !== undefined) {
    const liquidityPools = await ethers.getContractAt('LiquidityPools', options.liquidityPoolsAddress);

    pendingTxs.push(
      deployer.sendTransaction(
        token.approve(options.liquidityPoolsAddress, params.liquidityAmount, deployer.getOverrides()),
        'Approving token for LiquidityPools'
      )
    );

    pendingTxs.push(
      deployer.sendTransaction(
        liquidityPools.addLiquidity(token.address, params.liquidityAmount, deployer.getOverrides()),
        'Adding token to LiquidityPools'
      )
    );
  }

  await deployer.waitPromises(pendingTxs);

  if (
    options.homeNetwork !== undefined &&
    options.erc20BridgeMediatorAddress !== undefined &&
    network.name !== 'hardhat'
  ) {
    const networkConfig = config.networks[options.homeNetwork] as HttpNetworkConfig;
    const homeProvider = new ethers.providers.JsonRpcProvider(networkConfig.url, networkConfig.chainId);
    const homeSigner = new ethers.Wallet(options.privateKey, homeProvider);

    if (networkConfig.chainId !== undefined) {
      const params = resolveParameters(options);
      const deployer = new Deployer(params.displayLogs, false);

      const erc20BridgeMediatorFactory = await ethers.getContractFactory('ERC20BridgeMediator', homeSigner);
      const erc20BridgeMediator = erc20BridgeMediatorFactory.attach(options.erc20BridgeMediatorAddress);

      await deployer.sendTransaction(
        erc20BridgeMediator.addToken(options.symbol, network.config.chainId ?? 1, token.address),
        'Adding token to ERC20BridgeMediator'
      );
    }
  }

  return token;
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

  displayLogs: boolean;
  parallelDeployment: boolean;
  verify: boolean;
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

  displayLogs?: boolean;
  parallelDeployment?: boolean;
  verify?: boolean;
}
