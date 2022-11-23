import { BigNumber } from 'ethers';
import { ethers, network, config } from 'hardhat';
import { HttpNetworkConfig } from 'hardhat/types';
import {
  ERC20Bridge,
  FeeManager,
  BridgeValidatorFeePool,
  LiquidityPools,
  TokenManager,
  IBridgeApp__factory,
} from '../../typechain';
import { Deployer } from './deployer';

const defaultBridgeDeploymentParameters: BridgeDeploymentParameters = {
  feePercentage: BigNumber.from('10000000000000000'),
  validatorRefundFee: BigNumber.from('10000000000000000'),
  homeChainId: BigNumber.from(1),
  foundationAddress: '0x0000000000000000000000000000000000000001',
  bridgeApp: '0x0000000000000000000000000000000000000001',
  relayBridge: '0x0000000000000000000000000000000000000001',
  signerStorage: '0x0000000000000000000000000000000000000001',
  setApp: false,

  displayLogs: false,
  verify: false,
};

export async function deployBridgeContracts(options?: BridgeDeploymentOptions): Promise<BridgeDeploymentResult> {
  const params = resolveParameters(options);
  const deployer = new Deployer(params.displayLogs);

  const [validator] = await ethers.getSigners();

  if (network.name === 'hardhat') {
    const signerStorage = await deployer.deploy(ethers.getContractFactory('MockSignerStorage'), 'SignerStorage');
    await signerStorage.initialize(validator.address);

    params.signerStorage = signerStorage.address;
  }

  deployer.log('Deploying contracts\n');

  const res: BridgeDeployment = {
    tokenManager: await deployer.deploy(ethers.getContractFactory('TokenManager'), 'TokenManager'),
    feeManager: await deployer.deploy(ethers.getContractFactory('FeeManager'), 'FeeManager'),
    bridgeValidatorFeePool: await deployer.deploy(
      ethers.getContractFactory('BridgeValidatorFeePool'),
      'BridgeValidatorFeePool'
    ),
    liquidityPools: await deployer.deploy(ethers.getContractFactory('LiquidityPools'), 'LiquidityPools'),
    erc20Bridge: await deployer.deploy(ethers.getContractFactory('ERC20Bridge'), 'ERC20Bridge'),
  };

  deployer.log('Successfully deployed contracts\n');

  deployer.log('Initializing contracts\n');

  await deployer.sendTransaction(res.tokenManager.initialize(params.signerStorage), 'Initializing TokenManager');

  await deployer.sendTransaction(
    res.feeManager.initialize(
      params.signerStorage,
      res.liquidityPools.address,
      params.foundationAddress,
      res.bridgeValidatorFeePool.address,
      params.validatorRefundFee
    ),
    'Initializing FeeManager'
  );

  await deployer.sendTransaction(
    res.liquidityPools.initialize(
      params.signerStorage,
      res.tokenManager.address,
      res.erc20Bridge.address,
      res.feeManager.address,
      params.feePercentage
    ),
    'Initializing LiquidityPools'
  );

  await deployer.sendTransaction(
    res.erc20Bridge.initialize(
      params.signerStorage,
      res.tokenManager.address,
      res.liquidityPools.address,
      res.feeManager.address,
      params.bridgeApp,
      params.relayBridge
    ),
    'Initializing ERC20Bridge'
  );

  await deployer.sendTransaction(
    res.bridgeValidatorFeePool.initialize(
      params.signerStorage,
      res.erc20Bridge.address,
      validator.address,
      params.homeChainId
    ),
    'Initializing BridgeValidatorFeePool'
  );

  if (params.setApp && options?.homeNetwork !== undefined && options.privateKey !== undefined) {
    const networkConfig = config.networks[options.homeNetwork] as HttpNetworkConfig;
    const homeProvider = new ethers.providers.JsonRpcProvider(networkConfig.url, networkConfig.chainId);
    const homeSigner = new ethers.Wallet(options.privateKey, homeProvider);

    if (network.config.chainId !== undefined) {
      const bridgeApp = IBridgeApp__factory.connect(params.bridgeApp, homeSigner);

      await deployer.sendTransaction(
        bridgeApp.setContractAddress(network.config.chainId, res.erc20Bridge.address),
        'Setting bridge address in BridgeApp'
      );
    }
  }

  deployer.log('Successfully initialized contracts\n');

  if (params.verify) {
    await deployer.verifyObjectValues(res);
  }

  return {
    ...res,
    ...params,
  };
}

function resolveParameters(options?: BridgeDeploymentOptions): BridgeDeploymentParameters {
  let parameters = defaultBridgeDeploymentParameters;

  if (options === undefined) {
    return parameters;
  }

  if (options.feePercentage !== undefined) {
    parameters.feePercentage = options.feePercentage;
  }

  if (options.validatorRefundFee !== undefined) {
    parameters.validatorRefundFee = options.validatorRefundFee;
  }

  if (options.homeChainId !== undefined) {
    parameters.homeChainId = options.homeChainId;
  }

  if (options.foundationAddress !== undefined) {
    parameters.foundationAddress = options.foundationAddress;
  }

  if (options.bridgeApp !== undefined) {
    parameters.bridgeApp = options.bridgeApp;
  }

  if (options.relayBridge !== undefined) {
    parameters.relayBridge = options.relayBridge;
  }

  if (options.signerStorage !== undefined) {
    parameters.signerStorage = options.signerStorage;
  }

  if (options.setApp !== undefined) {
    parameters.setApp = options.setApp;
  }

  if (options.displayLogs !== undefined) {
    parameters.displayLogs = options.displayLogs;
  }

  if (options.verify !== undefined) {
    parameters.verify = options.verify;
  }

  return parameters;
}

export interface BridgeDeploymentResult extends BridgeDeployment, BridgeDeploymentParameters {}

export interface BridgeDeployment {
  tokenManager: TokenManager;
  feeManager: FeeManager;
  bridgeValidatorFeePool: BridgeValidatorFeePool;
  liquidityPools: LiquidityPools;
  erc20Bridge: ERC20Bridge;
}

export interface BridgeDeploymentParameters {
  feePercentage: BigNumber;
  validatorRefundFee: BigNumber;
  homeChainId: BigNumber;
  foundationAddress: string;
  bridgeApp: string;
  relayBridge: string;
  signerStorage: string;
  setApp: boolean;

  displayLogs: boolean;
  verify: boolean;
}

export interface BridgeDeploymentOptions {
  feePercentage?: BigNumber;
  validatorRefundFee?: BigNumber;
  homeChainId?: BigNumber;
  foundationAddress?: string;
  bridgeApp?: string;
  relayBridge?: string;
  signerStorage?: string;
  displayLogs?: boolean;
  setApp?: boolean;
  homeNetwork?: string;
  privateKey?: string;

  verify?: boolean;
  deployMocks?: boolean;
}
