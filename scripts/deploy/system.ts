import { ethers } from 'hardhat';
import { Deployer } from './deployer';

import {
  ERC20BridgeMediator,
} from '../../typechain';

const defaultSystemDeploymentParameters: SystemDeploymentParameters = {
  bridgeAppFactory: '0xcB0f2a13098f8e841e6Adfa5B17Ec00508b27665',

  displayLogs: false,
  verify: false,
  stakingKeys: [],
};

export async function deploySystemContracts(options?: SystemDeploymentOptions): Promise<SystemDeploymentResult> {
  const params = resolveParameters(options);
  const deployer = new Deployer(params.displayLogs);

  deployer.log('Deploying contracts\n');

  const res: SystemDeployment = {
    erc20BridgeMediator: await deployer.deploy(ethers.getContractFactory('ERC20BridgeMediator'), 'ERC20BridgeMediator'),
  };

  const bridgeAppFactory = await ethers.getContractAt('IBridgeAppFactory', options!.bridgeAppFactory!);
  await bridgeAppFactory.createApp(), 'Creating BridgeApp';
  const bridgeApp = await ethers.getContractAt('BridgeApp', options!.bridgeAppFactory!);
  await deployer.sendTransaction(bridgeApp.setMediator(res.erc20BridgeMediator.address), 'Setting BridgeMediator to BridgeApp');

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

  return parameters;
}

export interface SystemDeploymentResult extends SystemDeployment, SystemDeploymentParameters { }

export interface SystemDeployment {
  erc20BridgeMediator: ERC20BridgeMediator;
}

export interface SystemDeploymentParameters {
  bridgeAppFactory: string;

  displayLogs: boolean;
  verify: boolean;
  stakingKeys: string[];
}

export interface SystemDeploymentOptions {
  bridgeAppFactory?: string;

  displayLogs?: boolean;
  verify?: boolean;
  stakingKeys?: string[];
}
