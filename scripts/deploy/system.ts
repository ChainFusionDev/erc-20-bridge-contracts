import { ethers } from 'hardhat';
import { Deployer } from './deployer';
import { readContractsConfig, updateContractsConfig, writeContractsConfig } from './config';

import { ERC20BridgeMediator } from '../../typechain';

const defaultSystemDeploymentParameters: SystemDeploymentParameters = {
  displayLogs: false,
  verify: false,
};

export async function deploySystemContracts(options?: SystemDeploymentOptions): Promise<SystemDeploymentResult> {
  const params = resolveParameters(options);
  const deployer = new Deployer(params.displayLogs);

  deployer.log('Deploying contracts\n');

  const res: SystemDeployment = {
    erc20BridgeMediator: await deployer.deploy(ethers.getContractFactory('ERC20BridgeMediator'), 'ERC20BridgeMediator'),
  };

  if (params.bridgeAppFactory !== undefined) {
    const bridgeAppFactory = await ethers.getContractAt('IBridgeAppFactory', params.bridgeAppFactory);
    const bridgeAppAddress = await bridgeAppFactory.callStatic.createApp();
    await (await bridgeAppFactory.createApp()).wait();

    const bridgeApp = await ethers.getContractAt('IBridgeApp', bridgeAppAddress);
    await bridgeApp.setMediator(res.erc20BridgeMediator.address);

    const contractsConfig = await readContractsConfig();
    contractsConfig.bridgeAppFactory = bridgeAppFactory.address;
    contractsConfig.bridgeApp = bridgeApp.address;
    updateContractsConfig(contractsConfig, res);

    await writeContractsConfig(contractsConfig);
  }

  deployer.log('Successfully deployed contracts\n');

  deployer.log('Initializing contracts\n');

  deployer.log('Successfully initialized contracts\n');

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

  if (options.bridgeAppFactory !== undefined) {
    parameters.bridgeAppFactory = options.bridgeAppFactory;
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
}

export interface SystemDeploymentParameters {
  bridgeAppFactory?: string;

  displayLogs: boolean;
  verify: boolean;
}

export interface SystemDeploymentOptions {
  bridgeAppFactory?: string;

  displayLogs?: boolean;
  verify?: boolean;
}
