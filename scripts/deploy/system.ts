import { ethers } from 'hardhat';
import { Deployer } from './deployer';

import { ERC20BridgeMediator, IBridgeApp, IBridgeAppFactory } from '../../typechain';

const defaultSystemDeploymentParameters: SystemDeploymentParameters = {
  displayLogs: false,
  verify: false,
};

export async function deploySystemContracts(options?: SystemDeploymentOptions): Promise<SystemDeploymentResult> {
  const params = resolveParameters(options);
  const deployer = new Deployer(params.displayLogs);

  deployer.log('Deploying contracts\n');

  let res: SystemDeployment = {
    erc20BridgeMediator: await deployer.deploy(ethers.getContractFactory('ERC20BridgeMediator'), 'ERC20BridgeMediator'),
  };

  if (params.bridgeAppFactoryAddress !== undefined) {
    const bridgeAppFactory = await ethers.getContractAt('IBridgeAppFactory', params.bridgeAppFactoryAddress);
    const bridgeAppAddress = await bridgeAppFactory.callStatic.createApp();
    await (await bridgeAppFactory.createApp()).wait();

    const bridgeApp = await ethers.getContractAt('IBridgeApp', bridgeAppAddress);
    await bridgeApp.setMediator(res.erc20BridgeMediator.address);

    res.bridgeAppFactory = bridgeAppFactory;
    res.bridgeApp = bridgeApp;
  }

  deployer.log('Successfully deployed contracts\n');

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
  verify: boolean;
}

export interface SystemDeploymentOptions {
  bridgeAppFactoryAddress?: string;

  displayLogs?: boolean;
  verify?: boolean;
}
