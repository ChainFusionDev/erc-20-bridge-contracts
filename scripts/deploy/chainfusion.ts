import { ethers, network } from 'hardhat';
import { Deployer } from './deployer';

import { ERC20BridgeMediator, IBridgeApp, IBridgeAppFactory } from '../../typechain';

const defaultSystemDeploymentParameters: SystemDeploymentParameters = {
  displayLogs: false,
  parallelDeployment: false,
  verify: false,
};

export async function deploySystemContracts(options?: SystemDeploymentOptions): Promise<SystemDeploymentResult> {
  const params = resolveParameters(options);
  const deployer = new Deployer(params.displayLogs, params.parallelDeployment);

  const [deployerSigner] = await ethers.getSigners();
  const deployerAddress = await deployerSigner.getAddress();
  deployer.setNonce(await ethers.provider.getTransactionCount(deployerAddress));

  deployer.log(`🧾 Deploying ERC-20 Bridge contracts on ChainFusion (${network.name} chain)\n`);

  let res: SystemDeployment = {
    erc20BridgeMediator: await deployer.deploy(ethers.getContractFactory('ERC20BridgeMediator'), 'ERC20BridgeMediator'),
  };

  deployer.log('Successfully deployed ERC-20 Bridge contracts on ChainFusion\n');

  if (network.name !== 'hardhat' && params.bridgeAppFactoryAddress !== undefined) {
    const bridgeAppFactory = await ethers.getContractAt('IBridgeAppFactory', params.bridgeAppFactoryAddress);
    const bridgeAppAddress = await bridgeAppFactory.callStatic.createApp();

    await deployer.sendTransaction(bridgeAppFactory.createApp(), 'Creating BridgeApp');

    const bridgeApp = await ethers.getContractAt('IBridgeApp', bridgeAppAddress);

    await deployer.sendTransaction(
      bridgeApp.setMediator(res.erc20BridgeMediator.address),
      'Setting ERC20BridgeMediator in BridgeApp'
    );

    res.bridgeAppFactory = bridgeAppFactory;
    res.bridgeApp = bridgeApp;
  }

  if (params.verify) {
    await deployer.verifyObjectValues(res);
  }

  return {
    ...res,
    ...params,
  };
}

function resolveParameters(options?: SystemDeploymentOptions): SystemDeploymentParameters {
  let parameters = defaultSystemDeploymentParameters;

  if (options === undefined) {
    return parameters;
  }

  if (options.bridgeAppFactoryAddress !== undefined) {
    parameters.bridgeAppFactoryAddress = options.bridgeAppFactoryAddress;
  }

  if (options.displayLogs !== undefined) {
    parameters.displayLogs = options.displayLogs;
  }

  if (options.parallelDeployment !== undefined) {
    parameters.parallelDeployment = options.parallelDeployment;
  }

  if (options.verify !== undefined) {
    parameters.verify = options.verify;
  }

  return parameters;
}

export interface SystemDeploymentResult extends SystemDeployment, SystemDeploymentParameters {}

export interface SystemDeployment {
  erc20BridgeMediator: ERC20BridgeMediator;
  bridgeAppFactory?: IBridgeAppFactory;
  bridgeApp?: IBridgeApp;
}

export interface SystemDeploymentParameters {
  bridgeAppFactoryAddress?: string;

  displayLogs: boolean;
  parallelDeployment: boolean;
  verify: boolean;
}

export interface SystemDeploymentOptions {
  bridgeAppFactoryAddress?: string;

  displayLogs?: boolean;
  parallelDeployment?: boolean;
  verify?: boolean;
}
